package main

import (
	"fmt"

	. "LeetCode-Go/0000-library"
)

func main() {
	Input := []interface{}{3, 9, 20, nil, nil, 15, 7}
	root := BuildBinaryTree(Input)
	fmt.Println(root.WithExtent())
	// Output = true
	fmt.Println(isBalanced(root))

	Input = []interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4}
	root = BuildBinaryTree(Input)
	fmt.Println(root.WithExtent())
	// Output = false
	fmt.Println(isBalanced(root))

	Input = []interface{}{1,2,2,3,3,3,3,4,4,4,4,4,4,nil,nil,5,5}
	root = BuildBinaryTree(Input)
	fmt.Println(root.WithExtent())
	// Output = false
	fmt.Println(isBalanced(root))
}

/**
 * Definition for a binary tree node.
 */
type TreeNode = BinaryTreeNode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	nodes := []*TreeNode{root}
	depthMap := make(map[*TreeNode]int)
	for i := 0; i < len(nodes); {
		offset := len(nodes)
		for j := i; j < offset; j++ {
			node := nodes[j]
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
			depthMap[node] = 0
		}
		i = offset
	}

	for i := len(nodes) - 1; i >= 0; i-- {
		node := nodes[i]
		l, r := 0, 0
		if ld, ok := depthMap[node.Left]; ok {
			l = ld
		}
		if rd, ok := depthMap[node.Right]; ok {
			if l < rd {
				l, r = rd, l
			} else {
				r = rd
			}
		}

		if l-r > 1 {
			return false
		}
		depthMap[node] = l + 1
	}

	return true
}

func isBalancedR(root *TreeNode) bool {
	_, f := isBalancedRecursive(root)
	return f
}

func isBalancedRecursive(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	ld, lf := isBalancedRecursive(node.Left)
	if !lf {
		return 0, false
	}
	rd, rf := isBalancedRecursive(node.Right)
	if !rf {
		return 0, false
	}

	if ld < rd {
		ld, rd = rd, ld
	}

	return ld + 1, ld-rd <= 1
}

/*
Given a binary tree, determine if it is height-balanced.
给定一个二叉树，判断它是否是高度平衡的二叉树。

For this problem, a height-balanced binary tree is defined as:
本题中，一棵高度平衡二叉树定义为：

  a binary tree in which the left and right subtrees of every node differ in height by no more than 1.
  一个二叉树每个节点的左右两个子树的高度差的绝对值不超过1。


Example 1:
示例 1:

  Given the following tree [3,9,20,null,null,15,7]:
  给定二叉树 [3,9,20,null,null,15,7]

        3
       / \
      9  20
        /  \
       15   7

  Return true.
  返回true。


Example 2:
示例 2:

  Given the following tree [1,2,2,3,3,null,null,4,4]:
  给定二叉树 [1,2,2,3,3,null,null,4,4]

           1
          / \
         2   2
        / \
       3   3
      / \
     4   4

  Return false.
  返回false。
*/
