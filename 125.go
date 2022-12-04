package main

import (
	"strings"
)

func isPalindrome(s string) bool {
	ans := ""
	for i := 0; i < len(s); i++ {
		if (s[i] >= 65 && s[i] <= 90) || (s[i] >= 97 && s[i] <= 122) || (s[i] >= 48 && s[i] <= 57) {
			ans += string(s[i])
		}
	}
	ans = strings.ToLower(ans)
	j := len(ans) - 1
	for i := 0; i < len(ans); i++ {
		if ans[i] != ans[j] {
			return false
		}
		j--
	}
	return true
}
