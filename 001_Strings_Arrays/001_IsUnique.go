package ch01

/**
 * Will determine if a string has only unique characters
 */

func isUnique(s string) bool {
	sMap := make(map[rune]int)

	for _, r := range s {
		sMap[r]++
		if sMap[r] > 1 {
			return false
		}
	}

	return true
}

func isUniqueNoMap(s string) bool {

	for i:=0; i < len(s); i++ {
		for j:=i+1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}

	return true
}