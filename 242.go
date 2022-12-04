package main

func isAnagram(s string, t string) bool {
	mpS := map[byte]int{}
	mpT := map[byte]int{}
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		mpS[s[i]]++;
		mpT[t[i]]++;
	}

	for key := range mpS {
		if mpS[key] != mpT[key] {
			return false
		}
	}
	return true
}