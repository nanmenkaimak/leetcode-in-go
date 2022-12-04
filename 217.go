package main

func containsDuplicate(nums []int) bool {
	mp := map[int]bool{}
	cnt := 0
	for _, i := range nums {
		if mp[i] == false {
			mp[i] = true
			cnt++
		} else {
			return true
		}
	}

	if cnt == len(mp) {
		return false
	}
	
	return true
}