package library

import (
	"fmt"
)

func BuildBinaryTree(preOrderArray []interface{}) *TreeNode {
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

		if stack[len(stack)-1].Left != nil {
			stack[len(stack)-1].Right = nil
			stack = stack[:len(stack)-1]
			continue
		}

		stack[len(stack)-1].Left = nil
		if preOrderArray[i+1] == nil {
			stack = stack[:len(stack)-1]
		} else {
			node := &TreeNode{Val: preOrderArray[i+1].(int)}
			stack[len(stack)-1].Right = node
			stack[len(stack)-1] = node
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

func (t *TreeNode) WithExtent() string {
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
