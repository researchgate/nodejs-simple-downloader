package checksum

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

// CalculateSHA256 does calculate a sha-256 for the provided file
func CalculateSHA256(filePath string) (checksum string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	h := sha256.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return
	}

	checksum = hex.EncodeToString(h.Sum(nil))
	return
}
