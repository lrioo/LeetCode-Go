package main

import library "LeetCode-Go/0000-library"

func main() {

}

/**
 * Definition for a binary tree node.
 */
type TreeNode = library.BinaryTreeNode

func hasPathSum(root *TreeNode, sum int) bool {
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
