func solveNQueens(n int) [][]string {
	limited := make([][]int, n)
	res := [][]string{}
	queens := make([][]bool, n)
	for i := range limited {
		limited[i] = make([]int, n)
		queens[i] = make([]bool, n)
	}
	dfs(n, 0, &queens, &limited, &res)
	fmt.Println(res)
	return res
}

func dfs(n, count int, queens *[][]bool, limited *[][]int, res *[][]string) {
	if count == n {
		allStr := []string{}
		for i := 0; i < n; i++ {
			str := ""
			for k := 0; k < n; k++ {
				if (*queens)[i][k] {
					str += "Q"
				} else {
					str += "."
				}
			}
			allStr = append(allStr, str)
		}
		*res = append(*res, allStr)
		return
	}
	for j := 0; j < n; j++ {
		if (*limited)[count][j] == 0 {
			(*queens)[count][j] = true
			newLimited(count, j, count+1, limited)
			dfs(n, count+1, queens, limited, res)
			removeLimited(count+1, limited)
			(*queens)[count][j] = false
		}
	}
}

func newLimited(i, j, n int, limited *[][]int) {
	for row := 0; row < len(*limited); row++ {
		for col := 0; col < len(*limited); col++ {
			if (*limited)[row][col] == 0 {
				if row == i || col == j || (float32(row-i)/float32(col-j) == 1) || (float32(row-i)/float32(col-j) == -1) {
					(*limited)[row][col] = n
				}
			}
		}
	}
}

func removeLimited(n int, limited *[][]int) {
	for row := 0; row < len(*limited); row++ {
		for col := 0; col < len(*limited); col++ {
			if (*limited)[row][col] == n {
				(*limited)[row][col] = 0
			}
		}
	}
}