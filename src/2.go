package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var cur *ListNode;
	var carry = false
	for l1 != nil || l2 != nil || carry {
		temp := 0;
		if carry {
			temp++;
		}
		if l1 != nil {
			temp += l1.Val;
			l1 = l1.Next
		}
		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}
		if temp >= 10 {
			carry = true
		} else {
			carry = false
		}
		if head == nil {
			head = &ListNode{Val: temp%10}
			cur = head
			continue;
		}
		cur.Next = &ListNode{Val: temp%10}
		cur = cur.Next
	}
	return head
}