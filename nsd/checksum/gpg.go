package checksum

import (
	"fmt"
	"golang.org/x/crypto/openpgp"
	"os"
)

// VerifyGPG is cool
func VerifyGPG(signatureFilePath string, keyFilePath string, tarFilePath string) (err error) {
	keyFile, err := os.Open(keyFilePath)
	if err != nil {
		return
	}
	defer keyFile.Close()

	signatureFile, err := os.Open(signatureFilePath)
	if err != nil {
		return
	}
	defer signatureFile.Close()

	tarFile, err := os.Open(tarFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tarFile.Close()

	keyring, err := openpgp.ReadArmoredKeyRing(keyFile)
	if err != nil {
		return
	}
	entity, err := openpgp.CheckArmoredDetachedSignature(keyring, tarFile, signatureFile)
	if err != nil {
		return
	}

	fmt.Println(entity)

	return
}
