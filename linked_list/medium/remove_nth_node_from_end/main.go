package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}

	dummy := &ListNode{Next: head}
	first, second := dummy, dummy

	// Move first pointer n+1 steps ahead
	for i := 0; i <= n; i++ {
		if first == nil {
			return head // n is larger than the length of the list
		}
		first = first.Next
	}

	// Move both pointers until first reaches the end
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// second is now at the node before the one we want to remove
	if second.Next != nil {
		second.Next = second.Next.Next // Remove the nth node from the end
	}

	return dummy.Next // Return the new head of the list
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	head = removeNthFromEnd(head, 2)
	for head != nil {
		println(head.Val)
		head = head.Next
	}

	head = &ListNode{Val: 1}
	head = removeNthFromEnd(head, 1)
	if head == nil {
		println("List is empty")
	} else {
		println(head.Val)
	}
	head = &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	head = removeNthFromEnd(head, 1)
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}
