package main

func maxVowels(s string, k int) int {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	maxCount := 0
	currentCount := 0

	for i := 0; i < len(s); i++ {
		if vowels[s[i]] {
			currentCount++
		}

		if i >= k {
			if vowels[s[i-k]] {
				currentCount--
			}
		}

		if currentCount > maxCount {
			maxCount = currentCount
		}
	}
	return maxCount
}

func main() {
	s := "abciiidef"
	k := 3
	result := maxVowels(s, k)
	println(result) // Output: 3

	s = "aeiou"
	k = 2
	result = maxVowels(s, k)
	println(result) // Output: 2

	s = "leetcode"
	k = 3
	result = maxVowels(s, k)
	println(result) // Output: 2

}
