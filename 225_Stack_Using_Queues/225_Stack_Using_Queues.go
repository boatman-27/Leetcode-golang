package main

type MyStack struct {
	arr []int
}

func Constructor() MyStack {
	return MyStack{arr: []int{}}
}

func (this *MyStack) Push(x int) {
	this.arr = append(this.arr, x)
}

func (this *MyStack) Pop() int {
	if len(this.arr) == 0 {
		return -1
	}

	val := this.arr[len(this.arr)-1]
	this.arr = this.arr[:len(this.arr)-1]
	return val
}

func (this *MyStack) Top() int {
	return this.arr[len(this.arr)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.arr) == 0
}
