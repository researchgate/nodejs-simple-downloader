package checksum

import "testing"

func TestCalculateSHA256(t *testing.T) {
	expectedChecksum := "d9dba1dc545534158a19df00ba6bdf63c9818cd4aa6c373da8a92ba1253a7a02"

	checksum, err := CalculateSHA256("./testdata/randomfile.txt")
	if err != nil {
		t.Errorf("CalculateSHA256() did error %s", err)
	}
	if checksum != expectedChecksum {
		t.Errorf("CalculateSHA256() calculated checksum (%s) does not match expected (%s)", checksum, expectedChecksum)
	}
}
