package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/mholt/archiver/v3"
	Checksum "github.com/researchgate/nodejs-simple-downloader/nsd/checksum"
	Download "github.com/researchgate/nodejs-simple-downloader/nsd/download"
	Yarn "github.com/researchgate/nodejs-simple-downloader/nsd/yarn"

	"github.com/spf13/cobra"
)

var (
	yarnVersion         string
	yarnVersionfromFile string
	singleFile          string
	yarnCommand         = &cobra.Command{
		Use:   "yarn [path]",
		Short: "Download yarn to specific folder",
		Long:  "",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a path argument")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			downloadPath := args[0]

			err = prepareYarnFlags()
			if err != nil {
				return
			}

			yarnURL := fmt.Sprintf("https://github.com/yarnpkg/yarn/releases/download/v%s/yarn-", yarnVersion)
			if singleFile != "" {
				yarnURL = fmt.Sprintf(yarnURL+"%s.js", yarnVersion)
			} else {
				yarnURL = fmt.Sprintf(yarnURL+"v%s.tar.gz", yarnVersion)
			}
			yarnFilePath, err := Download.File(yarnURL)
			if err != nil {
				return
			}
			defer os.Remove(yarnFilePath)

			signatureURL := yarnURL + ".asc"
			signatureFilePath, err := Download.File(signatureURL)
			if err != nil {
				return
			}
			defer os.Remove(signatureFilePath)

			keyURL := "https://dl.yarnpkg.com/debian/pubkey.gpg"
			keyFilePath, err := Download.File(keyURL)
			if err != nil {
				return
			}
			defer os.Remove(keyFilePath)

			err = Checksum.VerifyGPG(signatureFilePath, keyFilePath, yarnFilePath)
			if err != nil {
				return
			}

			err = os.RemoveAll(downloadPath)
			if err != nil {
				return
			}

			if singleFile != "" {
				err = os.MkdirAll(downloadPath, 0755)
				if err != nil {
					return
				}

				destinationFilePath := path.Join(downloadPath, singleFile)
				err = os.Rename(yarnFilePath, destinationFilePath)
				if err != nil {
					return err
				}
				err = os.Chmod(destinationFilePath, 0755)
			} else {
				tar := archiver.NewTarGz()
				tar.StripComponents = 1
				err = tar.Unarchive(yarnFilePath, downloadPath)
			}

			if err != nil {
				return
			}

			return
		},
	}
)

func prepareYarnFlags() (err error) {
	if (yarnVersion != "" && yarnVersionfromFile != "") || (yarnVersion == "" && yarnVersionfromFile == "") {
		return errors.New("cannot figure out which version to install. Please specify one of --version or --from-file")
	}

	if yarnVersionfromFile != "" {
		yarnVersion, err = Yarn.VersionFromFile(yarnVersionfromFile)
	}

	return
}

func init() {
	yarnCommand.Flags().StringVarP(&yarnVersion, "version", "v", "", "Which version to install")
	yarnCommand.Flags().StringVarP(&singleFile, "single-file", "s", "", "Download only the single file distribution from yarn and save with the supplied name in the download path")
	yarnCommand.Flags().StringVarP(&yarnVersionfromFile, "version-from-file", "f", "", "Reads the version to be installed from a file. Supported is only package.json")
	rootCmd.AddCommand(yarnCommand)
}
