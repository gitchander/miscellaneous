package varint

import (
	"math"
	"testing"
)

type int64Sample struct {
	i64 int64
	u64 uint64
}

func TestZigZagGolang(t *testing.T) {
	testZigZagCoder(t, zigZagGolang{})
}

func TestZigZagMy(t *testing.T) {
	testZigZagCoder(t, zigZagMy{})
}

func testZigZagCoder(t *testing.T, coder ZigZagCoder) {
	samples := []int64Sample{

		{0, 0},
		{1, 2},
		{2, 4},
		{3, 6},
		{math.MaxInt8, math.MaxInt8 * 2},
		{math.MaxInt16, math.MaxInt16 * 2},
		{math.MaxInt32, math.MaxInt32 * 2},
		{math.MaxInt64, math.MaxInt64 * 2},

		// Some negative:
		{-1, 1},
		{-2, 3},
		{-3, 5},
		{math.MinInt8, math.MaxUint8},
		{math.MinInt16, math.MaxUint16},
		{math.MinInt32, math.MaxUint32},
		{math.MinInt64 + 7, math.MaxUint64 - 2*7},
		{math.MinInt64, math.MaxUint64},
	}
	for _, sample := range samples {
		var (
			haveU64 = coder.Encode(sample.i64)
			wantU64 = sample.u64
		)
		if haveU64 != wantU64 {
			t.Fatalf("Invalid zig-zag Encode: have %d, want %d", haveU64, wantU64)
		}

		var (
			haveI64 = coder.Decode(haveU64)
			wantI64 = sample.i64
		)
		if haveI64 != wantI64 {
			t.Fatalf("Invalid zig-zag Decode convert: have %d, want %d", haveI64, wantI64)
		}
	}
}
