package zipper

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

type archieve struct {
}

func NewArchiever() *archieve {
	return &archieve{}
}

func (a *archieve) Archieve(fileNames []string, readers ...io.Reader) (outR io.Reader, err error) {
	filePtr, err := os.Create("archive.zip")
	if err != nil {
		return nil, fmt.Errorf("error in creating file archive.zip due to %s", err.Error())
	}

	zipW := zip.NewWriter(filePtr)
	defer zipW.Close()
	defer zipW.Flush()

	for i, r := range readers {
		file := fileNames[i]
		w, err := zipW.Create(file)
		if err != nil {
			return nil, fmt.Errorf("error in creating file inside archive.zip due to %s", err.Error())
		}
		_, err = io.Copy(w, r)
		if err != nil {
			return nil, fmt.Errorf("error in copying file inside archive.zip due to %s", err.Error())
		}
	}

	_, err = filePtr.Seek(0, io.SeekStart)
	return filePtr, err
}
