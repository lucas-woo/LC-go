package main


type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
  pointer := len(postorder) - 1;
  return mergeTree(inorder, postorder, &pointer)
}

func mergeTree(inorder []int, postorder []int, pointer *int) *TreeNode {
  if len(inorder) == 0 {
    return nil
  }
  cur := &TreeNode{
    Val: postorder[(*pointer)],
    Left: nil,
    Right: nil,
  }
  curIndex := findIndex(postorder[(*pointer)], inorder)
  (*pointer)--;
  if curIndex != len(inorder) {
    cur.Right = mergeTree(inorder[curIndex+1:], postorder, pointer)
  }
  cur.Left = mergeTree(inorder[:curIndex], postorder, pointer)
  return cur
}

func findIndex(n int, inorder []int) int {
  for i, v := range inorder {
    if v == n {
      return i
    }
  }
  return 0;
}