package varint

import (
	"bytes"
	"encoding/hex"
	"math"
	"testing"
)

type testSample struct {
	Value uint64
	Bytes []byte
}

func makeTestSample(v uint64, data []byte) testSample {
	return testSample{
		Value: v,
		Bytes: data,
	}
}

func makeTestSampleHex(v uint64, s string) testSample {
	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return makeTestSample(v, data)
}

func TestVarint(t *testing.T) {
	samples := []testSample{
		makeTestSampleHex(0, "00"),
		makeTestSampleHex(1, "01"),
		makeTestSampleHex(252, "fc"),
		makeTestSampleHex(253, "fd00fd"),
		makeTestSampleHex(254, "fd00fe"),
		makeTestSampleHex(math.MaxUint8, "fd00ff"),
		makeTestSampleHex(math.MaxUint8+1, "fd0100"),
		makeTestSampleHex(math.MaxUint16-1, "fdfffe"),
		makeTestSampleHex(math.MaxUint16, "fdffff"),
		makeTestSampleHex(math.MaxUint16+1, "fe00010000"),
		makeTestSampleHex(math.MaxUint32-1, "fefffffffe"),
		makeTestSampleHex(math.MaxUint32, "feffffffff"),
		makeTestSampleHex(math.MaxUint32+1, "ff0000000100000000"),
		makeTestSampleHex(math.MaxUint64-1, "fffffffffffffffffe"),
		makeTestSampleHex(math.MaxUint64, "ffffffffffffffffff"),
	}
	data := make([]byte, MaxSize)
	for _, sample := range samples {
		en := Encode(data, sample.Value)
		var (
			haveBytes = data[:en]
			wantBytes = sample.Bytes
		)
		if !bytes.Equal(haveBytes, wantBytes) {
			t.Fatalf("invalid encoded bytes: have (%x), want (%x)", haveBytes, wantBytes)
		}

		y, dn, err := Decode(haveBytes)
		if err != nil {
			t.Fatal(err)
		}
		var (
			haveValue = y
			wantValue = sample.Value
		)
		if haveValue != wantValue {
			t.Fatalf("invalid decoded value: have %d, want %d", haveValue, wantValue)
		}

		var (
			haveLen = dn
			wantLen = len(sample.Bytes)
		)
		if haveLen != wantLen {
			t.Fatalf("invalid decoded length: have %d, want %d", haveLen, wantLen)
		}
	}
}
