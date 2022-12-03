package main

import "fmt"

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

func main() {
	nums := []int{1, 2, 3, 4, 1}
	fmt.Println(containsDuplicate(nums))
}