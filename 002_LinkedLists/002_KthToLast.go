package ch02

/**
 * Given a string, compress it into a readable format
 */

func kthToLastWithLen(n *Node, k int) *Node {
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

func kthToLastWithoutLen(n *Node, k int) *Node {
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
	
	length := 0
	i := 0
	nstart := n
	for n.next != nil{
		n = n.next
		i++
	}
	length = i

	// i should be in the middle of the array at this point, j should be the length
	for i:=0; i < length-k + 1; i++ {
		nstart = nstart.next
	}

	return nstart
}