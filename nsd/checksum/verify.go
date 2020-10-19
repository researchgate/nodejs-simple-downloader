package checksum

import (
	"bufio"
	"os"
	"strings"
)

// Verify does check if the provided checksum for the fileName does match the checksum
// from the nodejs checksum file
func Verify(checksum string, fileName string, nodeChecksumFilePath string) (verified bool, err error) {
	verified = false

	file, err := os.Open(nodeChecksumFilePath)
	if err != nil {
		return
	}
	defer file.Close()

	realFileName := fileName
	// We create TempFiles in the format `<random>#<filename>`
	if strings.Contains(fileName, "#") {
		realFileName = strings.Split(fileName, "#")[1]
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineChecksum, lineFileName := parseChecksumLine(scanner.Text())
		if realFileName == lineFileName {
			verified = lineChecksum == checksum
			return
		}
	}
	err = scanner.Err()

	return
}

func parseChecksumLine(line string) (string, string) {
	parts := strings.Split(line, "  ")

	return parts[0], parts[1]
}
