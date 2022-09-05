package utils

// https://en.wikipedia.org/wiki/Nibble

func byteToNibbles(b byte) (hi, lo byte) {
	hi = b >> 4
	lo = b & 0xf
	return
}

func nibblesToByte(hi, lo byte) (b byte) {
	b |= hi << 4
	b |= lo & 0xf
	return
}
