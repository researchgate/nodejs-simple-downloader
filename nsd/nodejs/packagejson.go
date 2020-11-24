package nodejs

import (
	"encoding/json"
	"io/ioutil"
)

// PackageJSONEngines holds the data from package.json
type PackageJSONEngines struct {
	Engines map[string]string `json:"engines"`
}

// GetPackageJSON returns a struct with data from package.json
func GetPackageJSON(filePath string) (packagejson *PackageJSONEngines, err error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	err = json.Unmarshal(content, &packagejson)

	return
}
