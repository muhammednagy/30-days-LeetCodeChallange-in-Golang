package main
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3298/
func findMaxLength(nums []int) int {
	counts := map[int]int{0: -1}
	maxLength, count := 0,0

	for i, v := range nums {
		if v == 0 {
			count--
		} else {
			count++
		}

		if v, ok := counts[count]; ok {
			maxLength = Max(maxLength,  i - v)
		} else {
			counts[count] = i
		}
	}
	return maxLength
}
func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}