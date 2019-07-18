package reverseString

func reverseString(s []byte) {
	ls := len(s)
	hls := ls / 2
	for i := 0; i < hls; i++ {
		s[i], s[ls-1-i] = s[ls-1-i], s[i]
	}
}
