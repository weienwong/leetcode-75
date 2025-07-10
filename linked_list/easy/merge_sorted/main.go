package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var mergedHead, current *ListNode

	if list1.Val < list2.Val {
		mergedHead = list1
		list1 = list1.Next
	} else {
		mergedHead = list2
		list2 = list2.Next
	}
	current = mergedHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return mergedHead
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}

	mergedList := mergeTwoLists(list1, list2)
	for mergedList != nil {
		println(mergedList.Val)
		mergedList = mergedList.Next
	}

	list1 = &ListNode{Val: 2, Next: &ListNode{Val: 4}}
	list2 = &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	mergedList = mergeTwoLists(list1, list2)
	for mergedList != nil {
		println(mergedList.Val)
		mergedList = mergedList.Next
	}
}
