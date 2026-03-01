package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func balanceBST(root *TreeNode) *TreeNode {
  sortedArr := make([]int, 0)
	createArr(root, &sortedArr)	
	return buildBST(sortedArr)
}

func buildBST(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	mid := len(arr) / 2
	cur := &TreeNode{
		Val: arr[mid],
	}
	cur.Left = buildBST(arr[0:mid])
	cur.Right = buildBST(arr[mid+1:])
	return cur
}

func createArr(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	createArr(root.Left, arr)
	*arr = append(*arr, root.Val)
	createArr(root.Right, arr)
}