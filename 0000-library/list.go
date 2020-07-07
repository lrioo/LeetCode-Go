package lib

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return "<nil>"
	}

	return fmt.Sprintf("%d", l.Val)
}

func (l *ListNode) WithExtend() string {
	if l == nil {
		return "<nil>"
	}

	s := fmt.Sprintf("%d", l.Val)
	cur := l.Next
	for cur != nil {
		s = fmt.Sprintf("%s->%d", s, cur.Val)
		cur = cur.Next
	}
	return s
}

func BuildList(input []int) *ListNode {
	if len(input) == 0 {
		return nil
	}

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

func BuildCycleList(input []int, pos int) *ListNode {
	if len(input) == 0 {
		return nil
	}

	head := &ListNode{
		Val:  input[0],
		Next: nil,
	}

	var cNode *ListNode
	if pos == 0 {
		cNode = head
	}

	pre := head
	for i := 1; i < len(input); i++ {
		next := &ListNode{
			Val:  input[i],
			Next: nil,
		}
		pre.Next = next
		pre = next

		if i == pos {
			cNode = next
		}
		if i == len(input)-1 {
			pre.Next = cNode
		}
	}

	return head
}

func BuildIntersectList(intersectVal int, listA []int, listB []int, skipA, skipB int) (*ListNode, *ListNode) {
	if intersectVal == 0 {
		return BuildList(listA), BuildList(listB)
	}

	var intersection *ListNode
	headA := &ListNode{
		Val:  listA[0],
		Next: nil,
	}

	pre := headA
	for i := 1; i < len(listA); i++ {
		next := &ListNode{
			Val:  listA[i],
			Next: nil,
		}
		pre.Next = next
		pre = next

		if i == skipA {
			intersection = next
		}
	}

	headB := &ListNode{
		Val:  listB[0],
		Next: nil,
	}

	pre = headB
	for i := 1; i < skipB; i++ {
		next := &ListNode{
			Val:  listB[i],
			Next: nil,
		}
		pre.Next = next
		pre = next
	}
	pre.Next = intersection

	return headA, headB
}
