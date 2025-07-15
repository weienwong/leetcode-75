package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isValidBSTHelper(root, nil, nil)
}

func isValidBSTHelper(node *TreeNode, min *int, max *int) bool {
	if node == nil {
		return true
	}

	if (min != nil && node.Val <= *min) || (max != nil && node.Val >= *max) {
		return false
	}

	return isValidBSTHelper(node.Left, min, &node.Val) && isValidBSTHelper(node.Right, &node.Val, max)
}

func main() {
	root := &TreeNode{Val: 2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3}}

	isValid := isValidBST(root)
	println(isValid) // Output: true

	root = &TreeNode{Val: 5,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 6}}}

	isValid = isValidBST(root)
	println(isValid) // Output: false
}
