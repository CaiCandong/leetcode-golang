package problem0155_1

type MinStack struct {
	Stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	if len(this.Stack) == 0 || this.minStack[len(this.minStack)-1] >= x {
		this.minStack = append(this.minStack, x)
	}
	this.Stack = append(this.Stack, x)
}

func (this *MinStack) Pop() {
	if len(this.Stack) == 0 {
		return
	}
	if this.minStack[len(this.minStack)-1] == this.Stack[len(this.Stack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	this.Stack = this.Stack[:len(this.Stack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
