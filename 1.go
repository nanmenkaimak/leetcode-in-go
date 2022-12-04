package main

func twoSum(nums []int, target int) []int {
    var ans = []int{}
    for i := 0; i < len(nums); i++{
        var n = target - nums[i]
        for j := i + 1; j < len(nums); j++{
            if n == nums[j]{
                ans = append(ans,i,j)
                break
            }
        }
    }
    return ans
}