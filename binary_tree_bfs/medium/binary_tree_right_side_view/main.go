package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	rightSide := []int{}

	for len(queue) > 0 {
		levelSize := len(queue)
		var rightMost *TreeNode

		for i := 0; i < levelSize; i++ {
			node := queue[i]
			if node != nil {
				rightMost = node
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
		}

		if rightMost != nil {
			rightSide = append(rightSide, rightMost.Val)
		}

		queue = queue[levelSize:] // Remove processed level nodes
	}

	return rightSide
}

func main() {
	root := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3,
			Right: &TreeNode{Val: 4}},
	}

	result := rightSideView(root)
	for _, val := range result {
		println(val) // Output: 1, 3, 4
	}
}
