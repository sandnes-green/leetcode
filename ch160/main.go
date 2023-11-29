// 160. 相交链表
// 简单
// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	mark := map[*ListNode]bool{}

	for headA != nil {
		mark[headA] = true
		headA = headA.Next
	}
	node := &ListNode{}
	flag := false
	for headB != nil {
		if _, ok := mark[headB]; ok {
			node = headB
			flag = true
			break
		}
		headB = headB.Next
	}
	if flag {
		return node
	}
	return nil
}

func main() {
	getIntersectionNode(&ListNode{}, &ListNode{})
}
