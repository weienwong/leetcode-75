package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	if head == nil || head.Next == nil {
		return 0
	}

	// we know the length of the linked list is even
	slow, fast := head, head
	var prev *ListNode
	prev = nil

	// Find the middle of the linked list using the slow and fast pointer technique
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}

	// Reverse the second half of the linked list
	var prevHalf *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prevHalf
		prevHalf = slow
		slow = next
	}

	// Calculate the maximum twin sum
	maxTwinSum := 0
	firstHalf := head
	secondHalf := prevHalf

	for firstHalf != nil && secondHalf != nil {
		twinSum := firstHalf.Val + secondHalf.Val
		if twinSum > maxTwinSum {
			maxTwinSum = twinSum
		}
		firstHalf = firstHalf.Next
		secondHalf = secondHalf.Next
	}

	return maxTwinSum

}

func main() {
	head := &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}}
	result := pairSum(head)
	println(result) // Output: 6

	head = &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}}
	result = pairSum(head)
	println(result) // Output: 7

	head = &ListNode{Val: 1, Next: &ListNode{Val: 100000, Next: &ListNode{Val: 1, Next: nil}}}
	result = pairSum(head)
	println(result) // Output: 100001
}
