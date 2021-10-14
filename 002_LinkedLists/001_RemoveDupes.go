package ch02

/**
 * Given a string, compress it into a readable format
 */

func removeDupes(l *Node) {
	n := l.Len()
	if n <= 1 {
		return
	}

	valArr := make(map[int]bool)
	prev := l
	for l.next != nil {
		if _, ok := valArr[l.value]; ok {
			// Remove node
			prev.next = l.next
		} else {
			valArr[l.value] = true
			prev = l
		}
		l = l.next
	}
}