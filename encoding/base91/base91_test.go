package base91

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"math"
	"testing"

	"github.com/gitchander/miscellaneous/encoding/base91/utils/random"
)

func TestRandomBytes(t *testing.T) {
	r := random.NewRandNow()
	for i := 0; i < 1000; i++ {
		bs := make([]byte, r.Intn(200))
		random.FillBytes(r, bs)

		en := EncodedLenMax(len(bs))
		ebs := make([]byte, en)
		n := Encode(ebs, bs)
		ebs = ebs[:n]

		//t.Logf("%s", ebs)

		dn := DecodedLenMax(len(ebs))
		dbs := make([]byte, dn)
		n, err := Decode(dbs, ebs)
		if err != nil {
			t.Fatal(err)
		}
		dbs = dbs[:n]

		if !bytes.Equal(bs, dbs) {
			t.Fatalf("[%x] != [%x]", bs, dbs)
		}
	}
}

func TestRandomString(t *testing.T) {
	r := random.NewRandNow()
	for i := 0; i < 1000; i++ {
		bs := make([]byte, r.Intn(200))
		random.FillBytes(r, bs)
		s := EncodeToString(bs)
		//t.Log(s)
		dbs, err := DecodeString(s)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(bs, dbs) {
			t.Fatalf("[%x] != [%x]", bs, dbs)
		}
	}
}

func TestEncodeSamples(t *testing.T) {
	samples := []struct {
		data   []byte
		result []byte
	}{
		{
			data:   hexDecode(""),
			result: hexDecode(""),
		},
		{
			data:   hexDecode("0000"),
			result: []byte("AAA"),
		},
		{
			data:   []byte{88},
			result: []byte("}A"),
		},
		{
			data:   []byte{89},
			result: []byte("~A"),
		},
		{
			data:   []byte{88, 88},
			result: []byte("s)C"),
		},
	}
	for _, sample := range samples {
		en := EncodedLenMax(len(sample.data))
		//t.Log(en)
		ebs := make([]byte, en)
		n := Encode(ebs, sample.data)
		ebs = ebs[:n]
		//t.Log(len(ebs), string(ebs))
		if !bytes.Equal(ebs, sample.result) {
			t.Fatalf("%s != %s", ebs, sample.result)
		}
	}
}

func TestBytes(t *testing.T) {
	bs := make([]byte, 2)
	n := math.MaxUint16 + 1
	for i := 0; i < n; i++ {
		binary.BigEndian.PutUint16(bs, uint16(i))
		en := EncodedLenMax(len(bs))
		ebs := make([]byte, en)
		n := Encode(ebs, bs)
		ebs = ebs[:n]
		// if n == 3 {
		// 	t.Logf("[% x]: %s", bs, ebs)
		// }
	}
}

func hexDecode(s string) []byte {
	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}
