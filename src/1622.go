package main


type Fancy struct {
    operations []bool
    operationsNumber []int
    sequence []int
		last []int
}


func Constructor() Fancy {
    return Fancy{
        operations: make([]bool, 0),
        operationsNumber: make([]int, 0),
        sequence: make([]int, 0),
        last: make([]int, 0),
    }
}


func (this *Fancy) Append(val int)  {
    this.sequence = append(this.sequence, val)
		this.last = append(this.last, len(this.operations))
}


func (this *Fancy) AddAll(inc int)  {
    this.operations = append(this.operations, false)
    this.operationsNumber = append(this.operationsNumber, inc)
}


func (this *Fancy) MultAll(m int)  {
    this.operations = append(this.operations, true)
    this.operationsNumber = append(this.operationsNumber, m)    
}


func (this *Fancy) GetIndex(idx int) int {
    if idx >= len(this.sequence) {
        return -1;
    }
    cur := this.sequence[idx]
		start := this.last[idx]
    for i := start; i < len(this.operations); i++ {
        if this.operations[i] {
            cur = (cur * this.operationsNumber[i]) % 1000000007
        } else {
            cur = (this.operationsNumber[i] + cur) % 1000000007
        }
    }
		this.last[idx] = len(this.operations)
		this.sequence[idx] = cur
    return cur;
}


/**
 * Your Fancy object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Append(val);
 * obj.AddAll(inc);
 * obj.MultAll(m);
 * param_4 := obj.GetIndex(idx);
 */