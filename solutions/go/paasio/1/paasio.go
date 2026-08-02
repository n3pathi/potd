package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.
type readCounter struct {
	mu        sync.Mutex
	delegate  io.Reader
	bytesRead int
	noOps     int
}
type writeCounter struct {
	mu           sync.Mutex
	delegate     io.Writer
	bytesWritten int
	noOps        int
}
type readWriteCounter struct {
	ReadCounter
	WriteCounter
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{delegate: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{delegate: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		ReadCounter:  NewReadCounter(readwriter),
		WriteCounter: NewWriteCounter(readwriter),
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	n, err := rc.delegate.Read(p)
	rc.mu.Lock()
	defer rc.mu.Unlock()
	rc.bytesRead += n
	rc.noOps++
	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	return int64(rc.bytesRead), rc.noOps
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n, err := wc.delegate.Write(p)
	wc.mu.Lock()
	defer wc.mu.Unlock()
	wc.bytesWritten += n
	wc.noOps++
	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	return int64(wc.bytesWritten), wc.noOps
}
