// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3287/
package main

func maxProfit(prices []int) int {
	maxProfit := 0
	if len(prices) <= 1 {
		return maxProfit
	}
	for i:= 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit +=  prices[i] - prices[i-1]
		}
	}
	return maxProfit

}