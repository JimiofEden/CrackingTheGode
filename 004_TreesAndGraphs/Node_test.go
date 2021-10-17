package ch04

import (
	"testing"
)

func TestCheckBinaryTree(t *testing.T) {
	tests := []struct {
		l *Node
		out bool
	}{}

	l1 := NewNode(2)
	l1.AddLeftLeaf(1)
	l1.AddRightLeaf(3)
	tests = append(tests, struct {
		l *Node
		out bool
	}{
		l: l1,
		out: true,
	})

	l2 := NewNode(1)
	l2.AddLeftLeaf(2)
	l2.AddRightLeaf(3)
	tests = append(tests, struct {
		l *Node
		out bool
	}{
		l: l2,
		out: true,
	})

	for i, test := range tests {
		res := test.l.IsBinaryTree()
		if res != test.out {
			t.Errorf("Test %v Failed Expected: \n%#v\n Actual result: \n%v\n", i+1, test.out, res)
		}
	}
}

func TestCheckBinarySearchTree(t *testing.T) {
	tests := []struct {
		l *Node
		out bool
	}{}

	l1 := NewNode(2)
	l1.AddLeftLeaf(1)
	l1.AddRightLeaf(3)
	tests = append(tests, struct {
		l *Node
		out bool
	}{
		l: l1,
		out: true,
	})

	l2 := NewNode(1)
	l2.AddLeftLeaf(2)
	l2.AddRightLeaf(3)
	tests = append(tests, struct {
		l *Node
		out bool
	}{
		l: l2,
		out: false,
	})

	l3 := NewNode(6)
	l3.AddLeftLeaf(2)
	l3.children[0].AddLeftLeaf(1)
	l3.children[0].AddRightLeaf(8)
	l3.AddRightLeaf(12)
	l3.children[1].AddLeftLeaf(10)
	l3.children[1].AddRightLeaf(14)
	tests = append(tests, struct {
		l *Node
		out bool
	}{
		l: l3,
		out: false,
	})

	for i, test := range tests {
		res := test.l.IsBinarySearchTree(0)
		if res != test.out {
			t.Errorf("Test %v Failed Expected: \n%#v\n Actual result: \n%v\n", i+1, test.out, res)
		}
	}
}