package main

import (
	"fmt"

	. "LeetCode-Go/0000-library"
)

func main() {
	Input := []interface{}{3, 9, 20, nil, nil, 15, 7}
	root := BuildBinaryTree(Input)
	fmt.Println(root.WithExtent())
	// Output = 3
	fmt.Println(maxDepth(root))
}

/**
 * Definition for a binary tree node.
 */
type TreeNode = BinaryTreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var max int
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		max++
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return max
}

func maxDepthR(root *TreeNode) int {
	return maxDepthRecursive(root)
}

func maxDepthRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var max int
	ld := maxDepthRecursive(root.Left)
	rd := maxDepthRecursive(root.Right)
	if ld > rd {
		max = ld
	} else {
		max = rd
	}

	return 1 + max
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
