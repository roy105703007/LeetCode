type NumArray struct {
	sumArr []int
}

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return NumArray{sumArr: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left > 0 {
		return this.sumArr[right] - this.sumArr[left-1]
	}
	return this.sumArr[right]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */