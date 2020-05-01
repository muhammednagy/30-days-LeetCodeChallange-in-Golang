package main
// https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/530/week-3/3304/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1
	for left < right {
		midPoint := left + (right - left) / 2
		if nums[midPoint] > nums[right] {
			left = midPoint+1
		} else {
			right = midPoint
		}
	}

	start := left
	left = 0
	right = len(nums) - 1

	if target >= nums[start] && target <= nums[right] {
		left = start
	} else {
		right = start
	}

	for left <= right {
		midPoint := left + (right - left) / 2
		if nums[midPoint] == target {
			return midPoint
		} else if nums[midPoint] < target {
			left = midPoint+1
		} else {
			right = midPoint-1
		}
	}

	return -1
}
