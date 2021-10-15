package ch02

/**
 * Given two nodes, delete one from the other
 */

func deleteMiddleNode(n *Node, k *Node) {
	l := n.Len()
	if l == 0 {
		return
	}

	prev := n
	for n.next != nil {
		if n.value == k.value && n.next == k.next {
			prev.next = n.next
			n = prev
			break;
		}
		prev = n
		n = n.next
	}
}