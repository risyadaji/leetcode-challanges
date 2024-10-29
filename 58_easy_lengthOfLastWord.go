package main

import "strings"

func lengthOfLastWord(s string) int {
	split := strings.Split(strings.TrimSpace(s), " ")

	lastWord := split[len(split)-1]

	return len(lastWord)
}
