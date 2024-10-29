package main

func strStr(haystack string, needle string) int {
	nHaystack, nNeedle := len(haystack), len(needle)
	if nNeedle > nHaystack {
		return -1
	}

	if needle == haystack {
		return 0
	}

	for i := 0; i <= nHaystack-nNeedle; i++ {
		substr := haystack[i : i+nNeedle]
		if substr == needle {
			return i
		}
	}

	return -1
}
