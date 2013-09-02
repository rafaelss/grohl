package grohl

import (
	"io"
	"sync"
)

type MutexWriter struct {
	Writer io.Writer
	mu     sync.Mutex
}

func SyncWriter(writer io.Writer) *MutexWriter {
	return &MutexWriter{Writer: writer}
}

func (w *MutexWriter) Write(line []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.Writer.Write(line)
}
