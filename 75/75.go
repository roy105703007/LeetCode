func sortColors(nums []int)  {
    if len(nums) == 1 {
        return
    }
    l := 0
    r := len(nums)-1
    for i:=0; i<len(nums); i++{
        if i > r {
            break
        }
        if nums[i] == 2 {
            r = getRightPoint(nums, i, r) // nums[r] will only be 0 or 1
            if nums[r] == 0 {
                nums[i], nums[r] = nums[r], nums[i]
                l = getLeftPoint(nums, l, i) 
                fmt.Println(nums[l])
                if nums[l] == 0 {
                    continue
                }
                nums[i], nums[l] = nums[l], nums[i]
                l += 1
                r -= 1
            } else if nums[r] == 1 {
                nums[i], nums[r] = nums[r], nums[i]
                r -= 1
            } else {
                return
            }
        } else if nums[i] == 0 {
            l = getLeftPoint(nums, l, i) 
            fmt.Println(nums[l])
            if nums[l] == 0 {
                continue
            }
            nums[i], nums[l] = nums[l], nums[i]
            l += 1
        }
    }
}
func getRightPoint (nums []int, l, r int) int{
    if r == l {
        return r
    }
    if nums[r] == 2 {
        return getRightPoint(nums, l, r-1)
    }
    return r
}
func getLeftPoint (nums []int, l, r int) int{
    if l == r {
        return l
    }
    if nums[l] == 0 {
        return getLeftPoint(nums, l+1, r)
    }
    return l
}
