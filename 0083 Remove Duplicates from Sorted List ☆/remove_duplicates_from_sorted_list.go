package main

import (
	library "LeetCode-Go/0000library"
	"fmt"
)

func main() {
	Input := []int{1, 1, 2}
	head := library.BuildList(Input)
	fmt.Println("Input:", head)
	//Output: 1->2
	fmt.Println("Output:", deleteDuplicates(head))

	Input = []int{1, 1, 2, 3, 3}
	head = library.BuildList(Input)
	fmt.Println("Input:", head)
	//Output: 1->2->3
	fmt.Println("Output:", deleteDuplicates(head))
}

/**
 * Definition for singly-linked list.
 */
type ListNode = library.ListNode

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
