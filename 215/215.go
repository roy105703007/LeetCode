func findKthLargest(nums []int, k int) int {
    m := len(nums)-k
	return selectSmallest(nums, 0, len(nums)-1, m)
}

func selectSmallest(nums []int, l, r, i int) int {
    if l >= r {
        return nums[l]
    }
    q := partition(nums, l, r)
    if q == i {
        return nums[q]
    }
    if q < i {
        return selectSmallest(nums, q+1, r, i)
    } else {
        return selectSmallest(nums, l, q-1, i)
    }
}

func partition(nums []int, l, r int) int {
	k := l + rand.Intn(r-l+1)
    nums[k], nums[r] = nums[r], nums[k]
    i := l - 1
    for j := l; j < r; j++ {
        if nums[j] <= nums[r] {
            i++
            nums[i], nums[j] = nums[j], nums[i]
        }
    }
    nums[r], nums[i+1] = nums[i+1], nums[r]
    return i+1
}
