package library

import (
	"fmt"
)

//Layer
func BuildBinaryTree(layerArray []interface{}) *BinaryTreeNode {
	if len(layerArray) == 0 || layerArray[0] == nil {
		return nil
	}

	stack := make([]*BinaryTreeNode, 0)
	head := &BinaryTreeNode{Val: layerArray[0].(int)}

	stack = append(stack, head)
	for i := 1; i < len(layerArray) && len(stack) > 0; i += 2 {
		d := stack[0]
		stack = stack[1:]
		if layerArray[i] != nil {
			l := &BinaryTreeNode{Val: layerArray[i].(int)}
			d.Left = l
			stack = append(stack, l)
		}
		if i+1 < len(layerArray) && layerArray[i+1] != nil {
			r := &BinaryTreeNode{Val: layerArray[i+1].(int)}
			d.Right = r
			stack = append(stack, r)
		}
	}
	return head
}

//DLR
func BuildBinaryTreeByDLR(DLRArray []interface{}) *BinaryTreeNode {
	if len(DLRArray) == 0 || DLRArray[0] == nil {
		return nil
	}

	if len(DLRArray)%2 == 0 {
		DLRArray = append(DLRArray, nil)
	}

	stack := make([]*BinaryTreeNode, 0)
	head := &BinaryTreeNode{Val: DLRArray[0].(int)}

	stack = append(stack, head)
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

	return head
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
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
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
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node = node.Right
		}
	}
	s = fmt.Sprintf("%s, null", s)

	return s
}
