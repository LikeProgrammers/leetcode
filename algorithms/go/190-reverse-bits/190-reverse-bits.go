package reverseBits

func reverseBits(num uint32) uint32 {
	l := 32
	var r_num uint32 = 0
	var bit uint32 = 1
	var value uint32 = 0
	for i := 0; i < l; i++ {
		value = 0
		if num&bit != 0 {
			value = 1
		}
		r_num <<= 1
		r_num += value
		bit <<= 1
	}
	return r_num
}
