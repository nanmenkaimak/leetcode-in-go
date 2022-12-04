package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	mp := map[int]int{}
	for _, v := range nums {
		mp[v]++
	}
	forsoring := []int{}

	for key := range mp {
		forsoring = append(forsoring, key)
	}

	sort.SliceStable(forsoring, func(i, j int) bool {
		return mp[forsoring[i]] > mp[forsoring[j]]
	})

	ans := []int{}
	for i := 0; i < k; i++ {
		ans = append(ans, forsoring[i])
	}
	return ans
}