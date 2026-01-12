package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
  pointer := 0
	return mergeTreeTwo(preorder, inorder, &pointer)
}
func mergeTreeTwo(preorder []int, inorder []int, pointer *int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	p := *pointer
	curVal := preorder[p]
	(*pointer)++;
	cur := &TreeNode{
		Val : curVal,
		Left: nil,
		Right: nil,
	}
	i := findIndexTree(inorder, curVal)
	cur.Left = mergeTreeTwo(preorder, inorder[:i], pointer)
	cur.Right = mergeTreeTwo(preorder, inorder[i+1:], pointer)
	return cur
}

func findIndexTree(inorder []int, val int) int {
	for i, v := range inorder {
		if v == val {
			return i
		}
	}
	return 0
}