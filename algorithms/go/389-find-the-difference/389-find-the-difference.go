package findTheDifference

func findTheDifference(s string, t string) byte {
	ms := map[byte]int{}
	mt := map[byte]int{}

	for i := 0; i < len(s); i++ {
		ms[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		mt[t[i]]++
	}
	for i := 0; i < len(t); i++ {
		count, found := ms[t[i]]
		if count != mt[t[i]] || found == false {
			return t[i]
		}
	}

	return '0'
}
