package main

import "fmt"

func reverseWords(s string) string {
	// Split the string by spaces and filter out empty strings

	words := []string{}
	start := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			if start < i {
				words = append(words, s[start:i])
			}
			start = i + 1
		}
	}

	// Reverse the words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Join the words with a single space
	result := ""
	for i, word := range words {
		if i > 0 {
			result = fmt.Sprintf("%s ", result)
		}
		result = fmt.Sprintf("%s%s", result, word)
	}

	return result

}

func main() {
	s := "the sky is blue"
	result := reverseWords(s)
	println(result) // Output: "blue is sky the"

	s = "  hello world  "
	result = reverseWords(s)
	println(result) // Output: "world hello"

	s = "a good   example"
	result = reverseWords(s)
	println(result) // Output: "example good a"
}
