package myurl

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"strings"
)

type FileSystem struct {
}

func NewDownloaderForFile() *FileSystem {
	return &FileSystem{}
}

func (f *FileSystem) Download(urll string) (r io.Reader, err error) {
	fullURLFile := urll

	// Build fileName from fullPath
	fileURL, err := url.Parse(fullURLFile)

	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")

	fileName := segments[len(segments)-1]

	dest, err := os.Create("New" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	source, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	size, err := io.Copy(dest, source)

	if err != nil {
		log.Fatal(err)
	}

	defer source.Close()
	fmt.Printf("Downloaded a file %s with size %d", fileName, size)

	dest.Seek(0, io.SeekStart)
	return dest, nil
}
