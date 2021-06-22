package main

import "fmt"

type MinStack struct {
	stack []int
	sorts []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	st, sr := make([]int, 0), make([]int, 0)
	return MinStack{stack: st, sorts: sr}
}

func (this *MinStack) Push(x int) {
	fmt.Printf("Push: %v\n", this)
	this.stack = append(this.stack, x)
    
    if len(this.sorts) == 0 {
		this.sorts = append(this.sorts, x)
		return
	}

	if x > this.sorts[0] {
		tmp := make([]int, len(this.sorts))
		copy(tmp, this.sorts)
		this.sorts = append([]int{x}, tmp...)
        return
	}
    
	for i := 0; i < len(this.sorts); i++ {
		if x > this.sorts[i] {
            tmp := make([]int, 1+i)
            copy(tmp, this.sorts[0:i])
            fmt.Println(tmp)
            fmt.Println(this.sorts)
			this.sorts = this.sorts[0:i]
            this.sorts = append(this.sorts, x)
            fmt.Println(this.sorts)
			this.sorts = append(this.sorts, tmp...)
			fmt.Println(this.sorts)
            return
		}
	}
    this.sorts = append(this.sorts, x)
}

func (this *MinStack) Pop() {
	fmt.Printf("Pop: %v\n", this)
    val := this.stack[len(this.stack)-1]
	this.stack = this.stack[0 : len(this.stack)-1]
    fmt.Println(val)
	for i := 0; i < len(this.sorts); i++ {
		if val == this.sorts[i] {
			this.sorts = append(this.sorts[0:i], this.sorts[i+1:]...)
            fmt.Println(i, this.sorts)
			return
		}
	}
}

func (this *MinStack) Top() int {
	fmt.Printf("Top: %v\n", this)
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.sorts[len(this.sorts)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
 
/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {

}
