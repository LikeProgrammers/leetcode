package isAnagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ms := map[int]int{}
	mt := map[int]int{}
	for _, v := range s {
		ms[int(v)-'a']++
	}
	for _, v := range t {
		mt[int(v)-'a']++
	}

	for s_value, s_count := range ms {
		if t_count, t_found := mt[s_value]; !t_found || t_count != s_count {
			return false
		}
	}

	return true
}
