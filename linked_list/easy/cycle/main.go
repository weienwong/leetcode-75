package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

func main() {
	head := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4}}}}
	// Creating a cycle for testing
	head.Next.Next.Next = head.Next // Cycle here
	result := hasCycle(head)
	println(result) // Output: true

	head = &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	head.Next.Next = head // Creating a cycle
	result = hasCycle(head)
	println(result) // Output: false

}
