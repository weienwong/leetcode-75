package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, targetSum int, currentSum int, pathCount map[int]int) int {
	if node == nil {
		return 0
	}

	// Update the current sum
	currentSum += node.Val

	// Count the number of paths that end at this node and sum to targetSum
	pathCount[currentSum]++
	// Calculate the number of valid paths that sum to targetSum
	numPaths := pathCount[currentSum-targetSum]

	// Recurse left and right children
	numPaths += dfs(node.Left, targetSum, currentSum, pathCount)
	numPaths += dfs(node.Right, targetSum, currentSum, pathCount)

	// Decrement the count of the current sum in the map before returning
	pathCount[currentSum]--

	return numPaths
}

func pathSum(root *TreeNode, targetSum int) int {
	pathCount := make(map[int]int)
	pathCount[0] = 1 // Initialize with a sum of 0 to handle paths that start from the root

	return dfs(root, targetSum, 0, pathCount)
}

func main() {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: -2,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: -3,
			Right: &TreeNode{
				Val: 11,
			},
		},
	}

	targetSum := 8
	result := pathSum(root, targetSum)
	println(result) // Output should be the number of paths that sum to targetSum
}
