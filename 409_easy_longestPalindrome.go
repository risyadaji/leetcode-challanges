package main

// Ref: https://leetcode.com/problems/longest-palindrome

func longestPalindrome(s string) int {
	n := len(s)
	pairs := 0

	m := make(map[rune]int)
	for _, char := range s {
		m[char]++
		if m[char]%2 == 0 {
			pairs++
		}
	}

	if pairs*2 == n {
		return n
	}

	return pairs*2 + 1
}
