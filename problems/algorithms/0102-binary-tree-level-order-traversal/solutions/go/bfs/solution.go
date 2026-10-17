package solution

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)

    if root == nil {
        return result
    }

    queue := []*TreeNode{root}

    for queueSize := len(queue); queueSize > 0; queueSize = len(queue) {
        levelResult := make([]int, queueSize)

        for idx, node := range queue {
            levelResult[idx] = node.Val
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        clear(queue[:queueSize])
        queue = queue[queueSize:]

        result = append(result, levelResult)
    }

    return result
}
