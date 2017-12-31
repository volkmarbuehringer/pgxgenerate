package writer

import (
	"bufio"
	"os"
	"path/filepath"
)

type Writer struct {
	f       *os.File
	dirname string
}

func (t *Writer) Close() error {
	if t.f != nil {
		return t.f.Close()
	}
	return nil
}

func (t *Writer) Create(n string) (*bufio.Writer, error) {
	var err error
	t.f, err = os.Create(filepath.Join(t.dirname, n))
	if err == nil {
		return bufio.NewWriter(t), nil
	} else {
		return nil, err
	}

}

func Init(dirname string) *Writer {
	t := Writer{dirname: dirname}
	return &t

}
func (t *Writer) Write(p []byte) (n int, err error) {
	l, err := t.f.Write(p)
	if err != nil {
		panic(err)
	}
	if len(p) != l {
		panic(err)
	}
	return l, err
}
