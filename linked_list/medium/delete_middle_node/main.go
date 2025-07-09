package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	// need to get the length of the linked list
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}

	if length == 1 {
		return nil // If there's only one node, return nil
	}

	middleIndex := length / 2
	dummy := &ListNode{Next: head}
	// dummy := head
	prev := dummy

	for i := 0; i < middleIndex; i++ {
		prev = prev.Next
	}

	// Now prev is at the node before the middle node
	if prev.Next != nil {
		prev.Next = prev.Next.Next // Skip the middle node
	}

	// fmt.Println(dummy)
	return dummy.Next // Return the modified list, skipping the dummy node
}

func main() {
	head := &ListNode{Val: 1,
		Next: &ListNode{Val: 3,
			Next: &ListNode{Val: 4,
				Next: &ListNode{Val: 7,
					Next: &ListNode{Val: 1,
						Next: &ListNode{Val: 2, Next: nil}}}}}}

	result := deleteMiddle(head)
	for result != nil {
		println(result.Val)
		result = result.Next
	}
}
