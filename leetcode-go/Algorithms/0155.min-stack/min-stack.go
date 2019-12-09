package problem0155

type MinStack struct {
	minStack []int
	min      int
}

func Constructor() MinStack {
	return MinStack{minStack: make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	if len(this.minStack) == 0 {
		this.min = x
	}
	this.minStack = append(this.minStack, x)
	if this.min > x {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	minStackLen := len(this.minStack)
	if minStackLen != 0 {
		top := this.Top()
		this.minStack = this.minStack[:minStackLen-1]
		if minStackLen-1 != 0 {
			if this.min <= top {
				this.min = this.Top()
				for _, v := range this.minStack {
					if this.min > v {
						this.min = v
					}
				}
			}
		} else {
			this.min = 0
		}
	}
}

func (this *MinStack) Top() int {
	minStackLen := len(this.minStack)
	if minStackLen != 0 {
		return this.minStack[minStackLen-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
