package ch01

/**
 * Given two strings, find if one is a rotation of another using IsSubString
 */

func stringRotation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == "" || s2 == "" {
		return false
	}

	sToSearch := s1 + s1

	return isSubstring(sToSearch, s2)
}

func isSubstring(source string, find string) bool {
	for i:=0; i < len(source)-len(find); i++ {
		if source[i:i+len(find)] == find {
			return true
		}
	}
	return false
}