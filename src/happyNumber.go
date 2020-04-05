package main
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3284/
import "strconv"


func isHappy(n int) bool {
	if n == 4 {
		return false
	}
	result := 0
	str := strconv.Itoa(n)
	for _, char := range str {
		output, _ := strconv.Atoi(string(char))
		result += output * output
	}
	if result == 1 {
		return true
	} else {
		return isHappy(result)
	}
}
