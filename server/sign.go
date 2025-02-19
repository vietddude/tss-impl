package server

import (
	"context"
	"encoding/base64"
	"fmt"

	pb "github.com/vietddude/tss-impl/proto"
	"github.com/vietddude/tss-impl/utils"
	"go.uber.org/zap"
)

func (s *MPCServer) Sign(ctx context.Context, sessionID string, parties []uint32, threshold int, msgHash []byte, shareData []byte) ([]byte, error) {
	p := s.getOrCreateParty(sessionID)
	defer s.removeParty(sessionID)

	p.Init(utils.ConvertToUint16(parties), threshold, s.createSenderFunc(sessionID))

	decryptedShare, err := s.getDecryptedShareData(ctx, sessionID, shareData)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt share data: %w", err)
	}
	p.SetShareData(decryptedShare)

	sig, err := p.Sign(ctx, msgHash)
	if err != nil {
		return nil, fmt.Errorf("signing failed: %w", err)
	}

	return sig, nil
}

func (s *MPCServer) InitSign(ctx context.Context, sessionID string, parties []uint32, threshold int, msgHash []byte, shareData []byte) error {
	s.logger.Info("initiating sign process", zap.String("session_id", sessionID))

	req := &pb.ActionRequest{
		SessionId: sessionID,
		Parties:   parties,
		Threshold: uint32(threshold),
		Action:    pb.Action_SIGN,
		MsgHash:   msgHash,
	}

	if err := s.notifyPeers(ctx, req); err != nil {
		return fmt.Errorf("peer notification failed: %w", err)
	}

	go s.runSign(sessionID, parties, threshold, msgHash, shareData)
	return nil
}

func (s *MPCServer) runSign(sessionID string, parties []uint32, threshold int, msgHash []byte, shareData []byte) {
	ctx := context.Background()

	sig, err := s.Sign(ctx, sessionID, parties, threshold, msgHash, shareData)
	if err != nil {
		s.logger.Error("signing failed",
			zap.String("session_id", sessionID),
			zap.Error(err))
		return
	}

	s.logger.Info("signing completed",
		zap.String("session_id", sessionID),
		zap.String("signature", base64.StdEncoding.EncodeToString(sig)))

	// Publish signature to Redis
	data := buildSignResponse(sessionID, sig)
	if err := s.publishToRedis(sessionID, data, fmt.Sprintf("sign:%s", sessionID)); err != nil {
		s.logger.Error("failed to publish signature to Redis", zap.Error(err))
	}
}

func buildSignResponse(sessionID string, sig []byte) map[string]interface{} {
	return map[string]interface{}{
		"session_id": sessionID,
		"signature":  base64.StdEncoding.EncodeToString(sig),
	}
}
