package download

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
)

// WriteCounter counts the number of bytes written to it. It implements to the io.Writer interface
// and we can pass this into io.TeeReader() which will report progress on each write cycle.
type WriteCounter struct {
	Current uint64
	Total   uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Current += uint64(n)
	wc.printProgress()
	return n, nil
}

func (wc WriteCounter) printProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\r%.0f%% (%s/%s)", (float64(wc.Current)/float64(wc.Total))*float64(100), humanize.Bytes(wc.Current), humanize.Bytes(wc.Total))
}

// File will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func File(url string) (tmpFile *os.File, err error) {
	// Create the file
	tmpFile, err = ioutil.TempFile("", "node-simple-downloader.*.tar.gz")
	if err != nil {
		return
	}
	defer tmpFile.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		os.Exit(1)
		// exit if not ok
	}
	defer resp.Body.Close()

	// the Header "Content-Length" will let us know
	// the total file size to download
	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	downloadSize := uint64(size)

	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{Total: downloadSize}
	// Write the body to file
	_, err = io.Copy(tmpFile, io.TeeReader(resp.Body, counter))

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Print("\n")

	return
}
