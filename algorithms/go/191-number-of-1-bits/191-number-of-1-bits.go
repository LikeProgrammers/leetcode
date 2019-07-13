package hammingWeight

func hammingWeight(num uint32) int {
	bit := 0
	var mask uint32 = 1
	for i := 0; i < 32; i++ {
		if (num & mask) != 0 {
			bit++
		}
		mask <<= 1
	}

	return bit
}
