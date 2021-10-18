package ch04

import "math"

func MinimalTree(arr []int) *TreeNode{

	if len(arr) == 0 {
		return nil
	}

	if len(arr) == 1 {
		return NewTreeNode(arr[0])
	}

	if len(arr) == 2 {
		v := NewTreeNode(arr[0])
		if arr[0] > arr[1] {
			v.Left = NewTreeNode(arr[1])
			return v
		}
		v.Right = NewTreeNode(arr[1])
		return v
	}

	n := math.Ceil(float64(len(arr)/2))
	i := int(n)
	var root = NewTreeNode(arr[i])

	if i != 0 {
		root.Left = MinimalTree(arr[:i])
	}
	if i < len(arr)-1 {
		root.Right = MinimalTree(arr[i+1:])
	}

	return root
}
