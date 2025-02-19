package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vietddude/tss-impl/config"
	"github.com/vietddude/tss-impl/db"
	"github.com/vietddude/tss-impl/proto"
	"github.com/vietddude/tss-impl/server"
	"google.golang.org/grpc"
)

func initializeServers(cfg *config.Config) ([]config.Node, error) {
	parts := strings.Split(cfg.Node.Address, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid address format: %s", cfg.Node.Address)
	}

	basePort, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid port number: %v", err)
	}

	servers := make([]config.Node, 0, cfg.NodeNumber)
	for i := 0; i < cfg.NodeNumber; i++ {
		server := config.Node{
			ID:      cfg.Node.ID + uint32(i),
			Address: fmt.Sprintf("localhost:%d", basePort+i),
		}
		servers = append(servers, server)
	}

	return servers, nil
}

func createPeerMap(servers []config.Node) *sync.Map {
	peerMap := &sync.Map{}
	for _, srv := range servers {
		peerMap.Store(srv.ID, srv.Address)
	}
	return peerMap
}

func startServer(id uint32, address string, peerMap *sync.Map, wg *sync.WaitGroup, dbPool *pgxpool.Pool, cfg *config.Config) {
	defer wg.Done()

	mpcServer := server.NewMPCServer(id, dbPool, cfg)
	setupPeerConnections(mpcServer, id, peerMap)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Printf("Server %d failed to listen on %s: %v", id, address, err)
		return
	}

	grpcServer := grpc.NewServer()
	proto.RegisterMPCServiceServer(grpcServer, mpcServer)

	log.Printf("Server %d is running on %s", id, address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("Server %d failed to serve: %v", id, err)
	}
}

func setupPeerConnections(mpcServer *server.MPCServer, id uint32, peerMap *sync.Map) {
	peerMap.Range(func(key, value interface{}) bool {
		peerID := key.(uint32)
		peerAddr := value.(string)
		if peerID != id {
			mpcServer.AddPeer(peerID, peerAddr)
			go mpcServer.ConnectToPeers()
		}
		return true
	})
}

func startAllServers(servers []config.Node, peerMap *sync.Map, dbPool *pgxpool.Pool, cfg *config.Config) {
	var wg sync.WaitGroup
	wg.Add(len(servers))

	for _, srv := range servers {
		go startServer(srv.ID, srv.Address, peerMap, &wg, dbPool, cfg)
	}

	wg.Wait()
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	servers, err := initializeServers(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize servers: %v", err)
	}

	dbPool, err := db.InitDB(&cfg.DB)
	if err != nil {
		log.Fatalf("Failed to initialize DB: %v", err)
	}
	defer dbPool.Close()

	peerMap := createPeerMap(servers)
	startAllServers(servers, peerMap, dbPool, cfg)
}
