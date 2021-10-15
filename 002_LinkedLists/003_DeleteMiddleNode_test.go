package ch02

import (
	"testing"
	"reflect"
)

func TestCheckDeleteMiddleNode(t *testing.T) {
	tests := []struct {
		l *Node
		k *Node
		out *Node
	}{}

	k := NewNode(3)

	l0 := NewNode(1)
	tests = append(tests, struct {
		l *Node
		k *Node
		out *Node
	}{
		l: l0,
		k: k,
		out: l0,
	})

	tests = append(tests, struct {
		l *Node
		k *Node
		out *Node
	}{
		l: l0,
		k: k,
		out: l0,
	})

	l1 := NewNode(1)
	l1.AddTail(2)
	l1.AddTail(3)
	l1.AddTail(4)
	l1.AddTail(5)
	r1 := l1
	r1.next.next = l1.next.next.next
	tests = append(tests, struct {
		l *Node
		k *Node
		out *Node
	}{
		l: l1,
		k: k,
		out: r1,
	})

	l2 := NewNode(1)
	l2.AddTail(2)
	l2.AddTail(2)
	l2.AddTail(3)
	l2.AddTail(4)
	l2.AddTail(4)
	l2.AddTail(5)
	r2 := l2
	r2.next.next.next = l2.next.next.next.next
	tests = append(tests, struct {
		l *Node
		k *Node
		out *Node
	}{
		l: l2,
		k: k,
		out: r2,
	})

	for i, test := range tests {
		deleteMiddleNode(test.l, test.k)
		for test.out.next != nil && test.l.next != nil {
			if !reflect.DeepEqual(test.out.value, test.l.value) {
				t.Errorf("Test %v Failed Expected: \n%#v\n Actual result: \n%v\n", i+1, test.out, test.l)
			}
			test.out = test.out.next
			test.l = test.l.next
		}
	}
}