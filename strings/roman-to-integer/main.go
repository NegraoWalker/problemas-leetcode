package main

import "fmt"

func romanToInt(s string) int {

	romanToInt := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romanToInt[s[i]] < romanToInt[s[i+1]] {
			result -= romanToInt[s[i]]
		} else {
			result += romanToInt[s[i]]
		}
	}

	return result

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

	s = "CD"
	fmt.Println(romanToInt(s))
}
