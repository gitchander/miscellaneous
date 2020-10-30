package circular

import (
	"bytes"
	"crypto/sha256"
	"math/rand"
	"testing"
	"time"
)

func TestRandWR(t *testing.T) {
	n := 37
	cb := NewBuffer(n)
	//var cb Buffer

	r := newRandNow()

	var (
		bufWrite = make([]byte, n)
		bufRead  = make([]byte, n)
	)

	var (
		hashWrite = sha256.New()
		hashRead  = sha256.New()
	)

	var err error

	for i := 0; i < 1000; i++ {

		// write
		{
			nw := r.Intn(n + 1)
			randBytes(r, bufWrite[:nw])
			nw, err = cb.Write(bufWrite[:nw])
			if err != nil {
				if err != ErrIsFull {
					t.Fatal(err)
				}
			}
			hashWrite.Write(bufWrite[:nw])
		}

		// read
		{
			nr := r.Intn(n + 1)
			nr, err = cb.Read(bufRead[:nr])
			if err != nil {
				if err != ErrIsEmpty {
					t.Fatal(err)
				}
			}
			hashRead.Write(bufRead[:nr])
		}
	}

	// last read
	if true {
		nr, err := cb.Read(bufRead)
		if err != nil {
			if err != ErrIsEmpty {
				t.Fatal(err)
			}
		}
		hashRead.Write(bufRead[:nr])
	}

	var (
		sumWrite = hashWrite.Sum(nil)
		sumRead  = hashRead.Sum(nil)
	)

	if !bytes.Equal(sumWrite, sumRead) {

		t.Logf("write hash: %X\n", sumWrite)
		t.Logf("read  hash: %X\n", sumRead)

		t.Fatalf("hash sums not equal")
	}
}

func newRandNow() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}

func randBytes(r *rand.Rand, data []byte) {
	for i := range data {
		data[i] = byte(r.Intn(256))
	}
}

func TestWriteIndexes(t *testing.T) {
	samples := []struct {
		Size    int
		WI      int // writeIndex
		RI      int // readIndex
		DataLen int // data length

		writeCount int
		writeIndex int // new writeIndex
	}{
		{
			Size:       5,
			WI:         0,
			RI:         0,
			DataLen:    3,
			writeCount: 3,
			writeIndex: 3,
		},
		{
			Size:       5,
			WI:         3,
			RI:         2,
			DataLen:    3,
			writeCount: 3,
			writeIndex: 1,
		},
	}
	for _, sample := range samples {
		cb := &Buffer{
			buf:        make([]byte, sample.Size),
			writeIndex: sample.WI,
			readIndex:  sample.RI,
		}
		data := make([]byte, sample.DataLen)
		writeCount, err := cb.Write(data)
		if err != nil {
			t.Fatal(err)
		}
		if writeCount != sample.writeCount {
			t.Fatalf("writeCount: (%d != %d)", writeCount, sample.writeCount)
		}
		if cb.writeIndex != sample.writeIndex {
			t.Fatalf("writeIndex: (%d != %d)", cb.writeIndex, sample.writeIndex)
		}
	}
}
