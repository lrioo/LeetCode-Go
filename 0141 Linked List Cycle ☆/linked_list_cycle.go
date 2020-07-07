package main

import (
	lib "LeetCode-Go/0000-library"
	"fmt"
)

func main() {
	input, pos := []int{3, 2, 0, -4}, 1
	head := lib.BuildCycleList(input, pos)
	//Output := true
	fmt.Println(hasCycle(head))

	input, pos = []int{1, 2}, 0
	head = lib.BuildCycleList(input, pos)
	//Output := true
	fmt.Println(hasCycle(head))

	input, pos = []int{1}, -1
	head = lib.BuildCycleList(input, pos)
	//Output := false
	fmt.Println(hasCycle(head))
}

/**
 * Definition for singly-linked list.
 */
type ListNode = lib.ListNode

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

/*
Given a linked list, determine if it has a cycle in it.
给定一个链表，判断链表中是否有环。

To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to.
If pos is -1, then there is no cycle in the linked list.
为了表示给定链表中的环，我们使用整数pos来表示链表尾连接到链表中的位置（索引从0开始）。如果pos是-1，则在该链表中没有环。


Example 1:
示例1：

  Input: head = [3,2,0,-4], pos = 1
  输入：head = [3,2,0,-4], pos = 1

  Output: true
  输出：true

  Explanation: There is a cycle in the linked list, where tail connects to the second node.
  解释：链表中有一个环，其尾部连接到第二个节点。


Example 2:
示例2：

  Input: head = [1,2], pos = 0
  输入：head = [1,2], pos = 0

  Output: true
  输出：true

  Explanation: There is a cycle in the linked list, where tail connects to the first node.
  解释：链表中有一个环，其尾部连接到第一个节点。


Example 3:
示例3：

  Input: head = [1], pos = -1
  输入：head = [1], pos = -1

  Output: false
  输出：false

  Explanation: There is no cycle in the linked list.
  解释：链表中没有环。


Follow up:
进阶：

  Can you solve it using O(1) (i.e. constant) memory?
  你能用O(1)（即，常量）内存解决此问题吗？
*/
