package main
// https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/530/week-3/3303/
func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i, _ := range dp {
		dp[i] = make([]int, len(grid[0]))
		for j,_ := range dp[i] {
			dp[i][j] += grid[i][j]
			if i > 0 && j > 0 {
				dp[i][j] += min(dp[i][j - 1], dp[i - 1][j])
			} else if i > 0 {
				dp[i][j] += dp[i - 1][j]
			} else if j > 0 {
				dp[i][j] += dp[i][j - 1]
			}
		}
	}
	return dp[len(grid) - 1][len(grid[0]) - 1]
}

func min(a int, b int) int{
	if a < b {
		return a
	}
	return b
}