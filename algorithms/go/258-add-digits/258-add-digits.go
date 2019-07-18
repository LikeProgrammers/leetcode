package addDigits

func addDigits1(num int) int {
	if num/10 <= 0 {
		return num
	}

	sum := 0
	for ; num > 0; num /= 10 {
		sum += num % 10
	}

	return addDigits(sum)
}

func addDigits(num int) int {
	if num > 9 {
		num %= 9
		if num == 0 {
			return 9
		} else {
			return num
		}
	} else {
		return num
	}
}
