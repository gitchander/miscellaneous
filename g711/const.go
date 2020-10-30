package g711

const (
	signBit   = 0x80 // Sign bit for a A-law byte.
	quantMask = 0x0f // Quantization field mask.
	segMask   = 0x70 // Segment field mask.
	segShift  = 4    // Left shift for segment number.
)
