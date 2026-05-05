func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var lengthOfList int = 1;
	var endOfList *ListNode;
	cur := head
	for {
		if cur.Next == nil {
			endOfList = cur
			break;
		}
		lengthOfList++;
		cur = cur.Next
	}
	endIndex := lengthOfList - (k % lengthOfList);
	cur = head
	i := 1;
	for {
		if cur == nil {
			break;
		}
		if i == endIndex {
			if cur.Next == nil {
				break
			}
			endOfList.Next = head
			head = cur.Next
			cur.Next = nil
			break;
		}
		cur = cur.Next
		i++;
	}
	return head
}