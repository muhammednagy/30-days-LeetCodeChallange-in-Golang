package main
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3283/
import (
	"sort"
)


func singleNumber(nums []int) int {
	sort.Ints(nums)
	result := 0
	for i := 0; i < len(nums); i++ {
		if len(nums) == 1 {
			return nums[0]
		}
		if i == 0 {
			if nums[i] != nums[i+1] {
				result = nums[i]
				break
			}
		} else if i == len(nums) {
			if nums[i] != nums[i-1] {
				result = nums[i]
				break
			}
		} else {
			if nums[i] != nums[i+1] && nums[i] != nums[i-1] {
				result = nums[i]
				break
			}
		}
	}
	return result
}