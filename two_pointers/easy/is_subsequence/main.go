package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	sIndex, tIndex := 0, 0
	for tIndex < len(t) {
		if s[sIndex] == t[tIndex] {
			sIndex++
			if sIndex == len(s) {
				return true
			}
		}
		tIndex++
	}
	return false
}

func main() {
	s := "abc"
	t := "ahbgdc"

	println(isSubsequence(s, t)) // Output: true

	println(isSubsequence("axc", t)) // Output: false
}
