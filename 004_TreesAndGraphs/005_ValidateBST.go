package ch04

func (n *TreeNode) IsBinarySearchTree() bool {
	if n == nil {
		return false
	}

	const MaxUint = ^uint(0) 
	const MinUint = 0 
	const MaxInt = int(MaxUint >> 1) 
	const MinInt = -MaxInt - 1
	
	return isBinarySearchTreeHelper(n, MinInt, MaxInt)
}
func isBinarySearchTreeHelper(n *TreeNode, min int, max int) bool {
	if n == nil {
		return true
	}
	if n.Val < min || n.Val > max {
		return false
	}

	return isBinarySearchTreeHelper(n.Left, min, n.Val-1) && isBinarySearchTreeHelper(n.Right, n.Val+1, max)
}