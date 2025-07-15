package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bfs(r *TreeNode) []int {
	if r == nil {
		return []int{}
	}

	queue := []*TreeNode{r}
	result := []int{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
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

	result1 := bfs(root1)
	for _, val := range result1 {
		println(val)
	}

}
