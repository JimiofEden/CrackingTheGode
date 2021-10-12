package ch01

/**
 * Will determine if one string is a permutation of another
 */

func checkPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	sMap := make(map[rune]int)

	for _, r := range s1 {
		sMap[r]++
	}

	for _, r := range s2 {
		sMap[r]--
		if sMap[r] < 0 {
			return false
		}
	}

	return true
}

