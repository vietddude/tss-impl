package utils

import (
	"bytes"
	"encoding/base64"
	"io"

	"github.com/klauspost/compress/zstd"
)

// Compress and Base64 encode the encrypted share
func CompressAndEncode(data []byte) (string, error) {
	var b bytes.Buffer
	w, err := zstd.NewWriter(&b)
	if err != nil {
		return "", err
	}
	_, err = w.Write(data)
	w.Close()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

func DecodeAndDecompress(encoded string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	r, err := zstd.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return io.ReadAll(r)
}
