package utils

import (
	"hash/maphash"
)

type BytesHasher struct {
	mh     *maphash.Hash
	hashes map[uint64]struct{} // hash table
}

func NewBytesHasher() *BytesHasher {
	return &BytesHasher{
		mh:     new(maphash.Hash),
		hashes: make(map[uint64]struct{}),
	}
}

func (p *BytesHasher) calcSum(data []byte) uint64 {
	p.mh.Reset()
	p.mh.Write(data)
	return p.mh.Sum64()
}

func (p *BytesHasher) AddIfNotExist(data []byte) bool {
	var (
		sum   = p.calcSum(data)
		_, ok = p.hashes[sum]
	)
	if !ok {
		p.hashes[sum] = struct{}{}
		return true
	}
	return false
}
