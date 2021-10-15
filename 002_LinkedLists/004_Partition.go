package ch02

/**
 * Given two nodes, delete one from the other
 */

func partition(n *Node, k int) *Node {
	l := n.Len()
	if l == 0 {
		return n
	}

	var lessThanNode *Node = nil
	var greaterThanNode *Node = nil

	for n != nil {

		if n.value < k {
			if lessThanNode == nil {
				lessThanNode = NewNode(n.value)
			} else {
				lessThanNode.AddTail(n.value)
			}
		} else {
			if greaterThanNode == nil {
				greaterThanNode = NewNode(n.value)
			} else {
				greaterThanNode.AddTail(n.value)
			}
		}

		n = n.next
	}

	for greaterThanNode != nil {
		lessThanNode.AddTail(greaterThanNode.value)
		greaterThanNode = greaterThanNode.next
	}
	
	return lessThanNode
}