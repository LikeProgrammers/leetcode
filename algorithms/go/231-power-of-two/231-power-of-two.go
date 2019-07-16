package isPowerOfTwo

func isPowerOfTwo(n int) bool {
	bit := 1
	t := 0
	l := 32
	if n <= 0 {
		return false
	}
	for i := 0; i < l; i++ {
		if n&bit != 0 {
			t++
		}
		if t > 1 {
			return false
		}
		bit <<= 1
	}

	return true
}
