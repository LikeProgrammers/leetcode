package isPerfectSquare

func isPerfectSquare(num int) bool {
	for i := 1; i <= num; i++ {
		p := i * i
		if p == num {
			return true
		} else if p > num {
			return false
		}
	}

	return false
}
