package main

import "math"

// implemented Kadane’s Algorithm
func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
func maxSubArray(nums []int) int {
	n := len(nums)
	localMax := 0
	globalMax :=  math.MinInt32

	for i := 0; i < n ; i++ {
		localMax = Max(nums[i], nums[i] +localMax)
		if localMax > globalMax {
			globalMax = localMax
		}
	}

	return globalMax

}
