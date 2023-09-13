package varint

// Comment from "encoding/binary":
// - signed integers are mapped to unsigned integers using "zig-zag"
//   encoding: Positive values x are written as 2*x + 0, negative values
//   are written as 2*(^x) + 1; that is, negative numbers are complemented
//   and whether to complement is encoded in bit 0.

// ZigZag, Int64Coder
type ZigZagCoder interface {
	Encode(int64) uint64
	Decode(uint64) int64
}

// ------------------------------------------------------------------------------
type zigZagGolang struct{}

var _ ZigZagCoder = zigZagGolang{}

// ZigZagEncode
func (zigZagGolang) Encode(x int64) uint64 {
	ux := uint64(x) << 1
	if x < 0 {
		ux = ^ux
	}
	return ux
}

// ZigZagDecode
func (zigZagGolang) Decode(ux uint64) int64 {
	x := int64(ux >> 1)
	if (ux & 1) != 0 {
		x = ^x
	}
	return x
}

// ------------------------------------------------------------------------------
type zigZagMy struct{}

var _ ZigZagCoder = zigZagMy{}

func (zigZagMy) Encode(x int64) uint64 {
	var ux uint64
	if x < 0 {
		ux = (uint64(-(x + 1)) << 1) + 1
	} else {
		ux = uint64(x) << 1
	}
	return ux
}

func (zigZagMy) Decode(ux uint64) int64 {
	var (
		// x = int64(ux / 2)
		x = int64(ux >> 1)
	)
	if (ux % 2) != 0 {
		x = -x - 1
	}
	return x
}

// ------------------------------------------------------------------------------
var (
	zigZag ZigZagCoder = zigZagGolang{}
	//zigZag ZigZagCoder = zigZagMy{}
)

func EncodeInt64(data []byte, v int64) int {
	u := zigZag.Encode(v)
	return Encode(data, u)
}

func DecodeInt64(data []byte) (int64, int, error) {
	u, n, err := Decode(data)
	v := zigZag.Decode(u)
	return v, n, err
}
