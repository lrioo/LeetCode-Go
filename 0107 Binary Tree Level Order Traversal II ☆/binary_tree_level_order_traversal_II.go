package main

import (
	"fmt"

	library "LeetCode-Go/0000-library"
)

func main() {
	Input := []interface{}{3, 9, 20, nil, nil, 15, 7}
	root := library.BuildBinaryTree(Input)
	fmt.Println(root.WithExtent())
	// Output = [[15,7], [9,20], [3]]
	fmt.Println(levelOrderBottom(root))
}

/**
 * Definition for a binary tree node.
 */
type TreeNode = library.BinaryTreeNode

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	return nil
}

/*
Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).
给定一个二叉树，返回其节点值自底向上的层次遍历。（即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）


For example:
例如：

  Given binary tree [3,9,20,null,null,15,7],
  给定二叉树 [3,9,20,null,null,15,7],

        3
       / \
      9  20
        /  \
       15   7

  return its bottom-up level order traversal as:
  返回其自底向上的层次遍历为：

      [
        [15,7],
        [9,20],
        [3]
      ]
*/
