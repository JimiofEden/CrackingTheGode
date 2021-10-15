package ch02

import "math"

/**
 * Given two nodes, add them together in reverse
 */

func sumListReverse(n *Node, k *Node) *Node {
	l := n.Len()
	if l == 0 {
		return n
	}

	if k == nil {
		return n
	}
	if n == nil {
		return k
	}

	runningSum := 0
	for i:=0; n != nil && k != nil; i++ {
		runningSum += (n.value + k.value) * int(math.Pow(10, float64(i)))
		if n != nil {
			n = n.next
		}
		if k != nil {
			k = k.next
		}
	}

	r := NewNode(runningSum%10)
	f := r
	runningSum /= 10

	for i:=1; runningSum > 0; i++ {
		f.next = NewNode(runningSum%10)
		runningSum /= 10
		f = f.next
	}
	
	return r
}