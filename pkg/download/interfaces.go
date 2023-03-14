package myurl

import "io"

type Downloader interface {
	Download(urll string) (io.Reader, error)
}
