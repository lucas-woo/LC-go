package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
  arr := make([]*ListNode, 0)
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	sorted := mergeSortList(arr)
	for i := 0; i < len(sorted) - 1; i++ {
		sorted[i].Next = sorted[i + 1]
	}
	sorted[len(arr) - 1].Next = nil
	return sorted[0]
}
func mergeSortList(arr []*ListNode) []*ListNode {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]
	return mergeList(mergeSortList(left), mergeSortList(right))
}
func mergeList(left, right []*ListNode) []*ListNode {
	sorted := make([]*ListNode, 0)
	i := 0;
	j := 0;
	for i < len(left) && j < len(right) {
		if left[i].Val < right[j].Val {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[j])
			j++			
		}
	}
	for i < len(left) {
		sorted = append(sorted, left[i])
		i++		
	}
	for j < len(right) {
		sorted = append(sorted, right[j])
		j++					
	}
	return sorted
}