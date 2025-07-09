package main

import "fmt"

func compress(chars []byte) int {
	n := len(chars)
	if n == 0 {
		return 0
	}

	writeIndex := 0
	count := 1

	for i := 1; i < n; i++ {
		if chars[i] == chars[i-1] {
			count++
		}
		if chars[i] != chars[i-1] || i == n-1 {
			chars[writeIndex] = chars[i-1]
			writeIndex++
			if count > 1 {
				countStr := []byte{}
				for count > 0 {
					countStr = append(countStr, byte('0'+count%10))
					count /= 10
				}
				// Reverse the count string to get the correct order
				for j := len(countStr) - 1; j >= 0; j-- {
					chars[writeIndex] = countStr[j]
					writeIndex++
				}
			}
			count = 1 // Reset count for the next character
		}
	}

	return len(chars[:writeIndex])
}

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	result := compress(chars)
	fmt.Println("Compressed length:", result)

}
