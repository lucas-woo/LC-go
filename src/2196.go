package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func createBinaryTree(descriptions [][]int) *TreeNode {
  treeMap := make(map[int]*TreeNode, 0);
	root := make(map[int]bool, 0)
	for i := 0; i < len(descriptions); i++ {
		p := descriptions[i][0];
		c := descriptions[i][1];
		isLeft := descriptions[i][2];
		if _, ok := root[p]; !ok {
			root[p] = true
		}
		root[c] = false

		if t, ok := treeMap[p]; ok {
			if isLeft == 1 {
				if v, ok := treeMap[c]; ok {
					t.Left = v;
				} else {
					treeMap[c] = &TreeNode{
						Val: c,
					}
					t.Left = treeMap[c]
				}
			} else {
				if v, ok := treeMap[c]; ok {
					t.Right = v;
				} else {
					treeMap[c] = &TreeNode{
						Val: c,
					}
					t.Right = treeMap[c]
				}				
			}
		} else {
			t = &TreeNode{
				Val: p,
			}
			treeMap[p] = t
			if isLeft == 1 {
				if v, ok := treeMap[c]; ok {
					t.Left = v;
				} else {
					treeMap[c] = &TreeNode{
						Val: c,
					}
					t.Left = treeMap[c]
				}
			} else {
				if v, ok := treeMap[c]; ok {
					t.Right = v;
				} else {
					treeMap[c] = &TreeNode{
						Val: c,
					}
					t.Right = treeMap[c]
				}				
			}
		}
	}

	for k, v := range root {
		if v {
			return treeMap[k]
		}
	}
	return treeMap[0]
}