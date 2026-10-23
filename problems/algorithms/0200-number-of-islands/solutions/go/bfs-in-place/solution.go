package solution

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	counter := 0

	// Map out the rest of the island
	// Choose BFS to avoid returning via lands already visited
	// BFS uses a queue
	buffer := make([][2]int, 0, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Skip visited cells and sea
			if grid[i][j] != '1' {
				continue
			}

			// Got a new island
			counter++

			queue := buffer
			queue = append(queue, [2]int{i, j})

			// Cheat by reusing grid as visited bitmap
			grid[i][j] = '2'

			for len(queue) > 0 {
				// Dequeue a land
				curr := queue[0]
				queue = queue[1:] // no pointer - no risk of memory leak

				x, y := curr[0], curr[1]

				if x > 0 && grid[x-1][y] == '1' {
					queue = append(queue, [2]int{x - 1, y})
					grid[x-1][y] = '2'
				}
				if x < m-1 && grid[x+1][y] == '1' {
					queue = append(queue, [2]int{x + 1, y})
					grid[x+1][y] = '2'
				}
				if y > 0 && grid[x][y-1] == '1' {
					queue = append(queue, [2]int{x, y - 1})
					grid[x][y-1] = '2'
				}
				if y < n-1 && grid[x][y+1] == '1' {
					queue = append(queue, [2]int{x, y + 1})
					grid[x][y+1] = '2'
				}
			}
		}
	}

	return counter
}
