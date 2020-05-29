package main

import (
	library "LeetCode-Go/0000-library"
	"fmt"
)

func main() {
	Input := []interface{}{1, 2, 2, 3, 4, 4, 3}
	root := library.BuildBinaryTree(Input)
	fmt.Println(root.WithExtent())
	fmt.Println(isSymmetric(root))

	Input = []interface{}{1, 2, 2, nil, 3, nil, 3}
	root = library.BuildBinaryTree(Input)
	fmt.Println(root.WithExtent())
	fmt.Println(isSymmetric(root))
}

/**
 * Definition for a binary tree node.
 */
type TreeNode = library.BinaryTreeNode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := []*TreeNode{root.Left, root.Right}
	for len(stack) > 0 && len(stack)%2 == 0 {
		p, q := stack[0], stack[1]
		stack = stack[2:]

		if p == nil || q == nil {
			if p == q {
				continue
			}
			return false
		}

		if p.Val != q.Val {
			return false
		}
		stack = append(stack, p.Left, q.Right, p.Right, q.Left)
	}

	return true
}

func isSymmetricRecursive(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		if p == q {
			return true
		}
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSymmetricTree(p.Left, q.Right) && isSymmetricTree(p.Right, q.Left)
}

/*
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
给定一个二叉树，检查它是否是镜像对称的。


For example:
例如：

  this binary tree [1,2,2,3,4,4,3] is symmetric:
  二叉树[1,2,2,3,4,4,3]是对称的：

        1
       / \
      2   2
     / \ / \
    3  4 4  3

But the following [1,2,2,null,3,null,3] is not:
但是[1,2,2,null,3,null,3]则不是镜像对称的:

        1
       / \
      2   2
       \   \
       3    3


Follow up:
进阶：

  Solve it both recursively and iteratively.
  你可以运用递归和迭代两种方法解决这个问题吗？
*/
