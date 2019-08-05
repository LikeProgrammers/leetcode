package defangipaddr

func defangIPaddr(address string) string {
	ret := []byte{}
	for _, v := range address {
		if v-'.' != 0 {
			ret = append(ret, byte(v))
		} else {
			ret = append(ret, '[')
			ret = append(ret, byte(v))
			ret = append(ret, ']')
		}
	}
	return string(ret)
}
