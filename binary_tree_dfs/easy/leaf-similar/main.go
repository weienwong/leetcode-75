package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1 := []int{}
	leaves2 := []int{}

	collectLeaves(root1, &leaves1)
	collectLeaves(root2, &leaves2)

	if len(leaves1) != len(leaves2) {
		return false
	}

	for i := 0; i < len(leaves1); i++ {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}

	return true

}

func collectLeaves(root *TreeNode, leaves *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*leaves = append(*leaves, root.Val)
		return
	}
	collectLeaves(root.Left, leaves)
	collectLeaves(root.Right, leaves)
}

func main() {
	root1 := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 5,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 4}}},
		Right: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 8}}}

	root2 := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7}},
		Right: &TreeNode{Val: 1,
			Left: &TreeNode{Val: 4},
			Right: &TreeNode{Val: 2,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 8}}}}

	result := leafSimilar(root1, root2)
	println(result) // Output should be true or false based on the leaf similarity
}
