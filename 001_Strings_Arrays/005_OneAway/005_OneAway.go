package ch01

/**
 * Given two strings, will determine if a string is only one edit away
 */

func oneAway(s1 string, s2 string) bool {
	// If they are the same string then this is true
	if s1 == s2 {
		return true
	}

	// If a string is more than one character differnt in length, then it is automatically more than one step away
	if myAbs(len(s1) - len(s2)) > 1 {
		return false
	}

	// iterate through both strings
	diffCount := 0
	for i := 0; i < myMax(len(s1), len(s2)); i++ {
		if i > len(s1) - 1 {
			diffCount++
			continue
		}
		if i > len(s2) - 1 {
			diffCount++
			continue
		}

		if s1[i] != s2[i] {
			diffCount++
		}
	}

	if diffCount > 1 {
		return false
	}

	return true
}

func myAbs(n int) int {
	if n < 0 {
		return -1*n
	}
	return n
}

func myMax(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}