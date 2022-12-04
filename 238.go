package main

func productExceptSelf(nums []int) []int {
	size := len(nums)
	ans := make([]int, size)
	fromBegin := make([]int, size)
	fromLast := make([]int, size)
	fromBegin[0] = 1
	fromLast[0] = 1
	for i := 1; i < size; i++ {
		fromBegin[i] = nums[i-1] * fromBegin[i-1]
		fromLast[i] = nums[size-i] * fromLast[i-1]
	}
	for i := 0; i < size; i++ {
		ans[i] = fromBegin[i] * fromLast[size-i-1]
	}
	return ans
}
