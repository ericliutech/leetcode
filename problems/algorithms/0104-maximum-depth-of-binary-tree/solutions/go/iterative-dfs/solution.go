package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Trace struct {
	Node   *TreeNode
	Height int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := []Trace{
		{
			Node:   root,
			Height: 1,
		},
	}
	maxh := 1

	for len(stack) > 0 {
		// pop one
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// update maxh
		maxh = max(maxh, curr.Height)

		// push children
		if curr.Node.Right != nil {
			stack = append(stack, Trace{
				Node:   curr.Node.Right,
				Height: curr.Height + 1,
			})
		}
		if curr.Node.Left != nil {
			stack = append(stack, Trace{
				Node:   curr.Node.Left,
				Height: curr.Height + 1,
			})
		}
	}

	return maxh
}
