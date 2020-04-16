package main
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/530/week-3/3301/
import "fmt"

func checkValidString(s string) bool {
	open := 0
	closed := 0
	length := len(s) - 1
	fmt.Println(length)
	// 40 is ASCII for (
	// 41 is ASCII for )
	// 42 is ASCII for *
	for i, v := range s {
		if v == 40 || v == 42 {
			open++
		} else {
			open--
		}
		if s[length - i] == 42 || s[length - i] == 41 {
			closed++
		} else {
			closed--
		}
		if open < 0 || closed < 0 {
			return false
		}
	}
	return true
}