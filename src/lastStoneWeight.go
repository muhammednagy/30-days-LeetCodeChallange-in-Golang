package main

import "sort"

// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3297/

func lastStoneWeight(stones []int) int {
	for i := len(stones) - 1 ; i >= 1 ; i-- {
		sort.Ints(stones)
		if stones[i] == stones[i-1] {
			if len(stones) == 2 {
				return 0
			}
			stones = stones[0:i-1]
			i--
		} else if stones[i] > stones[i-1] {
			stones = append(stones[:i-1], stones[i]-stones[i-1])
		} else if stones[i] < stones[i-1] {
			stones = append(stones[:i-1], stones[i-1]-stones[i])
		}
	}
	return stones[0]
}
