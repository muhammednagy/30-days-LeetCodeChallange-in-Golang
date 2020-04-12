package main

// https://leetcode.com/explore/other/card/30-day-leetcoding-challenge/529/week-2/3293/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return (1 + Max(height(node.Left), height(node.Right)));
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lheight := height(root.Left)
	rheight := height(root.Right)
	ldiameter := diameterOfBinaryTree(root.Left)
	rdiameter := diameterOfBinaryTree(root.Right)
	return Max(lheight + rheight, Max(ldiameter, rdiameter))
}