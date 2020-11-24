package nodejs

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func versionFromPackageJSONEngines(filePath string) (version string, err error) {
	packagejson, err := GetPackageJSON(filePath)
	if err != nil {
		return
	}

	return packagejson.Engines["node"], nil
}

func versionFromManagerFiles(versionFilePath string) (version string, err error) {
	content, err := ioutil.ReadFile(versionFilePath)
	if err != nil {
		return
	}

	version = strings.Trim(string(content), " \n\r")

	return
}

// VersionFromFile tries to read the Node.js version from the supplied file if supported
func VersionFromFile(versionFilePath string) (version string, err error) {
	switch filepath.Base(versionFilePath) {
	case "package.json":
		return versionFromPackageJSONEngines(versionFilePath)
	case ".nvmrc": // nvm
		fallthrough
	case ".node-version": // nodenv
		fallthrough
	case ".n-node-version": // n
		return versionFromManagerFiles(versionFilePath)
	default:
		err = fmt.Errorf("Unsupported file. Cannot read Node.js version from %s", versionFilePath)
	}

	return
}
