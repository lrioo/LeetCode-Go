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

	res := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		layer := make([]int, 0, len(queue))
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			layer = append(layer, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, nil)
		copy(res[1:], res)
		res[0] = layer
	}

	return res
}

func levelOrderBottomRecursive(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := treeRecursive(root)

	return res
}

func treeRecursive(nodes ...*TreeNode) [][]int {
	ns := make([]*TreeNode, 0, len(nodes)*2)
	layer := make([]int, 0, len(nodes))
	for _, node := range nodes {
		if node.Left != nil {
			ns = append(ns, node.Left)
		}
		if node.Right != nil {
			ns = append(ns, node.Right)
		}
		layer = append(layer, node.Val)
	}

	res := [][]int{layer}
	if len(ns) != 0 {
		res = append(treeRecursive(ns...), res...)
	}

	return res
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
