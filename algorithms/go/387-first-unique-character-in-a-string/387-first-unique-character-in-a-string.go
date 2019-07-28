package firstUniqChar

func firstUniqChar1(s string) int {
	ms := map[byte][]int{}
	first := -1

	for i := 0; i < len(s); i++ {
		v := s[i]
		ms[v] = append(ms[v], i)
	}

	for _, v := range ms {
		c := len(v)
		if (first == -1 || first > v[0]) && c == 1 {
			first = v[0]
		}
	}

	return first
}

func firstUniqChar2(s string) int {
	ms := map[byte]int{}

	for i := 0; i < len(s); i++ {
		ms[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if ms[s[i]] == 1 {
			return i
		}
	}

	return -1
}

func firstUniqChar(s string) int {
	ms := [26]int{}

	for _, v := range s {
		ms[v-'a']++
	}
	for i, v := range s {
		if ms[v-'a'] == 1 {
			return i
		}
	}

	return -1
}
