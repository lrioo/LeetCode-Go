package main

import "fmt"

func main() {
	Input1 := []interface{}{1, 2, nil, nil, 3, nil, nil}
	root1 := buildTree(Input1)
	fmt.Println(root1.withExtent())

	Input2 := []interface{}{1, 2, nil, nil, 3, nil, nil}
	root2 := buildTree(Input2)
	fmt.Println(root2)

	fmt.Println(isSameTree(root1, root2))
}

func buildTree(preOrderArray []interface{}) *TreeNode {
	if len(preOrderArray) == 0 || preOrderArray[0] == nil {
		return nil
	}

	if len(preOrderArray)%2 == 0 {
		preOrderArray = append(preOrderArray, nil)
	}

	stack := make([]*TreeNode, 0)
	head := &TreeNode{Val: preOrderArray[0].(int)}

	stack = append(stack, head)
	for i := 1; i < len(preOrderArray); i++ {
		if preOrderArray[i] != nil {
			node := &TreeNode{Val: preOrderArray[i].(int)}
			if stack[len(stack)-1].Left == nil {
				stack[len(stack)-1].Left = node
			} else {
				stack[len(stack)-1].Right = node
			}
			stack = append(stack, node)
			continue
		}

		stack[len(stack)-1].Left = nil
		if preOrderArray[i+1] == nil {
			stack = stack[:len(stack)-1]
		} else {
			node := &TreeNode{Val: preOrderArray[i+1].(int)}
			stack[len(stack)-1].Right = node
			stack = append(stack[:len(stack)-1], node)
		}
		i++
	}

	return head
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	if t == nil {
		return ""
	}

	stack := []*TreeNode{t}
	s := fmt.Sprintf("%d", t.Val)

	node := t.Left
	for node != nil || len(stack) > 0 {
		for node != nil {
			s = fmt.Sprintf("%s, %d", s, node.Val)
			stack = append(stack, node)
			node = node.Left
		}

		if len(stack) > 0 {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node = node.Right
		}
	}

	return s
}

func (t *TreeNode) withExtent() string {
	if t == nil {
		return ""
	}

	stack := []*TreeNode{t}
	s := fmt.Sprintf("%d", t.Val)

	node := t.Left
	for node != nil || len(stack) > 0 {
		for node != nil {
			s = fmt.Sprintf("%s, %d", s, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		s = fmt.Sprintf("%s, null", s)

		if len(stack) > 0 {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node = node.Right
		}
	}
	s = fmt.Sprintf("%s, null", s)

	return s
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return false
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
        [1,2,3],   [1,2,3]

  Output: true
  输出: true


Example 2:
示例 2:

  Input:
  输入:
           1         1
          /           \
         2             2
        [1,2],     [1,null,2]

  Output: false
  输出: false


Example 3:
示例 3:

  Input:
  输入:
           1         1
          / \       / \
         2   1     1   2
        [1,2,1],   [1,1,2]

  Output: false
  输出: false
*/
