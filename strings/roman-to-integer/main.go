package main

import "fmt"

func romanToInt(s string) int {
	values := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	sum := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && values[s[i]] < values[s[i+1]] {
			sum -= values[s[i]]
		} else {
			sum += values[s[i]]
		}
	}
	return sum
}

func main() {
	s := "III"
	fmt.Println(romanToInt(s))

	s = "LVIII"
	fmt.Println(romanToInt(s))

	s = "MCMXCIV"
	fmt.Println(romanToInt(s))

	s = "IV"
	fmt.Println(romanToInt(s))
}
