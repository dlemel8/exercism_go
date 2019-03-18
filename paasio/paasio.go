package paasio

import (
	"io"
	"sync"
)

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &counter{writer: writer}
}

func NewReadWriteCounter(readWriter io.ReadWriter) ReadWriteCounter {
	return &counter{writer: readWriter, reader: readWriter}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &counter{reader: reader}
}

type ioStats struct {
	bytes int64
	count int
	lock  sync.RWMutex
}

func (i *ioStats) update(n int) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.bytes += int64(n)
	i.count += 1
}

func (i *ioStats) state() (int64, int) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.bytes, i.count
}

type counter struct {
	writer     io.Writer
	writeStats ioStats
	reader     io.Reader
	readStats  ioStats
}

func (c *counter) Write(p []byte) (n int, err error) {
	n, err = c.writer.Write(p)
	c.writeStats.update(n)
	return n, err
}

func (c *counter) WriteCount() (n int64, nops int) {
	return c.writeStats.state()
}

func (c *counter) Read(p []byte) (n int, err error) {
	n, err = c.reader.Read(p)
	c.readStats.update(n)
	return n, err
}

func (c *counter) ReadCount() (n int64, nops int) {
	return c.readStats.state()
}
