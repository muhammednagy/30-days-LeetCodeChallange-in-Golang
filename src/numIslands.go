package main
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/530/week-3/3302/
func numIslands(grid [][]byte) int {
  count := 0
  for i := 0 ; i < len(grid) ; i++ {
	for j := 0; j < len(grid[i]) ; j++ {
	  if grid[i][j] == '1' {
		count++
		DFS(grid, i, j)
	  }
	}
  }
  return count
}

func DFS(grid [][]byte, i int, j int){
  if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
	return
  }
  grid[i][j] = '0'
  DFS(grid, i+1, j)
  DFS(grid, i-1, j)
  DFS(grid, i, j-1)
  DFS(grid, i, j+1)
}