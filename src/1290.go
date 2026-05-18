package main

func getDecimalValue(head *ListNode) int {
	total := 0
	for head != nil {
		total *= 2
		if head.Val == 1 {
			total++
		}
		head = head.Next
	}
	return total
}