package main
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3292/
type Node struct {
	data int
	prev *Node
	minValue *Node
}
type MinStack struct {
	top *Node
}

func Constructor() MinStack {
	return MinStack{}
}

/** initialize your data structure here. */


func (this *MinStack) Push(x int)  {
 node := Node{data: x, prev: this.top}
 if this.top != nil{
	 if x < this.top.minValue.data  {
		node.minValue = &node
	 } else {
	 	node.minValue = this.top.minValue
	 }
 } else {
	 node.minValue = &node
 }
 this.top = &node
}


func (this *MinStack) Pop()  {
	this.top = this.top.prev
}


func (this *MinStack) Top() int {
	return this.top.data
}


func (this *MinStack) GetMin() int {
	return this.top.minValue.data
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
