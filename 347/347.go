func topKFrequent(nums []int, k int) []int {
    var bucketMap map[int]int
    bucketMap = make(map[int]int)
    newArray := []int{}
    for i:=0; i < len(nums); i++ {
        _, isExist := bucketMap[nums[i]]
        if !isExist {
            bucketMap[nums[i]] = 1
            newArray = append(newArray, nums[i])
        } else {
            bucketMap[nums[i]] += 1
        }
    }
    newNums := selectSmallest(bucketMap, newArray, 0, len(newArray)-1, len(newArray)-k)
    fmt.Println(newNums)
    return newArray[newNums:]
}

func selectSmallest(m map[int]int, nums []int, l, r, i int) int {
    if l >= r {
        return l
    }
    q := partition(m, nums, l, r)
    if q == i {
        return q
    }
    if q < i {
        return selectSmallest(m, nums, q+1, r, i)
    } else {
        return selectSmallest(m, nums, l, q-1, i)
    }
}

func partition(m map[int]int, nums []int, l, r int) int {
	k := l + rand.Intn(r-l+1)
    nums[k], nums[r] = nums[r], nums[k]
    i := l - 1
    for j := l; j < r; j++ {
        if m[nums[j]] <= m[nums[r]] {
            i++
            nums[i], nums[j] = nums[j], nums[i]
        }
    }
    nums[r], nums[i+1] = nums[i+1], nums[r]
    return i+1
}
