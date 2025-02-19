package server

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/vietddude/tss-impl/config"
	"github.com/vietddude/tss-impl/party"
	pb "github.com/vietddude/tss-impl/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type peerConnection struct {
	client pb.MPCServiceClient
	stream pb.MPCService_StreamMessagesClient
	mu     sync.Mutex
}

type MPCServer struct {
	pb.UnimplementedMPCServiceServer
	parties     map[string]*party.Party
	partiesMu   sync.RWMutex
	peers       map[uint32]string
	peersMu     sync.RWMutex
	peerConns   map[string]*peerConnection
	peerConnsMu sync.RWMutex
	logger      *zap.Logger
	nodeID      uint32
	dbPool      *pgxpool.Pool
	cfg         *config.Config
	redisClient *redis.Client
}

func NewMPCServer(nodeID uint32, pool *pgxpool.Pool, cfg *config.Config) *MPCServer {
	logger, _ := zap.NewDevelopment()
	return &MPCServer{
		parties:   make(map[string]*party.Party),
		peers:     make(map[uint32]string),
		peerConns: make(map[string]*peerConnection),
		logger:    logger.With(zap.Uint32("node_id", nodeID)),
		nodeID:    nodeID,
		dbPool:    pool,
		cfg:       cfg,
		redisClient: redis.NewClient(&redis.Options{
			Addr: cfg.RedisAddr}),
	}
}

func (s *MPCServer) ConnectToPeers() {
	s.peersMu.RLock()
	defer s.peersMu.RUnlock()

	for peerID, peerAddr := range s.peers {
		s.peerConnsMu.RLock()
		_, exists := s.peerConns[peerAddr]
		s.peerConnsMu.RUnlock()
		if exists {
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		conn, err := grpc.DialContext(ctx, peerAddr, grpc.WithInsecure())
		cancel()
		if err != nil {
			s.logger.Error("Failed to connect to peer", zap.String("address", peerAddr), zap.Error(err))
			continue
		}

		client := pb.NewMPCServiceClient(conn)
		stream, err := client.StreamMessages(context.Background())
		if err != nil {
			s.logger.Error("Failed to create stream", zap.String("address", peerAddr), zap.Error(err))
			conn.Close()
			continue
		}

		s.peerConnsMu.Lock()
		s.peerConns[peerAddr] = &peerConnection{
			client: client,
			stream: stream,
		}
		s.peerConnsMu.Unlock()
		s.logger.Info("Connected to peer", zap.Uint32("peer_id", peerID), zap.String("address", peerAddr))
	}
}

func (s *MPCServer) getPeerConnection(peerAddr string) (*peerConnection, error) {
	s.peerConnsMu.RLock()
	peerConn, exists := s.peerConns[peerAddr]
	s.peerConnsMu.RUnlock()
	if exists {
		return peerConn, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, peerAddr, grpc.WithInsecure())
	cancel()
	if err != nil {
		return nil, err
	}

	client := pb.NewMPCServiceClient(conn)
	stream, err := client.StreamMessages(context.Background())
	if err != nil {
		conn.Close()
		return nil, err
	}

	peerConn = &peerConnection{
		client: client,
		stream: stream,
	}

	s.peerConnsMu.Lock()
	s.peerConns[peerAddr] = peerConn
	s.peerConnsMu.Unlock()

	return peerConn, nil
}

func (s *MPCServer) sendToPeer(peerAddr string, msg *pb.TSSMessage) error {
	peerConn, err := s.getPeerConnection(peerAddr)
	if err != nil {
		s.logger.Error("Failed to get peer connection", zap.String("address", peerAddr), zap.Error(err))
		return err
	}

	peerConn.mu.Lock()
	defer peerConn.mu.Unlock()

	if err = peerConn.stream.Send(msg); err != nil {
		newStream, streamErr := peerConn.client.StreamMessages(context.Background())
		if streamErr != nil {
			s.logger.Error("Failed to recreate stream", zap.String("address", peerAddr), zap.Error(streamErr))
			return streamErr
		}
		peerConn.stream = newStream
		return peerConn.stream.Send(msg)
	}
	return nil
}

func (s *MPCServer) AddParty(sessionID string, p *party.Party) {
	s.partiesMu.Lock()
	s.parties[sessionID] = p
	s.partiesMu.Unlock()
}

func (s *MPCServer) AddPeer(id uint32, address string) {
	s.peersMu.Lock()
	s.peers[id] = address
	s.peersMu.Unlock()
	s.logger.Info("Peer added", zap.Uint32("peer_id", id), zap.String("address", address))
}

// GRPC methods
func (s *MPCServer) StreamMessages(stream pb.MPCService_StreamMessagesServer) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			s.logger.Error("Failed to receive message", zap.Error(err))
			return err
		}

		s.partiesMu.RLock()
		p, exists := s.parties[msg.SessionId]
		s.partiesMu.RUnlock()
		if !exists {
			s.logger.Error("Party not found", zap.String("session_id", msg.SessionId))
			continue
		}
		p.OnMsg(msg.Payload, uint16(msg.From), msg.Broadcast)
	}
}

func (s *MPCServer) NotifyAction(ctx context.Context, req *pb.ActionRequest) (*pb.ActionResponse, error) {
	handler, exists := map[pb.Action]func(context.Context, *pb.ActionRequest) error{
		pb.Action_KEYGEN: func(ctx context.Context, req *pb.ActionRequest) error {
			return s.handleKeygen(ctx, req)
		},
		pb.Action_INIT_KEYGEN: func(ctx context.Context, req *pb.ActionRequest) error {
			return s.InitKeygen(ctx, req.SessionId, req.Parties, int(req.Threshold))
		},
		pb.Action_SIGN: func(ctx context.Context, req *pb.ActionRequest) error {
			_, err := s.Sign(ctx, req.SessionId, req.Parties, int(req.Threshold), req.MsgHash, req.ShareData)
			return err
		},
		pb.Action_INIT_SIGN: func(ctx context.Context, req *pb.ActionRequest) error {
			return s.InitSign(ctx, req.SessionId, req.Parties, int(req.Threshold), req.MsgHash, req.ShareData)
		},
	}[req.Action]

	if !exists {
		return &pb.ActionResponse{Success: false}, fmt.Errorf("unknown action: %v", req.Action)
	}

	go func() {
		if err := handler(context.Background(), req); err != nil {
			s.logger.Error("action handler failed", zap.Error(err))
		}
	}()

	return &pb.ActionResponse{Success: true}, nil
}
