package lib

import (
	"fmt"
)

//Layer
func BuildBinaryTree(layerArray []interface{}) *BinaryTreeNode {
	if len(layerArray) == 0 || layerArray[0] == nil {
		return nil
	}

	root := &BinaryTreeNode{Val: layerArray[0].(int)}
	queue := []*BinaryTreeNode{root}
	for i := 1; i < len(layerArray) && len(queue) > 0; i += 2 {
		node := queue[0]
		queue = queue[1:]
		if layerArray[i] != nil {
			l := &BinaryTreeNode{Val: layerArray[i].(int)}
			node.Left = l
			queue = append(queue, l)
		}
		if i+1 < len(layerArray) && layerArray[i+1] != nil {
			r := &BinaryTreeNode{Val: layerArray[i+1].(int)}
			node.Right = r
			queue = append(queue, r)
		}
	}
	return root
}

//DLR
func BuildBinaryTreeByDLR(DLRArray []interface{}) *BinaryTreeNode {
	if len(DLRArray) == 0 || DLRArray[0] == nil {
		return nil
	}

	if len(DLRArray)%2 == 0 {
		DLRArray = append(DLRArray, nil)
	}

	root := &BinaryTreeNode{Val: DLRArray[0].(int)}
	stack := []*BinaryTreeNode{root}
	for i := 1; i < len(DLRArray); i++ {
		if DLRArray[i] != nil {
			node := &BinaryTreeNode{Val: DLRArray[i].(int)}
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
		if DLRArray[i+1] == nil {
			stack = stack[:len(stack)-1]
		} else {
			node := &BinaryTreeNode{Val: DLRArray[i+1].(int)}
			stack[len(stack)-1].Right = node
			stack[len(stack)-1] = node
		}
		i++
	}

	return root
}

//LDR
func BuildBinaryTreeByLDR(LDRArray []interface{}) *BinaryTreeNode {
	return nil
}

//LRD
func BuildBinaryTreeByLRD(LRDArray []interface{}) *BinaryTreeNode {
	return nil
}

/**
 * Definition for a binary tree node.
 */
type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func (t *BinaryTreeNode) String() string {
	if t == nil {
		return ""
	}

	stack := []*BinaryTreeNode{t}
	s := fmt.Sprintf("%d", t.Val)

	node := t.Left
	for node != nil || len(stack) > 0 {
		for node != nil {
			s = fmt.Sprintf("%s, %d", s, node.Val)
			stack = append(stack, node)
			node = node.Left
		}

		if len(stack) > 0 {
			node, stack = stack[len(stack)-1], stack[:len(stack)-1]
			node = node.Right
		}
	}

	return s
}

func (t *BinaryTreeNode) WithExtent() string {
	if t == nil {
		return ""
	}

	stack := []*BinaryTreeNode{t}
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
			node, stack = stack[len(stack)-1], stack[:len(stack)-1]
			node = node.Right
		}
	}
	s = fmt.Sprintf("%s, null", s)

	return s
}
