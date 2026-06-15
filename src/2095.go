package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
  cur := head
	length := 1;
	for cur.Next != nil {
		cur = cur.Next
		length++;
	}
	mid := length / 2
	start := 0;
	cur = head
	for start < mid - 1 {
		cur = cur.Next
		start++
	}
	next := cur.Next
	cur.Next = next.Next
	return head
}