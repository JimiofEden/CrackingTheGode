package ch02

/**
 * Given a node, find out if it's a palindrome
 */

func palindrome(n *Node) bool {
	l := n.Len()
	if l == 0 {
		return false
	}

	pMap := make(map[int]int)
	for n != nil {
		pMap[n.value]++
		n = n.next
	}

	singleFlip := false
	for _, v := range pMap {
		if v%2 != 0 {
			if singleFlip {
				return false
			} else {
				singleFlip = true
			}
		}
	}
	
	return true
}