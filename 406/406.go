func reconstructQueue(people [][]int) [][]int {
	var kiMap map[int][]int
	kiMap = make(map[int][]int)
	kiSort := []int{}
	for i := 0; i < len(people); i++ {
		kiMap[people[i][1]] = append(kiMap[people[i][1]], people[i][0])
	}
	for key, _ := range kiMap {
		kiSort = append(kiSort, key)
		sort.Ints(kiMap[key])
		for i, j := 0, len(kiMap[key])-1; i < j; i, j = i+1, j-1 {
			kiMap[key][i], kiMap[key][j] = kiMap[key][j], kiMap[key][i]
		}
	}
	//sort.Sort(sort.Reverse(kiSort))
	sort.Ints(kiSort)
	newPeople := [][]int{}
	for key := range kiSort {
		//fmt.Println(kiSort[key], kiMap[kiSort[key]])
		for e := range kiMap[kiSort[key]] {
			row1 := []int{kiMap[kiSort[key]][e], kiSort[key]}
			rowPlace := rowIndex(newPeople, row1)
			newPeople = addRow(newPeople, row1, rowPlace)
			//fmt.Println(row1)
			//fmt.Println(rowPlace)
			//newPeople = append(newPeople, row1)
		}
	}
	//fmt.Println(kiMap)
	return newPeople
}

func rowIndex(nums [][]int, row []int) int {
	if len(nums) == 0 {
		return 0
	}
	bigThanRow := 0
	if bigThanRow == row[1] {
		return 0
	}
	for i := range nums {
		if bigThanRow == row[1] {
			return i
		}
		if nums[i][0] < row[0] {
			if bigThanRow == row[1] {
				return i
			} else {
				continue
			}
		}
		if nums[i][0] >= row[0] {
			bigThanRow += 1
		}
	}
	return len(nums)
}
func addRow(nums [][]int, row []int, rowPlace int) [][]int {
	if len(nums) == rowPlace {
		nums = append(nums, row)
	} else {
		nums = append(nums[:rowPlace+1], nums[rowPlace:]...)
		nums[rowPlace] = row
	}
	return nums
}