package lib

import "fmt"

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
