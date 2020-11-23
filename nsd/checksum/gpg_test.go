package checksum

import "testing"

func TestVerifyGPGSignature(t *testing.T) {
	err := VerifyGPG("./testdata/gpg/file.txt.asc", "./testdata/gpg/pubkey.gpg", "./testdata/gpg/file.txt")
	if err != nil {
		t.Errorf("TestVerifyGPGSignature() did error %s", err)
	}
}
