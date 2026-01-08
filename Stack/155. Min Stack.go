type MinStack struct {
    stack    []int
    minStack []int
}

func Constructor() MinStack {
    return MinStack{
        stack:    []int{},
        minStack: []int{},
    }
}

func (this *MinStack) Push(val int) {
    this.stack = append(this.stack, val)
    // Only push to minStack if it's empty or the new val is <= current min
    if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
        this.minStack = append(this.minStack, val)
    }
}

func (this *MinStack) Pop() {
    if len(this.stack) == 0 {
        return
    }
    
    topIdx := len(this.stack) - 1
    val := this.stack[topIdx]
    this.stack = this.stack[:topIdx]

    // If we popped the value that was the current minimum, pop it from minStack too
    if val == this.minStack[len(this.minStack)-1] {
        this.minStack = this.minStack[:len(this.minStack)-1]
    }
}

// Fixed: Only returns 1 value to match the interface
func (this *MinStack) Top() int {
    if len(this.stack) == 0 {
        return 0 
    }
    return this.stack[len(this.stack)-1]
}

// Fixed: Only returns 1 value to match the interface
func (this *MinStack) GetMin() int {
    if len(this.minStack) == 0 {
        return 0
    }
    return this.minStack[len(this.minStack)-1]
}