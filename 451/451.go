func frequencySort(s string) string {
	if s == "" {
		return ""
	}
	var bucketMap map[int]int
	bucketMap = make(map[int]int)
	newArray := []int{}
	bs := []byte(s)
	for i := 0; i < len(bs); i++ {
		_, isExist := bucketMap[int(bs[i])]
		if !isExist {
			bucketMap[int(bs[i])] = 1
			newArray = append(newArray, int(bs[i]))
		} else {
			bucketMap[int(bs[i])] += 1
		}
	}
	newArray = quickSort(bucketMap, newArray, 0, len(newArray)-1)
	var newString string
	for i := len(newArray) - 1; i >= 0; i-- {
		for j := 0; j < bucketMap[newArray[i]]; j++ {
			newString = newString + string(newArray[i])
		}
	}
	return newString
}
func partition(m map[int]int, arr []int, low, high int) ([]int, int) {
	pivot := m[arr[high]]
	i := low
	for j := low; j < high; j++ {
		if m[arr[j]] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
func quickSort(m map[int]int, arr []int, l, r int) []int {
	if l < r {
		var p int
		arr, p = partition(m, arr, l, r)
		arr = quickSort(m, arr, l, p-1)
		arr = quickSort(m, arr, p+1, r)
	}
	return arr
}