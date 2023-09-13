package varint

import (
	"encoding/binary"
	"fmt"
	"math"
)

// VarInt
// https://learnmeabitcoin.com/glossary/varint

// CompactSize
// https://bitcoin.org/en/developer-reference#raw-transaction-format

const (
	bytesPerUint8  = 1
	bytesPerUint16 = 2
	bytesPerUint32 = 4
	bytesPerUint64 = 8
)

const (
	compactMaxUint8 = 0xfc

	compactTagUint16 = 0xfd
	compactTagUint32 = 0xfe
	compactTagUint64 = 0xff
)

const (
	MaxSize = 1 + bytesPerUint64
)

var (
	byteOrder = binary.BigEndian
	//byteOrder = binary.LittleEndian
)

func Encode(data []byte, v uint64) int {
	switch {
	case v <= compactMaxUint8:
		{
			data[0] = uint8(v)
			return bytesPerUint8
		}
	case v <= math.MaxUint16:
		{
			data[0] = compactTagUint16
			byteOrder.PutUint16(data[1:], uint16(v))
			return 1 + bytesPerUint16
		}
	case v <= math.MaxUint32:
		{
			data[0] = compactTagUint32
			byteOrder.PutUint32(data[1:], uint32(v))
			return 1 + bytesPerUint32
		}
	default:
		{
			data[0] = compactTagUint64
			byteOrder.PutUint64(data[1:], v)
			return 1 + bytesPerUint64
		}
	}
}

func Decode(data []byte) (uint64, int, error) {
	tag := data[0]
	switch tag {
	case compactTagUint16:
		{
			v := byteOrder.Uint16(data[1:])
			var err error
			if v <= compactMaxUint8 {
				err = errCompactSize(uint64(v))
			}
			return uint64(v), 1 + bytesPerUint16, err
		}
	case compactTagUint32:
		{
			v := byteOrder.Uint32(data[1:])
			var err error
			if v <= math.MaxUint16 {
				err = errCompactSize(uint64(v))
			}
			return uint64(v), 1 + bytesPerUint32, err
		}
	case compactTagUint64:
		{
			v := byteOrder.Uint64(data[1:])
			var err error
			if v <= math.MaxUint32 {
				err = errCompactSize(v)
			}
			return v, 1 + bytesPerUint64, err
		}
	default:
		return uint64(tag), bytesPerUint8, nil
	}
}

func errCompactSize(v uint64) error {
	return fmt.Errorf("DecodeVarint: value %d isn't compact", v)
}
