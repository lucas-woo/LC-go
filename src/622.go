package main


type MyCircularQueue struct {
    queue []int
    head int
    rear int
    length int
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
        queue: make([]int, k),
    }
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() {
        return false
    }
    if this.IsEmpty() {
        this.head = this.rear
    }
    this.queue[this.rear] = value
    this.rear = (this.rear + 1) % len(this.queue)
    this.length++;
    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {
        return false
    }
    this.head = (this.head + 1) % len(this.queue)
    this.length--;
    return true
}


func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {
        return -1
    }
    return this.queue[this.head]
}


func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {
        return -1
    }    
    if this.rear == 0 {
        return this.queue[len(this.queue) - 1];
    } else {
        return this.queue[this.rear - 1]
    }
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.length == 0
}


func (this *MyCircularQueue) IsFull() bool {
    return this.length == len(this.queue)
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */