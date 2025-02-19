package server

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	sqlc "github.com/vietddude/tss-impl/db/sqlc"
	"github.com/vietddude/tss-impl/party"
	pb "github.com/vietddude/tss-impl/proto"
	"github.com/vietddude/tss-impl/utils"
	"go.uber.org/zap"
)

// Helper function to publish results to Redis
func (s *MPCServer) publishToRedis(sessionID string, data map[string]interface{}, channel string) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal result: %w", err)
	}

	if err := s.redisClient.Publish(context.TODO(), channel, jsonData).Err(); err != nil {
		return fmt.Errorf("failed to publish result: %w", err)
	}

	s.logger.Debug("published result to Redis", zap.String("session_id", sessionID))
	return nil
}

// Helper function to get or create a party
func (s *MPCServer) getOrCreateParty(sessionID string) *party.Party {
	s.partiesMu.Lock()
	defer s.partiesMu.Unlock()

	if p, exists := s.parties[sessionID]; exists {
		return p
	}

	p := party.NewParty(uint16(s.nodeID), s.logger.Sugar())
	s.parties[sessionID] = p
	return p
}

func (s *MPCServer) removeParty(sessionID string) {
	s.partiesMu.Lock()
	delete(s.parties, sessionID)
	s.partiesMu.Unlock()
}

func (s *MPCServer) notifyPeers(ctx context.Context, req *pb.ActionRequest) error {
	s.peerConnsMu.RLock()
	peerConns := s.peerConns
	s.peerConnsMu.RUnlock()

	var wg sync.WaitGroup
	errChan := make(chan error, len(peerConns))

	for addr, conn := range peerConns {
		if conn == nil {
			continue
		}

		wg.Add(1)
		go func(addr string, conn *peerConnection) {
			defer wg.Done()

			conn.mu.Lock()
			client := conn.client
			conn.mu.Unlock()

			if _, err := client.NotifyAction(ctx, req); err != nil {
				s.logger.Error("failed to notify peer",
					zap.String("address", addr),
					zap.Error(err))
				errChan <- fmt.Errorf("failed to notify %s: %w", addr, err)
			}
		}(addr, conn)
	}

	wg.Wait()
	close(errChan)

	var errs []string
	for err := range errChan {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		return fmt.Errorf("peer notification failed:\n%s", strings.Join(errs, "\n"))
	}

	return nil
}

// Helper function to fetch and decrypt share data
func (s *MPCServer) getDecryptedShareData(ctx context.Context, sessionID string, shareData []byte) ([]byte, error) {
	if shareData != nil {
		s.logger.Debug("Using provided share data")
		return utils.DecryptAESGCM(shareData, []byte(s.cfg.EncryptKey))
	}

	q := sqlc.New(s.dbPool)
	sessionUUID := utils.StringToPgUUID(sessionID)

	var encryptedShare []byte
	var err error

	switch s.nodeID {
	case 2:
		encryptedShare, err = q.GetShareKey1(ctx, sessionUUID)
	case 3:
		encryptedShare, err = q.GetShareKey2(ctx, sessionUUID)
	default:
		return nil, fmt.Errorf("unsupported nodeID: %d", s.nodeID)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get share key: %w", err)
	}

	return utils.DecryptAESGCM(encryptedShare, []byte(s.cfg.EncryptKey))
}

// Helper function to create a sender function for parties
func (s *MPCServer) createSenderFunc(sessionID string) func([]byte, bool, uint16) {
	return func(msg []byte, broadcast bool, to uint16) {
		streamMsg := &pb.TSSMessage{
			SessionId: sessionID,
			From:      uint32(s.nodeID),
			To:        uint32(to),
			Payload:   msg,
			Broadcast: broadcast,
		}

		if broadcast {
			s.peersMu.RLock()
			for peerID, peerAddr := range s.peers {
				if peerID != s.nodeID {
					go s.sendToPeer(peerAddr, streamMsg)
				}
			}
			s.peersMu.RUnlock()
		} else {
			s.peersMu.RLock()
			peerAddr, ok := s.peers[uint32(to)]
			s.peersMu.RUnlock()
			if ok {
				go s.sendToPeer(peerAddr, streamMsg)
			}
		}
	}
}
