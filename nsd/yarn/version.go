package nodejs

import (
	"fmt"
	"path/filepath"

	NodeJs "github.com/researchgate/nodejs-simple-downloader/nsd/nodejs"
)

func versionFromPackageJSONEngines(filePath string) (version string, err error) {
	packagejson, err := NodeJs.GetPackageJSON(filePath)
	if err != nil {
		return
	}

	return packagejson.Engines["yarn"], nil
}

// VersionFromFile tries to read the yarn version from the supplied file if supported
func VersionFromFile(versionFilePath string) (version string, err error) {

	switch filepath.Base(versionFilePath) {
	case "package.json":
		return versionFromPackageJSONEngines(versionFilePath)
	default:
		err = fmt.Errorf("Unsupported file. Cannot read yarn version from %s", versionFilePath)
	}

	return
}
