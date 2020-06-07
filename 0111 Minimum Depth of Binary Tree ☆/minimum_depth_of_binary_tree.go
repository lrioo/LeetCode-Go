package main

import (
	"fmt"

	library "LeetCode-Go/0000-library"
)

func main() {
	Input := []interface{}{3, 9, 20, nil, nil, 15, 7}
	root := library.BuildBinaryTree(Input)
	fmt.Println(root.WithExtent())
	// Output = 2
	fmt.Println(minDepth(root))
}

/**
 * Definition for a binary tree node.
 */
type TreeNode = library.BinaryTreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var min int
	if root.Left != nil {
		min = minDepth(root.Left)
	}
	if root.Right != nil {
		l := minDepth(root.Right)
		if min > 0 && min > l {
			min = l
		}
	}

	return min +1
}

/*
Given a binary tree, find its minimum depth.
给定一个二叉树，找出其最小深度。

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

Note: A leaf is a node with no children.
说明: 叶子节点是指没有子节点的节点。


Example:
示例:

  Given binary tree [3,9,20,null,null,15,7],
  给定二叉树[3,9,20,null,null,15,7],

        3
       / \
      9  20
        /  \
       15   7

  return its minimum depth = 2.
  返回它的最小深度2.
*/
