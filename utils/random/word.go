package random

const (
	bitsPerByte = 8

	bitsPerUint32 = 32
	bitsPerUint64 = 64
)

//------------------------------------------------------------------------------

// type word = uint32

// const bitsPerWord = bitsPerUint32

// func randWord(r *Rand) word {
// 	return r.Uint32()
// }

//------------------------------------------------------------------------------

type word = uint64

const bitsPerWord = bitsPerUint64

func randWord(r *Rand) word {
	return r.Uint64()
}

//------------------------------------------------------------------------------
