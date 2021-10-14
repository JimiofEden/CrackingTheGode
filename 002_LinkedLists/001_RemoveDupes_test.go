package ch02

import (
	"testing"
	"reflect"
)

func TestCheckRemoveDupes(t *testing.T) {
	tests := []struct {
		l *Node
		out *Node
	}{}

	l0 := NewNode(1)
	tests = append(tests, struct {
		l *Node
		out *Node
	}{
		l: l0,
		out: l0,
	})

	l1 := NewNode(1)
	l1.AddTail(2)
	l1.AddTail(3)
	l1.AddTail(4)
	l1.AddTail(5)
	tests = append(tests, struct {
		l *Node
		out *Node
	}{
		l: l1,
		out: l1,
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
		out *Node
	}{
		l: l2,
		out: l1,
	})

	l3 := NewNode(1)

	tests = append(tests, struct {
		l *Node
		out *Node
	}{
		l: l3,
		out: l3,
	})

	for i, test := range tests {
		removeDupes(test.l)
		if !reflect.DeepEqual(test.l, test.out) {
			t.Errorf("Test %v Failed Expected: \n%#v\n Actual result: \n%v\n", i, test, test.l)
		}
	}
}