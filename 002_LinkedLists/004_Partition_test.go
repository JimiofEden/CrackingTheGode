package ch02

import (
	"testing"
	"reflect"
)

func TestCheckPartition(t *testing.T) {
	tests := []struct {
		l *Node
		k int
		out *Node
	}{}

	k := 3

	l1 := NewNode(1)
	l1.AddTail(4)
	l1.AddTail(3)
	l1.AddTail(4)
	l1.AddTail(3)
	r1 := NewNode(1)
	r1.AddTail(4)
	r1.AddTail(3)
	r1.AddTail(4)
	r1.AddTail(3)
	tests = append(tests, struct {
		l *Node
		k int
		out *Node
	}{
		l: l1,
		k: k,
		out: r1,
	})

	k = 5

	l2 := NewNode(1)
	l2.AddTail(2)
	l2.AddTail(7)
	l2.AddTail(15)
	l2.AddTail(6)
	l2.AddTail(3)
	l2.AddTail(4)
	l2.AddTail(5)
	l2.AddTail(8)
	l2.AddTail(9)
	r2 := NewNode(1)
	r2.AddTail(2)
	r2.AddTail(3)
	r2.AddTail(4)
	r2.AddTail(7)
	r2.AddTail(15)
	r2.AddTail(6)
	r2.AddTail(5)
	r2.AddTail(8)
	r2.AddTail(9)
	tests = append(tests, struct {
		l *Node
		k int
		out *Node
	}{
		l: l2,
		k: k,
		out: r2,
	})

	for i, test := range tests {
		res := partition(test.l, test.k)
		for test.out.next != nil && res.next != nil {
			if !reflect.DeepEqual(test.out.value, res.value) {
				t.Errorf("Test %v Failed Expected: \n%#v\n Actual result: \n%v\n", i+1, test.out, res)
			}
			test.out = test.out.next
			res = res.next
		}
	}
}