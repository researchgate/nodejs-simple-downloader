package download

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

func TestDownloadTgz(t *testing.T) {
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./testdata/")))

	server := httptest.NewServer(router)
	defer server.Close()

	filePath, err := File(server.URL + "/basic.tar.gz")
	if filePath == "" {
		t.Errorf("File() returned no path")
	}
	defer os.Remove(filePath)

	if err != nil {
		t.Errorf("File() did error %s", err)
	}

	info, err := os.Stat(filePath)
	if os.IsNotExist(err) || info.IsDir() {
		t.Errorf("File() returned path to file, but file does not exist")
	}
}

func TestDownloadZip(t *testing.T) {
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./testdata/")))

	server := httptest.NewServer(router)
	defer server.Close()

	filePath, err := File(server.URL + "/basic.zip")
	if filePath == "" {
		t.Errorf("File() returned no path")
	}
	defer os.Remove(filePath)

	if err != nil {
		t.Errorf("File() did error %s", err)
	}

	info, err := os.Stat(filePath)
	if os.IsNotExist(err) || info.IsDir() {
		t.Errorf("File() returned path to file, but file does not exist")
	}
}

func TestDownloadInvalidFile(t *testing.T) {
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./testdata/")))

	server := httptest.NewServer(router)
	defer server.Close()

	_, err := File(server.URL + "/notexist.zip")
	if err == nil {
		t.Errorf("File() did not error")
	}
}
