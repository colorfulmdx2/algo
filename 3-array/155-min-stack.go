package __array

import "fmt"

type MinStack struct {
	stack    *StackNode
	minStack *StackNode
}

type StackNode struct {
	Val  int
	Next *StackNode
	Prev *StackNode
}

func Constructor() MinStack {
	return MinStack{stack: nil, minStack: nil}
}

func (this *MinStack) Push(val int) {
	node := &StackNode{
		Val:  val,
		Next: nil,
		Prev: this.stack,
	}

	this.stack = node

	if this.minStack == nil || val <= this.minStack.Val {
		newMinNode := &StackNode{
			Val:  val,
			Next: nil,
			Prev: this.minStack,
		}
		this.minStack = newMinNode
	}
}

func (this *MinStack) Pop() {
	if this.stack == nil {
		return
	}

	if this.stack.Val == this.minStack.Val {
		this.minStack = this.minStack.Prev
	}

	this.stack = this.stack.Prev

}

func (this *MinStack) Top() int {
	fmt.Println("stack", this.stack)
	if this.stack != nil {
		return this.stack.Val

	}
	return 0
}

func (this *MinStack) GetMin() int {
	fmt.Println("min", this.minStack)
	if this.minStack != nil {
		return this.minStack.Val

	}
	return 0
}
