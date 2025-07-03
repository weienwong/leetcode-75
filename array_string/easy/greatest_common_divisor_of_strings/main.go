package main

import (
	"fmt"
)

func gcdOfStrings(str1 string, str2 string) string {
	if fmt.Sprintf("%s%s", str1, str2) != fmt.Sprintf("%s%s", str2, str1) {
		return ""
	}

	return str1[:gcd(len(str1), len(str2))]

}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var str1, str2 string
	str1 = "ABCABC"
	str2 = "ABC"

	result := gcdOfStrings(str1, str2)
	println(result) // Output: "ABC"

	str1 = "ABABAB"
	str2 = "ABAB"

	result = gcdOfStrings(str1, str2)
	println(result) // Output: "AB"

	str1 = "LEET"
	str2 = "CODE"

	result = gcdOfStrings(str1, str2)
	println(result) // Output: ""

}
