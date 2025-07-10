package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// goodNodes counts the number of "good" nodes in a binary tree.
// A node is considered "good" if it is the highest value node on the path from the root to that node.
// The function uses a depth-first search (DFS) approach to traverse the tree and count the good nodes.
// It starts from the root and keeps track of the maximum value seen so far on the path to each node.
// If a node's value is greater than or equal to the maximum value seen so far, it is counted as a good node.
// The function returns the total count of good nodes in the tree.
func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, root.Val)
}

// dfs is a helper function that performs a depth-first search on the binary tree.
// It takes the current node and the maximum value seen so far on the path to that node.
func dfs(node *TreeNode, maxVal int) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.Val >= maxVal {
		count = 1         // Count this node as a good node
		maxVal = node.Val // Update the maximum value seen so far
	}

	// Recursively count good nodes in the left and right subtrees
	count += dfs(node.Left, maxVal)
	count += dfs(node.Right, maxVal)

	return count
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}

	result := goodNodes(root)
	println(result) // Output: 4
	root = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
	}

	result = goodNodes(root)
	println(result) // Output: 3
}
