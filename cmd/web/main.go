package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	myurl "github.com/pratika/assignment3/pkg/download"
	"github.com/pratika/assignment3/pkg/zipper"
)

func main() {

	fmt.Println("Downloading a File from URL")
	urll := "https://filesamples.com/samples/video/mp4/sample_1280x720_surfing_with_audio.mp4"

	var video myurl.Downloader
	video = myurl.NewDownloaderFromUrl()

	r1, err := video.Download(urll)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nDownloading a File from File System")
	filePrefix, err := filepath.Abs("/Users/pratika.tripathi/Desktop/Assignment/Assignment3/sample_1280x720_surfing_with_audio.mp4.mp4")
	if err != nil {
		log.Fatal(err)
	}

	video = myurl.NewDownloaderForFile()
	fileName := filePrefix
	r2, err := video.Download(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var arch zipper.ArchieverInterface
	arch = zipper.NewArchiever()

	arr := []string{"f1.mp4", "f2.mp4"}
	zipR, err := arch.Archieve(arr, r1, r2)
	if err != nil {
		log.Fatal(err)
	}
	zipW, err := os.Create("result.zip")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(zipW, zipR)
	if err != nil {
		log.Fatal(err)
	}
	zipW.Close()
}
