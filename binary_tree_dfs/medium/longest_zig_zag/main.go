package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, isLeft bool, length int, maxLength *int) {
	if node == nil {
		return
	}

	// Update the maximum length found so far
	if length > *maxLength {
		*maxLength = length
	}

	// If we are going left, the next step should be right and vice versa
	if isLeft {
		dfs(node.Left, false, length+1, maxLength)
		dfs(node.Right, true, 1, maxLength) // Reset length for right child
	} else {
		dfs(node.Right, true, length+1, maxLength)
		dfs(node.Left, false, 1, maxLength) // Reset length for left child
	}
}

// longestZigZag finds the length of the longest zigzag path in a binary tree.
// A zigzag path is defined as a path that alternates between left and right children.
// The function returns the length of the longest zigzag path found in the tree.
// The path length is defined as the number of edges in the path.
// The function uses depth-first search (DFS) to explore the tree and keep track of the maximum zigzag length.
// The time complexity is O(n), where n is the number of nodes in the tree.
// The space complexity is O(h), where h is the height of the tree due to the recursion stack.
// The function initializes the maximum length to 0 and starts the DFS from the root node.
// It explores both left and right children, alternating the direction and updating the length accordingly.
// The function returns the maximum length found during the DFS traversal.
// The function handles the case where the tree is empty by returning 0.
// The function is designed to work with a binary tree structure defined by the TreeNode type.

func longestZigZag(root *TreeNode) int {
	maxLength := 0

	if root == nil {
		return maxLength
	}

	// Start DFS from the root node, first going left and then right
	dfs(root.Left, true, 1, &maxLength)   // Start with left child
	dfs(root.Right, false, 1, &maxLength) // Start with right child

	return maxLength
}

func main() {
	root := &TreeNode{Val: 1,
		Right: &TreeNode{Val: 1,
			Left: &TreeNode{Val: 1,
				Right: &TreeNode{Val: 1,
					Left: &TreeNode{Val: 1,
						Right: &TreeNode{Val: 1}}}}}}

	result := longestZigZag(root)
	println(result) // Output: 3
}
