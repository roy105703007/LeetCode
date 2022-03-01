var dir = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func maxAreaOfIsland(grid [][]int) int {
	queue := [][]int{}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			queue = append(queue, []int{i, j})
			grid[i][j] = 0
			level := 0
			for len(queue) > 0 {
				curX := queue[0][0]
				curY := queue[0][1]
				queue = queue[1:]
				level++
				for k := range dir {
					nextX := curX + dir[k][0]
					nextY := curY + dir[k][1]
					if isInGrid(grid, nextX, nextY) && grid[nextX][nextY] == 1 {
						fmt.Println(nextX, nextY)
						grid[nextX][nextY] = 0
						queue = append(queue, []int{nextX, nextY})
					}
				}
			}
			if level > res {
				res = level
			}
		}
	}
	return res
}
func isInGrid(grid [][]int, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0])
}