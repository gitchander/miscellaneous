package g711

var seg_aend = [8]int16{
	0x001F, // 00000000 00011111
	0x003F, // 00000000 00111111
	0x007F, // 00000000 01111111
	0x00FF, // 00000000 11111111
	0x01FF, // 00000001 11111111
	0x03FF, // 00000011 11111111
	0x07FF, // 00000111 11111111
	0x0FFF, // 00001111 11111111
}

// linear2alaw() - Convert a 16-bit linear PCM value to 8-bit A-law
//
// linear2alaw() accepts an 16-bit integer and encodes it as A-law data.
//
// 	Linear Input Code	Compressed Code
// ------------------------	---------------
// 0000000wxyza			000wxyz
// 0000001wxyza			001wxyz
// 000001wxyzab			010wxyz
// 00001wxyzabc			011wxyz
// 0001wxyzabcd			100wxyz
// 001wxyzabcde			101wxyz
// 01wxyzabcdef			110wxyz
// 1wxyzabcdefg			111wxyz
//
// For further information see John C. Bellamy's Digital Telephony, 1982,
// John Wiley & Sons, pps 98-111 and 472-476.

// 2's complement (16-bit range)
func LinearToAlaw(pcm_val int16) uint8 {

	pcm_val = pcm_val >> 3

	var mask uint8
	if pcm_val >= 0 {
		mask = 0xD5 // sign (7th) bit = 1
	} else {
		mask = 0x55 // sign bit = 0
		pcm_val = -pcm_val - 1
	}

	// Convert the scaled magnitude to segment number.
	seg := search(pcm_val, seg_aend[:])

	// Combine the sign, segment, and quantization bits.

	if seg >= 8 { // out of range, return maximum value.
		return uint8(0x7F ^ mask)
	}

	aval := uint8(seg << segShift)
	if seg < 2 {
		aval |= uint8((pcm_val >> 1) & quantMask)
	} else {
		aval |= uint8((pcm_val >> seg) & quantMask)
	}
	return (aval ^ mask)
}

func search(val int16, table []int16) int {
	for i := range table {
		if val <= table[i] {
			return i
		}
	}
	return len(table)
}

// Convert an A-law value to 16-bit linear PCM
func AlawToLinear(a_val uint8) int16 {

	var t int16
	var seg int16

	a_val ^= 0x55

	t = int16(a_val&quantMask) << 4
	seg = int16(a_val&segMask) >> segShift
	switch seg {
	case 0:
		t += 8
		break
	case 1:
		t += 0x108
		break
	default:
		t += 0x108
		t <<= seg - 1
	}

	if (a_val & signBit) != 0 {
		return t
	} else {
		return -t
	}
}
