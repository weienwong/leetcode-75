package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	oddEvenList(head)

	for head != nil {
		println(head.Val)
		head = head.Next
	}
	// Output: 1 -> 3 -> 5 -> 2 -> 4

	head = &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 7}}}}}}}
	oddEvenList(head)

	for head != nil {
		println(head.Val)
		head = head.Next
	}
	// Output: 2 -> 3 -> 6 -> 7 -> 1 -> 5 -> 4s

}
