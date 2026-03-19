package main
/**
* Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 


import "slices" 
func modifiedList(nums []int, head *ListNode) *ListNode {
	if len(nums) == 0 || head == nil {
		return head
	}
	max := slices.Max(nums);
  hashmap := make([]bool, max + 1);
	for _, v := range nums {
		hashmap[v] = true;
	}
	var returnHead *ListNode;
	var prev *ListNode;
	cur := head
	for {
		if cur == nil {
			prev.Next = nil
			return returnHead
		}
		if cur.Val > max {
			if returnHead == nil {
				returnHead = cur
				prev = returnHead
			} else {
				prev.Next = cur
				prev = cur
			}			
			cur = cur.Next
			continue;
		}
		temp := hashmap[cur.Val];
		if !temp {
			if returnHead == nil {
				returnHead = cur
				prev = returnHead
			} else {
				prev.Next = cur
				prev = cur
			}
		}
		cur = cur.Next
	}
}