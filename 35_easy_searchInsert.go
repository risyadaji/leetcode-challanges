package main

// Ref: https://leetcode.com/problems/search-insert-position/description/
// using binary search tree

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2

		if target == nums[mid] {
			return mid
		} else if target <= nums[left] {
		}

		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right + 1
}
