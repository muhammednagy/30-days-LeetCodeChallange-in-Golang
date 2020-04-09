package main
// https://leetcode.com/explore/other/card/30-day-leetcoding-challenge/528/week-1/3289/
func countElements(arr []int) int {
	result := 0
	for _, element := range arr {
		for _, secondElement := range arr {
			if element + 1 == secondElement {
				result += 1
				break
			}
		}
	}
	return result
}