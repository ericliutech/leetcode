package solution

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	count := 0

	var exploreIsland func(x, y int)
	exploreIsland = func(x, y int) {
		if grid[x][y] == '1' {
			grid[x][y] = 'x'

			if x < m-1 {
				exploreIsland(x+1, y)
			}
			if y < n-1 {
				exploreIsland(x, y+1)
			}
			if x > 0 {
				exploreIsland(x-1, y)
			}
			if y > 0 {
				exploreIsland(x, y-1)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				exploreIsland(i, j)
			}
		}
	}

	return count
}
