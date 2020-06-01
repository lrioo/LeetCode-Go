package main

import (
	"fmt"

	library "LeetCode-Go/0000-library"
)

func main() {
	Input := []interface{}{3, 9, 20, nil, nil, 15, 7}
	root := library.BuildBinaryTree(Input)
	fmt.Println(root.WithExtent())
	// Output = 3
	fmt.Println(maxDepth(root))
}

/**
 * Definition for a binary tree node.
 */
type TreeNode = library.BinaryTreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 0
}

/*
Given a binary tree, find its maximum depth.
给定一个二叉树，找出其最大深度。

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

Note: A leaf is a node with no children.
说明: 叶子节点是指没有子节点的节点。


Example:
示例：

  Given binary tree [3,9,20,null,null,15,7],
  给定二叉树[3,9,20,null,null,15,7]，

        3
       / \
      9  20
        /  \
       15   7

  return its depth = 3.
  返回它的最大深度3。
*/
