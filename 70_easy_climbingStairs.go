package main

// Ref: https://leetcode.com/problems/climbing-stairs/description/

var mem [45]int

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	var n1, n2 int
	if mem[n-1] == 0 {
		mem[n-1] = climbStairs(n - 1)
	}
	if mem[n-2] == 0 {
		mem[n-2] = climbStairs(n - 2)
	}

	n1 = mem[n-1]
	n2 = mem[n-2]

	return n1 + n2
}
