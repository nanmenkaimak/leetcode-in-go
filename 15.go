package main

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				for nums[left] == nums[left - 1] && left < right {
					left++
				}
			}
		}
	}
	return res
}