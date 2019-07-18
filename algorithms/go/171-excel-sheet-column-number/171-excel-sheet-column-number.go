package titleToNumber

func titleToNumber(s string) int {
	const decimal = 26

	title := 0
	titles := map[byte]int{}
	for i := 0; i < decimal; i++ {
		titles[byte(int('A')+i)] = i + 1
	}

	ls := len(s)
	for i := 0; i < ls; i++ {
		title *= decimal
		title += titles[s[i]]
	}

	return title
}
