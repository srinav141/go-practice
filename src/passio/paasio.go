package paasio

import "io"

type ReadWriteCountImpl struct {
}

func (w *ReadWriteCountImpl) Write(p []byte) (int, error) {
	return 0, nil
}

func (e *ReadWriteCountImpl) WriteCount() (int64, int) {
	return 0, 0
}

func NewWriteCounter(w io.Writer) WriteCounter {

	return &ReadWriteCountImpl{}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return nil

}

func NewReadCounter(r io.Reader) ReadCounter {
	return nil
}
