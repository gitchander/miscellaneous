package circular

import (
	"errors"
	"io"
)

// (RI = WI):
// +---------+--+--+--+--+--+--+--+--+--+--+
// | indexes |  |  |RW|  |  |  |  |  |  |  |
// +---------+--+--+--+--+--+--+--+--+--+--+
// |  data   |--|--|--|--|--|--|--|--|--|--|
// +---------+--+--+--+--+--+--+--+--+--+--+

// (RI < WI):
// +---------+--+--+--+--+--+--+--+--+--+--+
// | indexes |  |  |RI|  |  |  |  |WI|  |  |
// +---------+--+--+--+--+--+--+--+--+--+--+
// |  data   |--|--|XX|XX|XX|XX|XX|--|--|--|
// +---------+--+--+--+--+--+--+--+--+--+--+

// (RI > WI):
// +---------+--+--+--+--+--+--+--+--+--+--+
// | indexes |  |  |WI|  |  |  |  |RI|  |  |
// +---------+--+--+--+--+--+--+--+--+--+--+
// |  data   |XX|XX|--|--|--|--|--|XX|XX|XX|
// +---------+--+--+--+--+--+--+--+--+--+--+

var (
	ErrIsFull  = errors.New("circular buffer is full")
	ErrIsEmpty = errors.New("circular buffer is empty")
)

// circular buffer
type Buffer struct {
	buf        []byte
	writeIndex int
	readIndex  int
}

var _ io.ReadWriter = &Buffer{}

func NewBuffer(size int) *Buffer {
	if size < 0 {
		size = 0
	}
	return &Buffer{
		buf: make([]byte, size+1),
	}
}

func (b *Buffer) Cap() int {
	n := len(b.buf)
	if n > 0 {
		n--
	}
	return n
}

func (b *Buffer) Empty() bool {
	return b.writeIndex == b.readIndex
}

func (b *Buffer) Full() bool {
	n := len(b.buf)
	if n == 0 {
		return false
	}
	return (b.writeIndex+1)%n == b.readIndex
}

func (b *Buffer) Reset() {
	b.writeIndex = 0
	b.readIndex = 0
}

func (b *Buffer) Write(data []byte) (n int, err error) {

	if b.writeIndex >= b.readIndex {
		k := len(b.buf) - b.writeIndex
		if (len(b.buf) > 0) && (b.readIndex == 0) {
			k--
		}
		n += b.writeSolid(data[n:], k)
	}

	if b.writeIndex < b.readIndex {
		k := b.readIndex - b.writeIndex - 1
		n += b.writeSolid(data[n:], k)
	}

	if len(data[n:]) > 0 {
		err = ErrIsFull
	}

	return
}

func (b *Buffer) writeSolid(data []byte, k int) int {
	k = minInt(k, len(data))
	if k > 0 {
		copy(b.buf[b.writeIndex:], data[:k])
		b.writeIndex = (b.writeIndex + k) % len(b.buf)
	}
	return k
}

func (b *Buffer) Read(data []byte) (n int, err error) {

	if b.Empty() {
		return 0, ErrIsEmpty
	}

	if b.readIndex > b.writeIndex {
		k := len(b.buf) - b.readIndex
		n += b.readSolid(data[n:], k)
	}

	if b.readIndex < b.writeIndex {
		k := b.writeIndex - b.readIndex
		n += b.readSolid(data[n:], k)
	}

	return
}

func (b *Buffer) readSolid(data []byte, k int) int {
	k = minInt(k, len(data))
	if k > 0 {
		copy(data[:k], b.buf[b.readIndex:])
		b.readIndex = (b.readIndex + k) % len(b.buf)
	}
	return k
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
