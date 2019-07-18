package convertToTitle

func convertToTitle(n int) string {
	const decimal = 26
	title := []byte{}

	titles := [decimal]byte{}
	for i := 0; i < decimal; i++ {
		titles[i] = byte(int('A') + i)
	}

	for ; n > 0; n /= decimal {
		n -= 1
		remainder := n % decimal
		title = append(title, titles[remainder])
	}

	lt := len(title)
	hlt := lt / 2
	for i := 0; i < hlt; i++ {
		title[i], title[lt-1-i] = title[lt-1-i], title[i]
	}

	return string(title)
}
