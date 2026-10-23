package solution

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	// Visited is a bitmap that tracks lands already visited
	visited := make([][]bool, m)
	for j := range visited {
		visited[j] = make([]bool, n)
	}

	buffer := make([][2]int, 0, m*n)

	counter := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Skip if visited already
			if visited[i][j] {
				continue
			}

			// Skip sea
			if grid[i][j] != '1' {
				visited[i][j] = true
				continue
			}

			// Got a new island
			counter++
			visited[i][j] = true

			// Map out the rest of the island
			// Choose BFS to avoid returning via lands already visited
			// BFS uses a queue
			queue := buffer
			queue = append(queue, [2]int{i, j})

			for len(queue) > 0 {
				// Dequeue a land
				curr := queue[0]
				queue = queue[1:]

				x, y := curr[0], curr[1]

				if grid[x][y] == '1' {
					if x > 0 && !visited[x-1][y] {
						queue = append(queue, [2]int{x - 1, y})
						visited[x-1][y] = true
					}
					if x < m-1 && !visited[x+1][y] {
						queue = append(queue, [2]int{x + 1, y})
						visited[x+1][y] = true
					}
					if y > 0 && !visited[x][y-1] {
						queue = append(queue, [2]int{x, y - 1})
						visited[x][y-1] = true
					}
					if y < n-1 && !visited[x][y+1] {
						queue = append(queue, [2]int{x, y + 1})
						visited[x][y+1] = true
					}
				}
			}
		}
	}

	return counter
}
