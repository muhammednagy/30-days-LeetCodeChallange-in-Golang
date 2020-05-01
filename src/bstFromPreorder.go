package main
// https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/530/week-3/3305/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	v := preorder[0]
	node := &TreeNode{v, nil, nil}
	index, left := bstFromPreorderRec(preorder, 1, node)
	node.Left = left
	_, right := bstFromPreorderRec(preorder, index, nil)
	node.Right = right
	return node
}
func bstFromPreorderRec(preorder []int, start int, ancestor *TreeNode) (int, *TreeNode) {
	if start >= len(preorder) {
		return start, nil
	}
	v := preorder[start]
	if ancestor != nil && v > ancestor.Val {
		return start, nil
	}
	node := &TreeNode{v, nil, nil}
	index, left := bstFromPreorderRec(preorder, start + 1, node)
	node.Left = left
	next, right := bstFromPreorderRec(preorder, index, ancestor)
	node.Right = right
	return next, node
}
