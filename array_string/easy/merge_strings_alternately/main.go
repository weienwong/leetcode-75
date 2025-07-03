package main

// 2 pointer approach
func mergeAlternately(word1 string, word2 string) string {
	result := ""
	var i, j int
	for i < len(word1) && j < len(word2) {
		result += string(word1[i]) + string(word2[j])
		i++
		j++
	}

	if i < len(word1) {
		result += word1[i:]
	} else {
		result += word2[j:]
	}

	return result
}

func main() {
	var word2, word1 string
	word1 = "abc"
	word2 = "pqr"
	result := mergeAlternately(word1, word2)
	println(result) // Output: "apbqcr"

	word1 = "ab"
	word2 = "pqrs"
	result = mergeAlternately(word1, word2)
	println(result) // Output: "apbqrs"

	word1 = "abcd"
	word2 = "pq"
	result = mergeAlternately(word1, word2)
	println(result) // Output: "apbqcd"
}
