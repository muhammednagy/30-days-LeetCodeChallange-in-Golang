package main
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3290/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func linkedListLength(head *ListNode) int {
	count := 1
	for head.Next != nil {
		count++
		head = head.Next
	}
	return count
}

func middleNode(head *ListNode) *ListNode {
	length := linkedListLength(head)
	element := length / 2
	node := head
	for i := 0; i < element; i++ {
		node = node.Next
	}
	return node
}