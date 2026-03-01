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
	return mergeTreeTwo(preorder, inorder, &pointer, 0, len(preorder) - 1)
}

func mergeTreeTwo(preorder []int, inorder []int, pointer *int, left, right int) *TreeNode {
	if left > right {
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
	cur.Left = mergeTreeTwo(preorder, inorder, pointer, left, i - 1)
	cur.Right = mergeTreeTwo(preorder, inorder, pointer, i + 1, right)
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