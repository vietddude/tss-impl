package party

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

// Test tạo ví Ethereum
func TestGenerateWallet(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	assert.NoError(t, err, "Lỗi khi tạo private key")

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	t.Logf("🟢 Địa chỉ ví: %s", address.Hex())

	assert.NotEmpty(t, address.Hex(), "Địa chỉ ví không hợp lệ")
}

// Test tạo và ký giao dịch
func TestSignTransaction(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	assert.NoError(t, err, "Lỗi khi tạo private key")

	toAddress := common.HexToAddress("0x187d26A192Ea5848e4bbB8A550CF63d83220D360")
	value := big.NewInt(10000000000000000) // 0.01 ETH
	gasLimit := uint64(21000)              // Gas limit cho 1 txn chuyển ETH
	gasPrice := big.NewInt(5000000000)     // 5 Gwei
	nonce := uint64(0)                     // Nonce tạm thời (thực tế phải lấy từ blockchain)
	chainID := big.NewInt(11155111)        // Sepolia Testnet Chain ID

	// 🏗 Tạo giao dịch
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// ✍️ Ký giao dịch
	signedTx, err := SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	assert.NoError(t, err, "Lỗi khi ký giao dịch")

	// 📤 In thông tin giao dịc
	t.Logf("📤 Giao dịch đã ký: %s", signedTx.Hash().Hex())
}

// Test gửi giao dịch lên blockchain (yêu cầu node Ethereum)
func TestSendTransaction(t *testing.T) {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/YOUR_INFURA_KEY")
	assert.NoError(t, err, "Không thể kết nối với Infura")

	privateKey, err := crypto.GenerateKey()
	assert.NoError(t, err, "Lỗi khi tạo private key")

	toAddress := common.HexToAddress("0x187d26A192Ea5848e4bbB8A550CF63d83220D360")
	value := big.NewInt(10000000000000000)
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(5000000000)
	nonce := uint64(0) // Lấy nonce thật từ blockchain
	chainID := big.NewInt(11155111)

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	assert.NoError(t, err, "Lỗi khi ký giao dịch")

	err = client.SendTransaction(context.Background(), signedTx)
	assert.NoError(t, err, "Lỗi khi gửi giao dịch")

	t.Logf("📤 Giao dịch đã gửi: %s", signedTx.Hash().Hex())
}

func SignTx(tx *types.Transaction, s types.Signer, prv *ecdsa.PrivateKey) (*types.Transaction, error) {
	h := s.Hash(tx)                    // Băm giao dịch
	sig, err := crypto.Sign(h[:], prv) // Ký giao dịch bằng private key
	if err != nil {
		return nil, err
	}

	// 🛠 Tách R, S, V từ chữ ký
	r := new(big.Int).SetBytes(sig[:32])
	sVal := new(big.Int).SetBytes(sig[32:64])
	v := sig[64] // V có giá trị 27 hoặc 28

	fmt.Printf("Signature: %x\n", sig)
	fmt.Printf("R: %x\n", r)
	fmt.Printf("S: %x\n", sVal)
	fmt.Printf("V: %x (%d)\n", v, v)

	// 🔄 Chuẩn hóa V về 0 hoặc 1 để recover public key
	// normalizedV := v - 27
	recoverableSig := append(sig[:64], v)

	// 🔑 Recover public key từ signature
	recoveredPubKey, err := crypto.SigToPub(h[:], recoverableSig)
	if err != nil {
		return nil, fmt.Errorf("lỗi recover public key: %w", err)
	}
	recoveredPubKeyBytes := crypto.FromECDSAPub(recoveredPubKey)
	fmt.Printf("Recovered Public Key: %x\n", recoveredPubKeyBytes)

	// 🔍 Lấy public key gốc từ private key
	originalPubKey := crypto.FromECDSAPub(&prv.PublicKey)

	// ✅ So sánh public key gốc và recovered
	if bytes.Equal(originalPubKey, recoveredPubKeyBytes) {
		fmt.Println("✅ Public key recovered chính xác!")
	} else {
		fmt.Println("❌ Public key recovered sai!")
	}

	return tx.WithSignature(s, sig)
}
