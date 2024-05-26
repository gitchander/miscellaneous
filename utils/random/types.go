package random

type Bit = uint

func RandBit(r *Rand) Bit {
	return Bit(r.Int63() & 1)
}

func RandBool(r *Rand) bool {
	return RandBit(r) == 1
}
