package main

import "fmt"

func main() {
	Input := []int{1, 1, 2}
	head := buildList(Input)
	fmt.Println("Input:", head)
	//Output: 1->2
	fmt.Println("Output:", deleteDuplicates(head))

	Input = []int{1, 1, 2, 3, 3}
	head = buildList(Input)
	fmt.Println("Input:", head)
	//Output: 1->2->3
	fmt.Println("Output:", deleteDuplicates(head))
}

func buildList(input []int) *ListNode {
	head := &ListNode{
		Val:  input[0],
		Next: nil,
	}

	pre := head
	for i := 1; i < len(input); i++ {
		next := &ListNode{
			Val:  input[i],
			Next: nil,
		}
		pre.Next = next
		pre = next
	}

	return head
}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return ""
	}

	s := fmt.Sprintf("%d", l.Val)
	cur := l.Next
	for cur != nil {
		s = fmt.Sprintf("%s->%d", s, cur.Val)
		cur = cur.Next
	}
	return s
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre, cur := head, head.Next
	for cur != nil {
		if pre.Val == cur.Val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = pre.Next
	}

	return head
}

/*
Given a sorted linked list, delete all duplicates such that each element appear only once.
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

Example 1:
示例1:

  Input: 1->1->2
  输入: 1->1->2

  Output: 1->2
  输出: 1->2


Example 2:
示例2:

  Input: 1->1->2->3->3
  输入: 1->1->2->3->3

  Output: 1->2->3
  输出: 1->2->3
*/
