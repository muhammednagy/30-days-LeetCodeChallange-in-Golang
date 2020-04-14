package main
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3299/
func stringShift(str string, shift [][]int) string {
	length := len(str)
	for _, v := range shift {
		fmt.Println(v[1])
		if v[0] == 0 {
			str = str[v[1]:] + str[0:v[1]]
		} else {
			str = str[length-v[1]:] + str[0:length-v[1]]
		}
		fmt.Println(str)
	}
	return str
}