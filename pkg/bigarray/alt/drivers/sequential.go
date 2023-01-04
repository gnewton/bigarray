package drivers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Sequential struct {
	index  int64
	file   *os.File
	sizeof int
}

type SequentialWriter struct {
	Sequential
	writer *bufio.Writer
	frozen bool
}

type SequentialReader struct {
	Sequential
	reader *bufio.Reader
}

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
	w.index++
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

func NewSequentialReader(f *os.File, sizeof int) (*SequentialReader, error) {
	if f == nil {
		return nil, errors.New("file is nil")
	}
	if sizeof <= 0 {
		return nil, fmt.Errorf("sizeof cannot be < 1; have: %d", sizeof)
	}

	r := new(SequentialReader)
	r.file = f
	r.sizeof = sizeof
	r.reader = bufio.NewReader(f)

	return r, nil
}

func (r *SequentialReader) ReadAt(p []byte, index int64) (n int, err error) {
	if index != r.index {
		return 0, fmt.Errorf("Offset incorrect: have: %d; need: %d", index, r.index)
	}
	r.index++
	return r.reader.Read(p)
}
func (w *SequentialReader) Done() error {
	return w.file.Close()
}
