package ch02

/**
 * Given a string, compress it into a readable format
 */

func kthToLast(n *Node, k int) *Node {
	l := n.Len()
	if k > l {
		return NewNode(-1)
	}
	if k == 0 {
		k = 1
	}
	if k == 1 && l == 1 {
		return n
	}
	
	for i:=0; i < l-k; i++ {
		n = n.next
	}

	return n
}