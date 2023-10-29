package leetcode

type MinStack struct {
	stk    []int
	minStk []int
}

func Constructor() MinStack {
	return MinStack{stk: make([]int, 0), minStk: make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	this.stk = append(this.stk, val)
	var n = len(this.minStk)
	if n > 0 {
		var t = this.minStk[n-1]
		if t < val {
			val = t
		}
	}
	this.minStk = append(this.minStk, val)
}

func (this *MinStack) Pop() {
	var n = len(this.stk)
	this.stk = this.stk[:n-1]
	this.minStk = this.minStk[:n-1]
}

func (this *MinStack) Top() int {
	return this.stk[len(this.stk)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStk[len(this.minStk)-1]
}
