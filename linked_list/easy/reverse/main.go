package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		nextTemp := current.Next // Store the next node
		current.Next = prev      // Reverse the link
		prev = current           // Move prev to current
		current = nextTemp       // Move to the next node
	}

	return prev // New head of the reversed list
}

func main() {
	head := &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 5}}}}}
	reversedHead := reverseList(head)

	for node := reversedHead; node != nil; node = node.Next {
		println(node.Val)
	}

}
