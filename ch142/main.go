package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) *ListNode {
	slow, fast := head, head // 乌龟和兔子同时从起点出发
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // 乌龟走一步
		fast = fast.Next.Next // 兔子走两步
		if fast == slow {     // 兔子追上乌龟（套圈），说明有环
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

func main() {

}
