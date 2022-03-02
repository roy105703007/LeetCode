var dir = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == byte(49) {
				count++
				searchIsland(&grid, i, j)
			}
		}
	}
	return count
}

func searchIsland(grid *[][]byte, x, y int) {
	(*grid)[x][y] = byte(48)
	for i := range dir {
		nextX := x + dir[i][0]
		nextY := y + dir[i][1]
		if nextX >= 0 && nextY >= 0 && nextX < len(*grid) && nextY < len((*grid)[0]) && (*grid)[nextX][nextY] == byte(49) {
			searchIsland(grid, nextX, nextY)
		}
	}
}