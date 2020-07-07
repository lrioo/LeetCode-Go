package main

import (
	lib "LeetCode-Go/0000-library"
	"fmt"
)

func main() {
	intersectVal, listA, listB, skipA, skipB := 8, []int{4, 1, 8, 4, 5}, []int{5, 0, 1, 8, 4, 5}, 2, 3
	headA, headB := lib.BuildIntersectList(intersectVal, listA, listB, skipA, skipB)
	//Output := 8
	fmt.Println(getIntersectionNode(headA, headB))

	intersectVal, listA, listB, skipA, skipB = 2, []int{0, 9, 1, 2, 4}, []int{3, 2, 4}, 3, 1
	headA, headB = lib.BuildIntersectList(intersectVal, listA, listB, skipA, skipB)
	//Output = 2
	fmt.Println(getIntersectionNode(headA, headB))

	intersectVal, listA, listB, skipA, skipB = 0, []int{2, 6, 4}, []int{1, 5}, 3, 2
	headA, headB = lib.BuildIntersectList(intersectVal, listA, listB, skipA, skipB)
	//Output = nil
	fmt.Println(getIntersectionNode(headA, headB))
}

/**
 * Definition for singly-linked list.
 */
type ListNode = lib.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a, b := headA, headB
	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}

		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}
	return a
}

/*
Write a program to find the node at which the intersection of two singly linked lists begins.
编写一个程序，找到两个单链表相交的起始节点。


Example 1:
示例 1：

  Input:
  输入：
    intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3

  Output:
  输出：
    Reference of the node with value = 8

  Input Explanation:
    The intersected node's value is 8 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [4,1,8,4,5].
    From the head of B, it reads as [5,6,1,8,4,5]. There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.
  输入解释：
    相交节点的值为8（注意，如果两个链表相交则不能为0）。从各自的表头开始算起，链表A为[4,1,8,4,5]，链表B为[5,0,1,8,4,5]。
    在A中，相交节点前有2个节点；在B中，相交节点前有3个节点。


Example 2:
示例 2：

  Input:
  输入：
    intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1

  Output:
  输出：
    Reference of the node with value = 2

  Input Explanation:
    The intersected node's value is 2 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [1,9,1,2,4].
    From the head of B, it reads as [3,2,4]. There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in B.
  输入解释：
    相交节点的值为2（注意，如果两个链表相交则不能为0）。从各自的表头开始算起，链表A为[0,9,1,2,4]，链表B为[3,2,4]。
    在A中，相交节点前有3个节点；在B中，相交节点前有1个节点。


Example 3:
示例 3：

  Input:
  输入：
    intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2

  Output: null
  输出：null

  Input Explanation:
    From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5]. Since the two lists do not intersect, intersectVal must be 0,
    while skipA and skipB can be arbitrary values.
  输入解释：
    从各自的表头开始算起，链表A为[2,6,4]，链表B为[1,5]。由于这两个链表不相交，所以intersectVal必须为0，而skipA和skipB可以是任意值。

  Explanation: The two lists do not intersect, so return null.
  解释：这两个链表不相交，因此返回null。


Notes:
注意：

  If the two linked lists have no intersection at all, return null.
  如果两个链表没有交点，返回 null.

  The linked lists must retain their original structure after the function returns.
  在返回结果后，两个链表仍须保持原有的结构。

  You may assume there are no cycles anywhere in the entire linked structure.
  可假定整个链表结构中没有循环。

  Each value on each linked list is in the range [1, 10^9].
  每个链表中节点的值范围为[1, 10^9]。

  Your code should preferably run in O(n) time and use only O(1) memory.
  程序尽量满足O(n)时间复杂度，且仅用O(1)内存。
*/
