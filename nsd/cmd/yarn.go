package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/mholt/archiver/v3"
	Checksum "github.com/researchgate/nodejs-simple-downloader/nsd/checksum"
	Download "github.com/researchgate/nodejs-simple-downloader/nsd/download"

	"github.com/spf13/cobra"
)

var (
	yarnVersion string
	yarnCommand = &cobra.Command{
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

			// https://github.com/yarnpkg/yarn/releases/download/v'.$version.'/yarn-v'.$version.'.tar.gz

			err = prepareYarnFlags()
			if err != nil {
				return
			}

			yarnURL := fmt.Sprintf("https://github.com/yarnpkg/yarn/releases/download/v%s/yarn-v%s.tar.gz", yarnVersion, yarnVersion)
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

			tar := archiver.NewTarGz()
			tar.StripComponents = 1
			err = tar.Unarchive(yarnFilePath, downloadPath)

			if err != nil {
				return
			}

			return
		},
	}
)

func prepareYarnFlags() (err error) {
	if yarnVersion == "" {
		return errors.New("No version specified")
	}

	return
}

func init() {
	yarnCommand.Flags().StringVarP(&yarnVersion, "version", "v", "", "Which version to install")
	rootCmd.AddCommand(yarnCommand)
}
