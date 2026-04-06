package main

func middleNode(head *ListNode) *ListNode {
  mid := (findLengthMiddleNode(head) /2) + 1
	i := 1
	for {
		if i == mid {
			return head
		}
		head = head.Next
		i++
	}
}

func findLengthMiddleNode (head *ListNode) int {
	rN := 0;
	for head != nil {
		rN++
		head = head.Next
	}
	return rN
}