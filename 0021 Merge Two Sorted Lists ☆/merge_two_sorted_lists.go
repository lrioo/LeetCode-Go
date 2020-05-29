package main

import (
	library "LeetCode-Go/0000-library"
	"fmt"
)

func main() {
	Input1 := []int{1, 2, 4}
	head1 := library.BuildList(Input1)
	fmt.Println("Input:", head1)

	Input2 := []int{1, 3, 4}
	head2 := library.BuildList(Input2)
	fmt.Println("Input:", head2)

	//Output: 1->1->2->3->4->4
	fmt.Println(mergeTwoLists(head1, head2))
}

//ListNode Definition for singly-linked list.
type ListNode = library.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	pre := head
	for {
		if l1 == nil {
			if l2 != nil {
				pre.Next = l2
			}
			break
		} else {
			if l2 == nil {
				pre.Next = l1
				break
			}
		}

		if l1.Val <= l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}

	return head.Next
}

/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。


Example:
示例：

  Input: 1->2->4, 1->3->4
  输入：1->2->4, 1->3->4

  Output: 1->1->2->3->4->4
  输出：1->1->2->3->4->4
*/
