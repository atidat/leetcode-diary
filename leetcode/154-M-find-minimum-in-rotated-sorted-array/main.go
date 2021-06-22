package main

type MinStack struct {
	stack []int
	sort  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	st := make([]int, 0)
	sr := make([]int, 0)
	return MinStack{stack: st, sort: sr}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	for k := range this.stack {
		if x < this.sort[k] {
			temp := this.sort[k+1:]
			this.sort = append(this.sort[0:k], x)
			this.sort = append(this.sort, temp...)
			break
		}
	}
}

func (this *MinStack) Pop() {
	this.stack = 
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {

}

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
