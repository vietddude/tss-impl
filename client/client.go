package main

import (
	"context"
	"crypto/elliptic"
	"encoding/base64"
	"encoding/json"
	"log"
	"time"

	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/redis/go-redis/v9"
	pb "github.com/vietddude/tss-impl/proto"
	"github.com/vietddude/tss-impl/utils"
	"google.golang.org/grpc"
)

func init() {
	tss.RegisterCurve("elliptic.p256Curve", elliptic.P256())
}
func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMPCServiceClient(conn)
	// redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	// Đầu tiên thực hiện keygen
	// sessionID := uuid.NewString()
	// notifyCtx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	// defer cancel()

	// // Gửi yêu cầu NotifyAction cho keygen
	// _, err = client.NotifyAction(notifyCtx, &pb.ActionRequest{
	// 	SessionId: sessionID,
	// 	Parties:   []uint32{1, 2, 3},
	// 	Threshold: 2,
	// 	Action:    pb.Action_INIT_KEYGEN,
	// })
	// if err != nil {
	// 	log.Fatalf("failed to notify keygen action: %v", err)
	// }

	// // Subscribe để nhận kết quả keygen
	// pubsub := redisClient.Subscribe(context.Background(), "mpc_keygen")
	// defer pubsub.Close()

	// // Chờ và xử lý kết quả keygen
	// shareData := waitForKeygenResult(pubsub.Channel())

	// // Save share data to file
	// err = utils.SaveToJSON([]byte(shareData), fmt.Sprintf("share_data_%s.txt", sessionID))
	// if err != nil {
	// 	log.Fatalf("failed to save share data: %v", err)
	// }

	// time.Sleep(10 * time.Second)

	// Khởi tạo signing với share data từ keygen
	signCtx, signCancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer signCancel()

	// Tạo message hash để sign (ví dụ)
	msgHash := []byte("hello")

	shareData, err := utils.LoadFromJSON("share_data_fea95c52-dc04-4858-8358-74c19e7149b2.txt")
	if err != nil {
		log.Fatalf("failed to load share data: %v", err)
	}

	// Gửi yêu cầu NotifyAction cho signing
	signRes, err := client.NotifyAction(signCtx, &pb.ActionRequest{
		SessionId: "fea95c52-dc04-4858-8358-74c19e7149b2",
		Parties:   []uint32{1, 2, 3},
		Threshold: 2,
		MsgHash:   msgHash,
		ShareData: shareData, // Truyền share data từ keygen
		Action:    pb.Action_INIT_SIGN,
	})
	if err != nil {
		log.Fatalf("failed to notify signing action: %v", err)
	}

	log.Printf("Sign initialization response: %v\n", signRes)
}

func waitForKeygenResult(ch <-chan *redis.Message) []byte {
	for msg := range ch {
		var result map[string]string

		// Parse JSON
		err := json.Unmarshal([]byte(msg.Payload), &result)
		if err != nil {
			log.Printf("Failed to parse JSON: %v", err)
			continue
		}

		// Decode base64 share và trả về trực tiếp
		encryptedShare, err := base64.StdEncoding.DecodeString(result["share_data"])
		if err != nil {
			log.Printf("Failed to decode base64 share: %v", err)
			continue
		}
		return encryptedShare
	}
	return nil
}
