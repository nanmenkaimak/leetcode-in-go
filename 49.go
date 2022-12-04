package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	mp := map[string][]string{}

	for _, word := range strs {
		s := strings.Split(word, "")
		sort.Strings(s)
		w := strings.Join(s, "")
		mp[w] = append(mp[w], word)
	}
	for _, word := range mp {
		ans = append(ans, word)
	}
	return ans
}
