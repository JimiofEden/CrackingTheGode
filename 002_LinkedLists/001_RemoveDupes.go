package ch02

/**
 * Given a string, compress it into a readable format
 */

func removeDupes(l *SinglyLinkedList) {
	n := l.Len()
	if n <= 1 {
		return
	}

	valArr := make(map[int]bool)
	
	
}