package main

func isPalindromeNumber(x int) bool {
	num := x
	palindrome := 0

	for num > 0 {
		mod := num % 10
		num /= 10
		palindrome = palindrome*10 + mod
	}

	return x == palindrome
}
