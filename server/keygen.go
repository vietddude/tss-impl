package server

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	sqlc "github.com/vietddude/tss-impl/db/sqlc"
	pb "github.com/vietddude/tss-impl/proto"
	"github.com/vietddude/tss-impl/utils"
	"go.uber.org/zap"
)

func (s *MPCServer) Keygen(ctx context.Context, sessionID string, parties []uint32, threshold int) (string, []byte, error) {
	p := s.getOrCreateParty(sessionID)
	defer s.removeParty(sessionID)

	p.Init(utils.ConvertToUint16(parties), threshold, s.createSenderFunc(sessionID))

	shareData, err := p.KeyGen(ctx)
	if err != nil {
		return "", nil, fmt.Errorf("keygen failed: %w", err)
	}
	p.SetShareData(shareData)

	pk, err := p.TPubKey()
	if err != nil {
		return "", nil, fmt.Errorf("failed to get threshold public key: %w", err)
	}
	return crypto.PubkeyToAddress(*pk).Hex(), shareData, nil
}

func (s *MPCServer) InitKeygen(ctx context.Context, sessionID string, parties []uint32, threshold int) error {
	s.logger.Info("initiating keygen process", zap.String("session_id", sessionID))

	req := &pb.ActionRequest{
		SessionId: sessionID,
		Parties:   parties,
		Threshold: uint32(threshold),
		Action:    pb.Action_KEYGEN,
	}

	if err := s.notifyPeers(ctx, req); err != nil {
		return fmt.Errorf("peer notification failed: %w", err)
	}

	go s.runKeygen(sessionID, parties, threshold)
	return nil
}

// Helper function to run Keygen and publish results
func (s *MPCServer) runKeygen(sessionID string, parties []uint32, threshold int) {
	ctx := context.Background()

	pubKey, shareData, err := s.Keygen(ctx, sessionID, parties, threshold)
	if err != nil {
		s.logger.Error("keygen failed",
			zap.String("session_id", sessionID),
			zap.Error(err))
		return
	}

	s.logger.Info("keygen completed",
		zap.String("session_id", sessionID),
		zap.String("pub_key", pubKey))

	encrypted, err := utils.EncryptAESGCM(shareData, []byte(s.cfg.EncryptKey))
	if err != nil {
		s.logger.Error("failed to encrypt share data",
			zap.String("session_id", sessionID),
			zap.Error(err))
		return
	}

	data := buildKegenResponse(sessionID, pubKey, encrypted)
	if err := s.publishToRedis(sessionID, data, fmt.Sprintf("keygen:%s", sessionID)); err != nil {
		s.logger.Error("failed to publish results",
			zap.String("session_id", sessionID),
			zap.Error(err))
	}
}

func (s *MPCServer) handleKeygen(ctx context.Context, req *pb.ActionRequest) error {
	_, shareData, err := s.Keygen(ctx, req.SessionId, req.Parties, int(req.Threshold))
	if err != nil {
		return fmt.Errorf("keygen failed: %w", err)
	}

	encrypted, err := utils.EncryptAESGCM(shareData, []byte(s.cfg.EncryptKey))
	if err != nil {
		return fmt.Errorf("failed to encrypt share data: %w", err)
	}

	sessionUUID := utils.StringToPgUUID(req.SessionId)
	q := sqlc.New(s.dbPool)

	var dbErr error
	switch s.nodeID {
	case 2:
		dbErr = q.InsertShareKey1(ctx, sqlc.InsertShareKey1Params{
			SessionID:      sessionUUID,
			EncryptedShare: encrypted,
		})
	case 3:
		dbErr = q.InsertShareKey2(ctx, sqlc.InsertShareKey2Params{
			SessionID:      sessionUUID,
			EncryptedShare: encrypted,
		})
	}

	if dbErr != nil {
		return fmt.Errorf("failed to insert share key: %w", dbErr)
	}

	return nil
}

func buildKegenResponse(sessionID string, pubKey string, shareData []byte) map[string]interface{} {
	return map[string]interface{}{
		"sesion_id":  sessionID,
		"pub_key":    pubKey,
		"share_data": base64.StdEncoding.EncodeToString(shareData),
	}
}
