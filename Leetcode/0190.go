package leetcode

// my solution
func reverseBits(num uint32) uint32 {
	b := []uint32{}
	for i := 0; i < 32; i++ {
		b = append(b, num%2)
		num /= 2
	}
	t := uint32(1)
	for i := 31; i >= 0; i-- {
		num += b[i] * t
		t *= 2
	}

	return num
}

// shift digit
func reverseBits(num uint32) uint32 {
	var res uint32 = 0
	for i := 0; i < 32 && num > 0; i++ {
		res += (num & 1) << (31 - i)
		num >>= 1
	}
	return res
}

// Divide and Conquer
func reverseBits(num uint32) uint32 {
	const M1 uint32 = 0x55555555 // 0101
	const M2 uint32 = 0x33333333 // 0011
	const M4 uint32 = 0x0F0F0F0F // 00001111
	const M8 uint32 = 0x00FF00FF // 0000000011111111

	num = num>>1&M1 | (num&M1)<<1
	num = num>>2&M2 | (num&M2)<<2
	num = num>>4&M4 | (num&M4)<<4
	num = num>>8&M8 | (num&M8)<<8
	return num>>16 | num<<16
}
