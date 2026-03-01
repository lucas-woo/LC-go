package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "strconv"
func sumRootToLeaf(root *TreeNode) int {
	rN := 0
	findSumRootToLeaf(root, "", &rN)
	return rN
}

func findSumRootToLeaf(root *TreeNode, s string, final *int) {

	cur := ""
	if root.Val == 0 {
		cur = "0"
	} else {
		cur = "1"
	}
	s = s + cur
	if root.Left == nil && root.Right == nil {
		i, _ := strconv.ParseInt(s, 2, 64)
		*final += int(i)
		return
	}
	if root.Left != nil {
		findSumRootToLeaf(root.Left, s , final)
	}
	if root.Right != nil {
		findSumRootToLeaf(root.Right, s , final)
	}
} 