package isIsomorphic

import "strings"

func isIsomorphic(s string, t string) bool {
	for i := 0; i < len(s); i++ {
		if strings.IndexByte(s, s[i]) != strings.IndexByte(t, t[i]) {
			return false
		}
	}
	return true
}

// "ab", "aa" "abab", "baba"
func isIsomorphic2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	l := len(s)

	rsm := map[int]bool{}
	tm := map[byte]bool{}
	rs := []byte(s)
	for i := 0; i < l; i++ {
		if tm[t[i]] {
			continue
		}
		tm[t[i]] = true
		for j := 0; j < l; j++ {
			if rs[j] == s[i] {
				if rsm[j] {
					continue
				}
				rsm[j] = true
				rs[j] = t[i]
			}
		}
	}

	rrs := string(rs)
	if rrs == t {
		return true
	} else {
		return false
	}
}
