func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		v, ok := hashMap[target-nums[i]]
		if ok {
			return []int{v, i}
		} else {
			hashMap[nums[i]] = i
		}
	}
	return nil
}