func numSquares(n int) int {
	squareList := []int{}
	queue := []int{n}
	level := 0
	visit := make(map[int]bool, n)
	i := 1
	for i*i <= n {
		if i*i == n {
			return 1
		}
		squareList = append(squareList, i*i)
		i++
	}
	for len(queue) > 0 {
		level++
		count := len(queue)
		for c := 0; c < count; c++ {
			cur := queue[0]
			queue = queue[1:]
			for i := range squareList {
				next := cur - squareList[i]
				fmt.Println(next)
				if next < 0 {
					break
				} else if next == 0 {
					return level
				} else if !visit[next] {
					visit[next] = true
					queue = append(queue, next)
				}
			}
		}
	}
	fmt.Println(queue)
	return 0
}