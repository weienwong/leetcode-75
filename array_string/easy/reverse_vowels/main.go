package main

func reverseVowels(s string) string {
	vowelMap := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	left, right := 0, len(s)-1
	runes := []rune(s)

	for left < right {

		if !vowelMap[byte(runes[left])] {
			left++
		}

		if !vowelMap[byte(runes[right])] {
			right--
		}

		// Swap the vowels
		if left < right {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}
	}

	return string(runes)

}

func main() {
	s := "IceCreAm"
	result := reverseVowels(s)
	println(result) // Output: "AceCreIm"

	s = "leetcode"
	result = reverseVowels(s)
	println(result) // Output: "leotcede"
}
