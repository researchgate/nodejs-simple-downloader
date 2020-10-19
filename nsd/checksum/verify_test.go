package checksum

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	expectedChecksum := "6e878d5e7aedaffb16de27ed65ee8d8351282c146caf8aa3ef726fded26226c5"
	expectedFile := "1234#node-v12.19.0-linux-x64.tar.xz"
	verified, err := Verify(expectedChecksum, expectedFile, "./testdata/SHASUMS256.txt")
	if err != nil {
		t.Errorf("Verify() did error %s", err)
	}
	if !verified {
		t.Errorf("Verify() returned false, expected true")
	}
}

func TestVerifyWithNoHashFilename(t *testing.T) {
	expectedChecksum := "6e878d5e7aedaffb16de27ed65ee8d8351282c146caf8aa3ef726fded26226c5"
	expectedFile := "node-v12.19.0-linux-x64.tar.xz"
	verified, err := Verify(expectedChecksum, expectedFile, "./testdata/SHASUMS256.txt")
	if err != nil {
		t.Errorf("Verify() did error %s", err)
	}
	if !verified {
		t.Errorf("Verify() returned false, expected true")
	}
}

func TestVerifyFailsOnInvalidchecksum(t *testing.T) {
	expectedChecksum := "5bd51bcc2017a1aca716e7b07cac5ed3e5ae0b475815f1eed084232cdf598004"
	expectedFile := "1234#node-v12.19.0-linux-x64.tar.xz"
	verified, err := Verify(expectedChecksum, expectedFile, "./testdata/SHASUMS256.txt")
	if err != nil {
		t.Errorf("Verify() did error %s", err)
	}
	if verified {
		t.Errorf("Verify() returned true, expected false")
	}
}

func TestVerifyFailsOnFilenameNotFound(t *testing.T) {
	expectedChecksum := "6e878d5e7aedaffb16de27ed65ee8d8351282c146caf8aa3ef726fded26226c5"
	expectedFile := "1234#node-v12.20.0-linux-x64.tar.xz"
	verified, err := Verify(expectedChecksum, expectedFile, "./testdata/SHASUMS256.txt")
	if err != nil {
		t.Errorf("Verify() did error %s", err)
	}
	if verified {
		t.Errorf("Verify() returned true, expected false")
	}
}

func TestParseChecksumLine(t *testing.T) {
	expectedChecksum := "256ce45b2aad4f4d7da6e282f94f1c8cfdef20cd0c4e346c9a158116fc944825"
	expectedFile := "node-v12.19.0-aix-ppc64.tar.gz"
	line := fmt.Sprintf("%s  %s", expectedChecksum, expectedFile)
	checksum, file := parseChecksumLine(line)
	if checksum != expectedChecksum {
		t.Errorf("parseChecksumLine(\"%s\") did not return checksum %s", line, expectedChecksum)
	}
	if file != expectedFile {
		t.Errorf("parseChecksumLine(\"%s\") did not return file %s", line, expectedFile)
	}
}
