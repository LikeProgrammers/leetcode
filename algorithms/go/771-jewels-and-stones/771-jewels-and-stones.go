package stones

import (
	"strings"
)

func numJewelsInStones(J string, S string) int {
	ret := 0
	for _, v := range J {
		ret += strings.Count(S, string(v))
	}
	return ret
}
