package main

// Ref: https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicatesInPlaceArr(nums []int) int {
	lastIdx := 0
	for i := 1; i < len(nums); i++ {
		if nums[lastIdx] != nums[i] {
			lastIdx++
			nums[lastIdx] = nums[i]
		}
	}

	return lastIdx + 1
}
