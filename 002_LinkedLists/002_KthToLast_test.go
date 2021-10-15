package ch02

import (
	"testing"
	"reflect"
)

func TestCheckKthToLast(t *testing.T) {
	tests := []struct {
		l *Node
		k int
		out *Node
	}{}

	l0 := NewNode(1)
	tests = append(tests, struct {
		l *Node
		k int
		out *Node
	}{
		l: l0,
		k: 0,
		out: l0,
	})

	tests = append(tests, struct {
		l *Node
		k int
		out *Node
	}{
		l: l0,
		k: 1,
		out: l0,
	})

	l1 := NewNode(1)
	l1.AddTail(2)
	l1.AddTail(3)
	l1.AddTail(4)
	l1.AddTail(5)
	tests = append(tests, struct {
		l *Node
		k int
		out *Node
	}{
		l: l1,
		k: 1,
		out: NewNode(5),
	})
	tests = append(tests, struct {
		l *Node
		k int
		out *Node
	}{
		l: l1,
		k: 2,
		out: NewNode(4),
	})
	tests = append(tests, struct {
		l *Node
		k int
		out *Node
	}{
		l: l1,
		k: 3,
		out: NewNode(3),
	})

	l2 := NewNode(1)
	l2.AddTail(2)
	l2.AddTail(2)
	l2.AddTail(3)
	l2.AddTail(4)
	l2.AddTail(4)
	l2.AddTail(5)

	tests = append(tests, struct {
		l *Node
		k int
		out *Node
	}{
		l: l2,
		k: 5,
		out: NewNode(2),
	})

	for i, test := range tests {
		res := kthToLastWithLen(test.l, test.k)
		if !reflect.DeepEqual(test.out.value, res.value) {
			t.Errorf("Test %v Failed Expected: \n%#v\n Actual result: \n%v\n", i+1, test.out.value, res.value)
		}
	}

	for i, test := range tests {
		res := kthToLastWithoutLen(test.l, test.k)
		if !reflect.DeepEqual(test.out.value, res.value) {
			t.Errorf("Test %v Failed Expected: \n%#v\n Actual result: \n%v\n", i+1, test.out.value, res.value)
		}
	}
}