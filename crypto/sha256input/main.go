package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"math/rand"
	"runtime"
	"time"
)

const (
	bytesPerUint64 = 8
	bitsPerByte    = 8
)

func newRandNow() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

type Result struct {
	Sample []byte
	Data   []byte
	Pos    []int
}

func fillBytes(r *rand.Rand, data []byte) {
	for len(data) >= bytesPerUint64 {
		x := r.Uint64()
		binary.BigEndian.PutUint64(data, x)
		data = data[bytesPerUint64:]
	}
	if len(data) > 0 {
		x := r.Uint64()
		for len(data) > 0 {
			data[0] = byte(x)
			data = data[1:]
			x >>= bitsPerByte
		}
	}
}

func findHash(h hash.Hash, samples [][]byte, size int) *Result {

	for _, sample := range samples {
		if len(sample) > h.Size() {
			panic("invalid sample size")
		}
	}

	r := newRandNow()

	data := make([]byte, size)

	for {
		// for i := range data {
		// 	data[i] = byte(r.Intn(256))
		// }

		fillBytes(r, data)

		h.Reset()
		h.Write(data)
		sum := h.Sum(nil)

		for _, sample := range samples {
			//fmt.Printf("sample: %s\n", sample)

			// if bytes.Equal(sum[:len(sample)], sample) {
			index := bytes.Index(sum[:], sample)
			if index != -1 {

				fmt.Printf("index %d, %s\n", index, string(sum[:]))

				return &Result{
					Sample: sample,
					Data:   data,
					Pos: []int{
						index,
						index + len(sample),
					},
				}
			}
		}
	}
}

type hashPiece struct {
	data []byte
	pos  []int
}

var (
	hw1 = []hashPiece{
		{
			data: hexDecode("9e301d67"), // 'hell'
			pos:  []int{25, 29},
		},
		{
			data: hexDecode("1682b8ef"), // 'o wo'
			pos:  []int{16, 20},
		},
		{
			data: hexDecode("3f4eece4"), // 'rld!'
			pos:  []int{3, 7},
		},
	}

	hw2 = []hashPiece{
		{
			data: hexDecode("d161ebea71a03dbe4d68"),
			pos:  []int{6, 10},
		},
		{
			data: hexDecode("5e9440f64324da7aef02"),
			pos:  []int{14, 18},
		},
		{
			data: hexDecode("7a95e9b7929fb894e224"),
			pos:  []int{7, 11},
		},
	}

	hw3 = []hashPiece{
		{
			data: hexDecode("e7d319c5fc3cde57"), // 'hello'
			pos:  []int{13, 18},
		},
		{
			data: hexDecode("773ad03ad7a28a18"), // ' '
			pos:  []int{23, 24},
		},
		{
			data: hexDecode("f9b1b6643ad08f91"), // 'world'
			pos:  []int{9, 14},
		},
		{
			data: hexDecode("6ec0112796dbc1ed"), // '!'
			pos:  []int{17, 18},
		},
	}
)

func greet() string {
	samples := hw3
	var buf bytes.Buffer
	for _, sample := range samples {
		var (
			sum = sha256.Sum256(sample.data)
			pos = sample.pos
		)
		buf.Write(sum[pos[0]:pos[1]])
	}
	return buf.String()
}

func hexDecode(s string) []byte {
	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func run() {

	var size int
	var strValue string

	flag.IntVar(&size, "size", 8, "size of input value")
	flag.StringVar(&strValue, "sample", "", "find sample")

	flag.Parse()

	samples := [][]byte{
		[]byte(strValue),
	}

	//-----------------------------------------------
	// size := 4
	// samples := [][]byte{
	// 	[]byte("hell"),
	// 	[]byte("o wo"),
	// 	[]byte("rld!"),
	// }
	//-----------------------------------------------
	// size := 8
	// samples := [][]byte{
	// 	[]byte("hello"),
	// 	[]byte("world"),
	// }
	//-----------------------------------------------
	// size := 10
	// samples := [][]byte{
	// 	// []byte("hello "),
	// 	// []byte("world!"),
	// 	[]byte("chander"),
	// }
	//-----------------------------------------------

	numCPU := runtime.NumCPU()

	result := make(chan *Result)

	for i := 0; i < numCPU; i++ {
		go func() {
			h := sha256.New()
			result <- findHash(h, samples, size)
		}()
	}

	res := <-result

	fmt.Println("result:")
	fmt.Printf("sample: %s\n", res.Sample)
	fmt.Printf("data: %x\n", res.Data)
	fmt.Printf("pos: %v\n", res.Pos)
}

func main() {
	start := time.Now()
	run()
	fmt.Println(time.Since(start))
}
