package canConstruct

func canConstruct(ransomNote string, magazine string) bool {
	mr := map[byte]int{}
	mm := map[byte]int{}

	lm := len(magazine)
	for i := 0; i < lm; i++ {
		mm[magazine[i]]++
	}
	lr := len(ransomNote)
	for i := 0; i < lr; i++ {
		v := ransomNote[i]
		mr[v]++
		c, found := mm[v]
		if found == false || mr[v] > c {
			return false
		}
	}

	return true
}
