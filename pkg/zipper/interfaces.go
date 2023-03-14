package zipper

import "io"

type ArchieverInterface interface {
	Archieve(names []string, readers ...io.Reader) (outR io.Reader, err error)
}
