package main


type MyQueue struct {
	queue []int
	headPointer int
}

func Constructor() MyQueue {
  return MyQueue{queue: make([]int, 0), headPointer: 0}
}


func (this *MyQueue) Push(x int) {
	(*this).queue = append((*this).queue, x);
}


func (this *MyQueue) Pop() int {
  if this.Empty() {
		return -1;
	}
	this.headPointer++;
	return this.queue[this.headPointer - 1]
}


func (this *MyQueue) Peek() int {
  if this.Empty() {
		return -1;
	}    
	return this.queue[this.headPointer]
}


func (this *MyQueue) Empty() bool {
  return (len((*this).queue) - (*this).headPointer) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */