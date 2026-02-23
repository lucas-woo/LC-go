package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func recoverTree(root *TreeNode) {
	arr := make([]*TreeNode, 0)
	inOrderTrav(root, &arr)
	prev := arr[0]
	var not *TreeNode
	var not2 *TreeNode
	for i := 1; i < len(arr); i++ {
		if arr[i].Val < prev.Val {
			if not != nil {
				not.Val, arr[i].Val = arr[i].Val, not.Val
				return
			}
			not = prev;
			not2 = arr[i]
		}
		prev = arr[i]
	}
	not.Val, not2.Val = not2.Val, not.Val
}
func inOrderTrav(root *TreeNode, arr *[]*TreeNode) {
	if root == nil {
		return
	}
	inOrderTrav(root.Left, arr)
	*arr = append(*arr, root)
	inOrderTrav(root.Right, arr)
}