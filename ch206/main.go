// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}

	return prev
}

func reverseList1(head *ListNode) *ListNode {
	return recur(head, nil)
}

func recur(cur *ListNode, prev *ListNode) *ListNode {
	fmt.Println(cur)
	fmt.Println(prev)
	if cur == nil {
		return prev
	}

	cur.Next = prev

	res := recur(cur.Next, cur)

	return res
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Println(reverseList(head))
	fmt.Println(reverseList1(head))
}
