package alt

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type SequentialWriter struct {
	index    int64
	file     *os.File
	filename string
	writer   *bufio.Writer
	frozen   bool
	sizeof   int
}

const SeqFilePrefix = "big_"

func NewSequentialWriter(f *os.File, sizeof int) (*SequentialWriter, error) {
	if f == nil {
		return nil, errors.New("file is nil")
	}
	if sizeof <= 0 {
		return nil, fmt.Errorf("sizeof cannot be < 1; have: %d", sizeof)
	}

	w := new(SequentialWriter)
	w.file = f
	w.sizeof = sizeof
	w.writer = bufio.NewWriter(f)

	return w, nil
}

func (w *SequentialWriter) WriteAt(p []byte, index int64) (n int, err error) {
	if w.frozen {
		return 0, fmt.Errorf("Writer is frozen: unable to write to")
	}

	if index != w.index {
		return 0, fmt.Errorf("Offset incorrect: have: %d; need: %d", index, w.index)
	}
	index++
	return w.writer.Write(p)
}

func (w *SequentialWriter) Done() error {
	if w.writer == nil {
		return fmt.Errorf("Done(): writer is nil")
	}
	if err := w.writer.Flush(); err != nil {
		return err
	}
	if err := w.file.Close(); err != nil {
		return err
	}
	w.frozen = true
	return nil
}

// return fmt.Errorf("Not implemented")
