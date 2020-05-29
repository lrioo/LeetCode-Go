package main

import (
	library "LeetCode-Go/0000-library"
	"fmt"
)

func main() {
	Input1 := []interface{}{1, 2, nil, nil, 3, nil, nil}
	root1 := library.BuildBinaryTree(Input1)
	fmt.Println(root1.WithExtent())

	Input2 := []interface{}{1, 2, nil, nil, 3, nil, nil}
	root2 := library.BuildBinaryTree(Input2)
	fmt.Println(root2.WithExtent())

	fmt.Println(isSameTree(root1, root2))

	Input1 = []interface{}{1, 2, nil, nil, nil}
	root1 = library.BuildBinaryTree(Input1)
	fmt.Println(root1.WithExtent())

	Input2 = []interface{}{1, nil, 2, nil, nil}
	root2 = library.BuildBinaryTree(Input2)
	fmt.Println(root2.WithExtent())

	fmt.Println(isSameTree(root1, root2))

	Input1 = []interface{}{1, 1, nil, nil, 2, nil, nil}
	root1 = library.BuildBinaryTree(Input1)
	fmt.Println(root1.WithExtent())

	Input2 = []interface{}{1, 2, nil, nil, 1, nil, nil}
	root2 = library.BuildBinaryTree(Input2)
	fmt.Println(root2.WithExtent())

	fmt.Println(isSameTree(root1, root2))
}

/**
 * Definition for a binary tree node.
 */
type TreeNode = library.BinaryTreeNode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	stack := []*TreeNode{p, q}
	for len(stack) > 0 && len(stack)%2 == 0 {
		p, q = stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]

		if p == nil || q == nil {
			if p == q {
				continue
			}
			return false
		}

		if p.Val != q.Val {
			return false
		}
		stack = append(stack, p.Right, q.Right, p.Left, q.Left)
	}

	return true
}

func isSameTreeRecursive(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		if p == q {
			return true
		}
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreeRecursive(p.Left, q.Left) && isSameTreeRecursive(p.Right, q.Right)
}

/*
Given two binary trees, write a function to check if they are the same or not.
给定两个二叉树，编写一个函数来检验它们是否相同。

Two binary trees are considered the same if they are structurally identical and the nodes have the same value.
如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

Example 1:
示例1:

  Input:
  输入:
           1         1
          / \       / \
         2   3     2   3
        [1,2,3]   [1,2,3]

  Output: true
  输出：true


Example 2:
示例2:

  Input:
  输入:
           1         1
          /           \
         2             2
       [1,2]      [1,null,2]

  Output: false
  输出：false


Example 3:
示例3:

  Input:
  输入:
           1         1
          / \       / \
         2   1     1   2
        [1,2,1]   [1,1,2]

  Output: false
  输出：false
*/
