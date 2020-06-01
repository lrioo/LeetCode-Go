package main

import (
	"fmt"

	library "LeetCode-Go/0000-library"
)

func main() {
	Input := []int{-10, -3, 0, 5, 9}
	// Output = [0,-3,9,-10,null,5,null]
	fmt.Println(sortedArrayToBST(Input).WithExtent())
}

/**
 * Definition for a binary tree node.
 */
type TreeNode = library.BinaryTreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return nil
}

/*
Given an array where elements are sorted in ascending order, convert it to a height balanced BST.
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

For this problem, a height-balanced binary tree is defined as a binary tree in
which the depth of the two subtrees of every node never differ by more than 1.
本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。


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
