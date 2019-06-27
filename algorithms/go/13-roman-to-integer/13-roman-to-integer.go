package main

import (
	"fmt"
)

func main() {
	romans := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "LVIII", "MCMXCIV"}
	for _, s := range romans {
		ret := romanToInt(s)
		fmt.Println(ret)
	}
}

func romanToInt(s string) int {
	ret := 0

	symbol_value := map[string]int{
		"M": 1000, "CM": 900, "D": 500, "CD": 400,
		"C": 100, "XC": 90, "L": 50, "XL": 40,
		"X": 10, "IX": 9, "V": 5, "IV": 4,
		"I": 1,
	}

	num := 0
	ls := len(s)
	for i := 0; i < ls; {
		value := 0
		found := false
		if i < ls-1 {
			value, found = symbol_value[s[i:i+2]]
			if found {
				num += value
				i += 2
			}
		}

		if found == false {
			value, found = symbol_value[s[i:i+1]]
			if found {
				num += value
				i += 1
			} else {
				break
			}
		}
	}

	ret = num
	return ret
}
