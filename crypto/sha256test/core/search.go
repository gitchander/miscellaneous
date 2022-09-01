package core

import (
	"bytes"
	"context"
	"sync"
)

type Result struct {
	Blocks []GenesisBlock `json:"blocks"`
}

// genesis block
// origin block !!!
type GenesisBlock struct {
	Data []byte `json:"data"`
	Pos  [2]int `json:"pos"`
}

func Search(ctx context.Context, data []byte, sampleSize int, originSize int) (*Result, error) {

	child, cancel := context.WithCancel(ctx)
	defer cancel()

	samples := splitBytes(data, sampleSize)

	var (
		guard sync.Mutex

		gbs     = make([]GenesisBlock, len(samples))
		errOnce error
	)

	var wg sync.WaitGroup
	wg.Add(len(samples))

	for i, sample := range samples {
		go func(index int, bs []byte) {
			defer wg.Done()
			gb, err := blockSearch(child, bs, originSize)
			if err != nil {
				guard.Lock()
				if errOnce == nil {
					errOnce = err
				}
				guard.Unlock()
				cancel()
				return
			}
			guard.Lock()
			gbs[index] = *gb
			guard.Unlock()
		}(i, sample)
	}

	wg.Wait()

	if errOnce != nil {
		return nil, errOnce
	}

	r := &Result{
		Blocks: gbs,
	}

	return r, nil
}

func blockSearch(ctx context.Context, sample []byte, originSize int) (*GenesisBlock, error) {

	r := nextRandom()

	data := make([]byte, originSize)

	h := newHash()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		randFillBytes(r, data)

		h.Reset()
		h.Write(data)
		sum := h.Sum(nil)

		index := bytes.Index(sum, sample)
		if index != -1 {
			gb := &GenesisBlock{
				Data: data,
				Pos: [2]int{
					index,
					index + len(sample),
				},
			}
			return gb, nil
		}
	}
}
