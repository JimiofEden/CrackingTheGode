package ch04

import (
	"testing"
)

func TestCheckBinarySearchTree(t *testing.T) {
	tests := []struct {
		l *TreeNode
		out bool
	}{}

	l1 := NewTreeNode(2)
	l1.AddLeftLeaf(1)
	l1.AddRightLeaf(3)
	tests = append(tests, struct {
		l *TreeNode
		out bool
	}{
		l: l1,
		out: true,
	})

	l2 := NewTreeNode(1)
	l2.AddLeftLeaf(2)
	l2.AddRightLeaf(3)
	tests = append(tests, struct {
		l *TreeNode
		out bool
	}{
		l: l2,
		out: false,
	})

	l3 := NewTreeNode(6)
	l3.AddLeftLeaf(2)
	l3.Left.AddLeftLeaf(1)
	l3.Left.AddRightLeaf(8)
	l3.AddRightLeaf(12)
	l3.Right.AddLeftLeaf(10)
	l3.Right.AddRightLeaf(14)
	tests = append(tests, struct {
		l *TreeNode
		out bool
	}{
		l: l3,
		out: false,
	})

	l4 := NewTreeNode(6)
	l4.AddLeftLeaf(2)
	l4.Left.AddLeftLeaf(1)
	l4.AddRightLeaf(12)
	l4.Right.AddLeftLeaf(10)
	l4.Right.AddRightLeaf(14)
	tests = append(tests, struct {
		l *TreeNode
		out bool
	}{
		l: l4,
		out: true,
	})

	for i, test := range tests {
		res := test.l.IsBinarySearchTree()
		if res != test.out {
			t.Errorf("Test %v Failed Expected: \n%#v\n Actual result: \n%v\n", i+1, test.out, res)
		}
	}
}