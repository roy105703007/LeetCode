func ladderLength(beginWord string, endWord string, wordList []string) int {
	endIndex := -1
	wordMap := make(map[int][]int, len(wordList))
	wordLen := len(wordList[0])
	visit := make(map[int]int, len(wordList))
	queue := []int{}
	for i := range wordList {
		for j := range wordList {
			if i == j {
				continue
			}
			flag := 0
			for k := 0; k < wordLen; k++ {
				if flag > 1 {
					break
				}
				if wordList[i][k] != wordList[j][k] {
					flag++
				}
			}
			if flag == 1 {
				wordMap[i] = append(wordMap[i], j)
			}
		}
	}
	level := 0
	for i := range wordList {
		if beginWord == wordList[i] {
			queue = append(queue, i)
			visit[i] = 1
		}
		if endWord == wordList[i] {
			endIndex = i
		}
	}
	if len(queue) == 0 {
		for i := range wordList {
			flag := 0
			for j := range wordList[0] {
				if flag > 1 {
					break
				}
				if wordList[i][j] != beginWord[j] {
					flag++
				}
			}
			if flag == 1 {
				if i == endIndex {
					return 2
				}
				queue = append(queue, i)
				visit[i] = 1
			}
		}
		level = 1
	}
	for len(queue) > 0 {
		fmt.Println(visit)
		level++
		curLen := len(queue)
		for i := 0; i < curLen; i++ {
			next := queue[0]
			queue = queue[1:]
			for j := range wordMap[next] {
				if wordMap[next][j] == endIndex {
					return level + 1
				}
				if visit[wordMap[next][j]] == 0 {
					visit[wordMap[next][j]] = level
					queue = append(queue, wordMap[next][j])
				}
			}
		}
	}
	fmt.Println(wordMap)
	return 0
}