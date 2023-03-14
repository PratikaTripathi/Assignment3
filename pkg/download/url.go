package myurl

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Url struct {
}

func NewDownloaderFromUrl() *Url {
	return &Url{}
}
func (h *Url) Download(urll string) (io.Reader, error) {
	fullURLFile := urll

	// Build fileName from fullPath
	fileURL, err := url.Parse(fullURLFile)

	if err != nil {
		log.Fatal(err)
	}

	path := fileURL.Path
	segments := strings.Split(path, "/")
	fmt.Println(segments)

	fileName := segments[len(segments)-1]
	//fmt.Println(fileName)
	dest, err := os.OpenFile(fileName+".mp4", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}

	// Put content on file
	resp, err := client.Get(fullURLFile)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	size, err := io.Copy(dest, resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Downloaded a file %s with size %d", fileName, size)

	dest.Seek(0, io.SeekStart)
	return dest, nil
}
