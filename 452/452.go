func findMinArrowShots(points [][]int) int {
    
    points = mergeSort(points)
    end := points[0][len(points[0])-1]
    count := 1
    for i := 1; i < len(points); i++{
        if points[i][0] > end {
            end = points[i][len(points[i])-1]
            count += 1
        }
    }
    return count
}

func mergeSort (nums [][]int) [][]int{
    if len(nums) < 2 {
        return nums
    }
    l := mergeSort(nums[:len(nums)/2])
    r := mergeSort(nums[len(nums)/2:])
    return merge(l, r)
}

func merge(l, r [][]int) [][]int{
    final := [][]int{}
    i := 0
    j := 0
    for i<len(l) && j<len(r) {
        if l[i][len(l[i])-1] < r[j][len(r[j])-1] {
            final = append(final, l[i])
            i++
        } else {
            final = append(final, r[j])
            j++
        }
    }
    for ; i < len(l); i++ {
        final = append(final, l[i])
    }
    for ; j < len(r); j++ {
        final = append(final, r[j])
    }
    return final
}
