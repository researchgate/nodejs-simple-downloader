package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mholt/archiver/v3"
	Checksum "github.com/researchgate/nodejs-simple-downloader/nsd/checksum"
	Download "github.com/researchgate/nodejs-simple-downloader/nsd/download"
	NodeJs "github.com/researchgate/nodejs-simple-downloader/nsd/nodejs"

	"github.com/spf13/cobra"
)

var (
	version         string
	fromFile        string
	downloadCommand = &cobra.Command{
		Use:   "download [path]",
		Short: "Download nodejs to specific folder",
		Long:  "",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a path argument")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			downloadPath := args[0]

			err = prepareFlags()
			if err != nil {
				return
			}

			nodeURL := fmt.Sprintf(string(NodeJs.CurrentURL)+"node-v%s-%s.%s", version, version, NodeJs.CurrentArch, NodeJs.CurrentExtension)
			nodeFilePath, err := Download.File(nodeURL)
			if err != nil {
				return
			}
			defer os.Remove(nodeFilePath)

			checksumURL := fmt.Sprintf(string(NodeJs.CurrentURL)+"SHASUMS256.txt", version)
			checkusmFilePath, err := Download.File(checksumURL)
			if err != nil {
				return
			}
			defer os.Remove(checkusmFilePath)

			checksum, err := Checksum.CalculateSHA256(nodeFilePath)
			if err != nil {
				return
			}

			verified, err := Checksum.Verify(checksum, path.Base(nodeFilePath), checkusmFilePath)
			if err != nil {
				return
			}
			if !verified {
				return errors.New("Checksum mismatch. Aborting")
			}

			err = os.RemoveAll(downloadPath)
			if err != nil {
				return
			}

			switch NodeJs.CurrentExtension {
			case NodeJs.Zip:
				zip := archiver.NewZip()
				zip.StripComponents = 1
				err = zip.Unarchive(nodeFilePath, downloadPath)
			case NodeJs.TarXz:
				tar := archiver.NewTarXz()
				tar.StripComponents = 1
				err = tar.Unarchive(nodeFilePath, downloadPath)
			case NodeJs.TarGz:
				tar := archiver.NewTarGz()
				tar.StripComponents = 1
				err = tar.Unarchive(nodeFilePath, downloadPath)
			default:
				return errors.New("Invalid archive format. Aborting")
			}

			if err != nil {
				return
			}

			return
		},
	}
)

func prepareFlags() (err error) {
	if version != "" && fromFile != "" {
		return errors.New("cannot figure out which version to install. Please only specify one of --version or --from-file")
	}
	versionSpecified := true
	if version == "" && fromFile == "" {
		versionSpecified = false
		fromFile = ".nvmrc"
	}

	if fromFile != "" {
		fromFile, err = filepath.Abs(fromFile)
		if err != nil {
			return
		}

		content, err := ioutil.ReadFile(fromFile)
		if err != nil {
			if !versionSpecified {
				return errors.New("No version specified and could not find any version file in the current directory")
			}
			return err
		}

		version = strings.Trim(string(content), " \n\r")
	}

	return
}

func init() {
	downloadCommand.Flags().StringVarP(&version, "version", "v", "", "Which version to install")
	downloadCommand.Flags().StringVarP(&fromFile, "from-file", "r", "", "Reads the version to be installed from a file. Either specify the filename or if empty it will try to read from .nvmrc file.")
	downloadCommand.MarkFlagFilename("from-file")
	rootCmd.AddCommand(downloadCommand)
}
