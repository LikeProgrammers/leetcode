package reverseVowels

func reverseVowels(s string) string {
	vl := []byte{
		'a', 'e', 'i', 'o', 'u',
		'A', 'E', 'I', 'O', 'U',
	}

	lvl := len(vl)
	vowels := map[byte]bool{}
	for i := 0; i < lvl; i++ {
		vowels[vl[i]] = true
	}

	ss := []byte(s)
	ls := len(ss)
	vs := []int{}
	for i := 0; i < ls; i++ {
		if _, found := vowels[ss[i]]; found {
			vs = append(vs, i)
		}
	}

	lvs := len(vs)
	hlvs := lvs / 2
	for i := 0; i < hlvs; i++ {
		index := vs[i]
		r_index := vs[lvs-1-i]
		ss[index], ss[r_index] = ss[r_index], ss[index]
	}

	return string(ss)
}
