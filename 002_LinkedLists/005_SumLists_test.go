package ch02

import (
	"testing"
	"reflect"
)

func TestCheckSumLists(t *testing.T) {
	tests := []struct {
		l *Node
		k *Node
		out *Node
	}{}


	l1 := NewNode(1)
	l1.AddTail(2)
	l1.AddTail(3)

	k1 := NewNode(3)
	k1.AddTail(2)
	k1.AddTail(1)

	r1 := NewNode(4)
	r1.AddTail(4)
	r1.AddTail(4)

	tests = append(tests, struct {
		l *Node
		k *Node
		out *Node
	}{
		l:   l1,
		k:   k1,
		out: r1,
	})

	l2 := NewNode(1)
	l2.AddTail(2)
	l2.AddTail(3)

	k2 := NewNode(1)
	k2.AddTail(2)
	k2.AddTail(5)

	r2 := NewNode(2)
	r2.AddTail(4)
	r2.AddTail(8)

	tests = append(tests, struct {
		l *Node
		k *Node
		out *Node
	}{
		l: 	 l2,
		k: 	 k2,
		out: r2,
	})

	l3 := NewNode(9)
	l3.AddTail(9)
	l3.AddTail(9)

	k3 := NewNode(1)

	r3 := NewNode(0)
	r3.AddTail(0)
	r3.AddTail(0)
	r3.AddTail(1)

	tests = append(tests, struct {
		l *Node
		k *Node
		out *Node
	}{
		l:   l3,
		k:   k3,
		out: r3,
	})

	for i, test := range tests {
		res := sumListReverse(test.l, test.k)
		for test.out.next != nil && res.next != nil {
			if !reflect.DeepEqual(test.out.value, res.value) {
				t.Errorf("Test %v Failed Expected: \n%#v\n Actual result: \n%v\n", i+1, test.out, res)
			}
			test.out = test.out.next
			res = res.next
		}
	}
}