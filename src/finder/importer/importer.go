package importer

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const Storage = "/tmp/"

func CollectAllPastes(pastes map[string]string) (map[string]string, error) {

	log.Printf("Trying to import all pastes")

	var (
		err error
		destination string
	)

	files := make(map[string]string, 0)
	for k, v := range pastes {

		if startsWithHttp(v) == true {

			destination = Storage + k + ".md"
			log.Printf("paste %s - Trying download it from source ( %s )", k, v)
			err = downloadURLToFile (v, destination)

		} else {

			log.Printf("%s - seems to be a local file. Import it as is for paste %s", v, k)
			destination = v
		}

		files[k] = destination
	}
	return files, err
	
}
func startsWithHttp(line string) bool {
	return strings.HasPrefix(line, "http")
}

// DownloadURLToFile downloads Body response from URL to a file.
func downloadURLToFile(urlStr string, fileName string) error {
	var f *os.File

	resp, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)

	return err
}

