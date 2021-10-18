package ch04

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func NewTreeNode(n int) *TreeNode {
	l := new(TreeNode)
	l.Val = n
	return l
}

func (n *TreeNode) AddLeftLeaf(v int) {
	if n.Left == nil {
		n.Left = NewTreeNode(v)
	} else {
		n.Left.Val = v
	}

}

func (n *TreeNode) AddRightLeaf(v int) {
	if n.Right == nil {
		n.Right = NewTreeNode(v)
	} else {
		n.Right.Val = v
	}
}

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

func (n *TreeNode) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}