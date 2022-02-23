var dir = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func shortestPathBinaryMatrix(grid [][]int) int {
	if len(grid) == 1 && len(grid[0]) == 1 {
		return 1
	}
	if grid[0][0] == 1 || grid[len(grid)-1][len(grid[0])-1] == 1 {
		return -1
	}
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}
	visited[0][0] = true
	dis := make([][]int, len(grid))
	for i := range dis {
		dis[i] = make([]int, len(grid[i]))
	}
	dis[0][0] = 1
	queue := [][]int{}
	queue = append(queue, []int{0, 0})
	for len(queue) > 0 {
		curX := queue[0][0]
		curY := queue[0][1]
		queue = queue[1:]
		for j := 0; j < 8; j++ {
			nextX := curX + dir[j][0]
			nextY := curY + dir[j][1]
			if isInBoard(nextX, nextY, grid) && !visited[nextX][nextY] &&
				grid[nextX][nextY] == 0 {
				if nextX == len(grid)-1 && nextY == len(grid[0])-1 {
					return dis[curX][curY] + 1
				}
				queue = append(queue, []int{nextX, nextY})
				visited[nextX][nextY] = true
				dis[nextX][nextY] = dis[curX][curY] + 1
			}
		}
	}
	return -1
}

func isInBoard(x, y int, grid [][]int) bool {
	return x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0])
}
