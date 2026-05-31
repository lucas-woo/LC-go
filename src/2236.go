package main

//this question can't be real bruh
func checkTree(root *TreeNode) bool {
	if root.Val == (root.Left.Val + root.Right.Val) {
		return true
	} else {
		return false
	}
}