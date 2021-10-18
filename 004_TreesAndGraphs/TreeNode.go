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

func (n *TreeNode) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
    
    for p != nil && q != nil {
        if p.Val != q.Val {
            return false
        }
        return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
    }
    if p != nil || q != nil {
        return false
    }
    return true
}