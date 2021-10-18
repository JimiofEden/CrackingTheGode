package ch04

import (
	"testing"
)

func TestCheckMinimalTree(t *testing.T) {
	tests := []struct {
		l []int
		out *TreeNode
	}{}

	l1 := []int{1, 2, 3}
	o1 := NewTreeNode(2)
	o1.AddLeftLeaf(1)
	o1.AddRightLeaf(3)
	tests = append(tests, struct {
		l []int
		out *TreeNode
	}{
		l: l1,
		out: o1,
	})

	l2 := []int{1, 2, 3, 4, 5}
	o2 := NewTreeNode(3)
	o2.AddLeftLeaf(1)
	o2.Left.AddRightLeaf(2)
	o2.AddRightLeaf(4)
	o2.Right.AddRightLeaf(5)
	tests = append(tests, struct {
		l []int
		out *TreeNode
	}{
		l: l2,
		out: o2,
	})

	l3 := []int{1, 2, 3, 4, 5, 6}
	o3 := NewTreeNode(4)
	o3.AddLeftLeaf(2)
	o3.Left.AddLeftLeaf(1)
	o3.Left.AddRightLeaf(3)
	o3.AddRightLeaf(5)
	o3.Right.AddRightLeaf(6)
	tests = append(tests, struct {
		l []int
		out *TreeNode
	}{
		l: l3,
		out: o3,
	})

	for i, test := range tests {
		res := MinimalTree(test.l)

		if !IsSameTree(res, test.out) {
			t.Errorf("Test %v Failed Expected: \n%#v\n Actual result: \n%v\n", i+1, test.out, res)
		}
	}
}