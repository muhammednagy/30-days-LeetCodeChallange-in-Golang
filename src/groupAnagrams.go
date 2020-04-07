package main
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3288/
import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strs {
		bytes := []byte(v)
		sort.SliceStable(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		s := string(bytes)
		m[s] = append(m[s], v)
	}
	var SortedStrings [][]string
	for e := range m {
		SortedStrings = append(SortedStrings, m[e])
	}
	return SortedStrings
}