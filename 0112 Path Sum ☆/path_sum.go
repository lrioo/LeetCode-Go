package main

import (
	library "LeetCode-Go/0000-library"
	"fmt"
)

func main() {
	Input, sum := []interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}, 22
	root := library.BuildBinaryTree(Input)
	fmt.Println(root.WithExtent())
	// Output = true
	fmt.Println(hasPathSum(root, sum))
}

/**
 * Definition for a binary tree node.
 */
type TreeNode = library.BinaryTreeNode

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left == nil && node.Right == nil {
			if node.Val == sum {
				return true
			}
			continue
		}

		if node.Left != nil {
			node.Left.Val += node.Val
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			node.Right.Val += node.Val
			stack = append(stack, node.Right)
		}
	}

	return false
}

func hasPathSumR(root *TreeNode, sum int) bool {
	return hasPathSumRecursive(root, sum)
}

func hasPathSumRecursive(node *TreeNode, sum int) bool {
	if node == nil {
		return false
	}

	if node.Left == nil && node.Right == nil && node.Val == sum {
		return true
	}

	if hasPathSumRecursive(node.Left, sum-node.Val) {
		return true
	}
	if hasPathSumRecursive(node.Right, sum-node.Val) {
		return true
	}

	return false
}

/*
Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

Note: A leaf is a node with no children.
说明: 叶子节点是指没有子节点的节点。


Example:
示例:

  Given the below binary tree and sum = 22,
  给定如下二叉树，以及目标和sum=22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1

  return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
  返回true，因为存在目标和为22的根节点到叶子节点的路径5->4->11->2。
*/
