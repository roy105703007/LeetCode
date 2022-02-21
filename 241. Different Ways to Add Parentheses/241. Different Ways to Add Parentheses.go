package main

import "strconv"

func diffWaysToCompute(expression string) []int {
	nums := []int{}
	for i := 0; i < len(expression); i++ {
		if string(expression[i]) == "-" || string(expression[i]) == "+" ||
			string(expression[i]) == "*" {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for j := range left {
				for k := range right {
					switch string(expression[i]) {
					case "+":
						nums = append(nums, left[j]+right[k])
					case "-":
						nums = append(nums, left[j]-right[k])
					case "*":
						nums = append(nums, left[j]*right[k])
					}
				}
			}
		}
	}
	if len(nums) == 0 {
		i1, err := strconv.Atoi(expression)
		if err == nil {
			nums = append(nums, i1)
		}
	}
	return nums
}
