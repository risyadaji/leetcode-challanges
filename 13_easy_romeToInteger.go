package main

// Reference: https://leetcode.com/problems/roman-to-integer/description/

func romanToInt(s string) int {
	romans := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	sum := 0
	for i := 0; i < len(s); i++ {
		sum += romans[s[i]]
		if i != 0 {
			if romans[s[i-1]] < romans[s[i]] {
				sum -= 2 * romans[s[i-1]]
			}
		}
	}

	return sum
}
