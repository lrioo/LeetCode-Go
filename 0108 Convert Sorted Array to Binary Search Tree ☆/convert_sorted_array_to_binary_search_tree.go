package main

import (
	"fmt"

	. "LeetCode-Go/0000-library"
)

func main() {
	Input := []int{-10, -3, 0, 5, 9}
	// Output = [0,-3,9,-10,null,5,null]
	fmt.Println(sortedArrayToBST(Input).WithExtent())
}

/**
 * Definition for a binary tree node.
 */
type TreeNode = BinaryTreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{}
	left, right, node := 0, len(nums)-1, root

	stack := []int{left, right}
	nodeStack := []*TreeNode{node}
	for len(stack) > 0 {
		left, right, stack = stack[len(stack)-2], stack[len(stack)-1], stack[:len(stack)-2]
		node, nodeStack = nodeStack[len(nodeStack)-1], nodeStack[:len(nodeStack)-1]

		middle := left + (right-left)/2
		node.Val = nums[middle]

		l, r := left, middle-1
		if l <= r {
			node.Left = &TreeNode{}
			nodeStack = append(nodeStack, node.Left)
			stack = append(stack, l, r)
		}

		l, r = middle+1, right
		if l <= r {
			node.Right = &TreeNode{}
			nodeStack = append(nodeStack, node.Right)
			stack = append(stack, l, r)
		}
	}

	return root
}

func sortedArrayToBSTR(nums []int) *TreeNode {
	return sortedArrayToBSTRecursive1(nums)
	//return sortedArrayToBSTRecursive2(nums)
}

func sortedArrayToBSTRecursive1(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	middle := len(nums) / 2
	root := &TreeNode{
		Val: nums[middle],
	}
	root.Left = sortedArrayToBSTRecursive1(nums[:middle])
	root.Right = sortedArrayToBSTRecursive1(nums[middle+1:])

	return root
}

func sortedArrayToBSTRecursive2(nums []int) *TreeNode {
	return createBST(nums, 0, len(nums)-1)
}

func createBST(nums []int, left, right int) *TreeNode {
	if left <= right {
		middle := (left + right) / 2
		root := &TreeNode{Val: nums[middle]}
		root.Left = createBST(nums, left, middle-1)
		root.Right = createBST(nums, middle+1, right)
		return root
	}

	return nil
}

/*
Given an array where elements are sorted in ascending order, convert it to a height balanced BST.
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

For this problem, a height-balanced binary tree is defined as a binary tree in
which the depth of the two subtrees of every node never differ by more than 1.
本题中，一个高度平衡二叉树是指一个二叉树每个节点的左右两个子树的高度差的绝对值不超过1。


Example:
示例:

  Given the sorted array: [-10,-3,0,5,9],
  给定有序数组: [-10,-3,0,5,9],

  One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:
  一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

          0
         / \
       -3   9
       /   /
     -10  5
*/
