package ch02

import (
	"testing"
	"reflect"
)

func TestCheckRemoveDupes(t *testing.T) {
	tests := []struct {
		l *SinglyLinkedList
		out *SinglyLinkedList
	}{}

	l1 := NewSinglyLinkedList(1)
	l1.AddTail(2)
	l1.AddTail(3)
	l1.AddTail(4)
	l1.AddTail(5)
	tests = append(tests, struct {
		l *SinglyLinkedList
		out *SinglyLinkedList
	}{
		l: l1,
		out: l1,
	})

	l2 := NewSinglyLinkedList(1)
	l2.AddTail(2)
	l2.AddTail(2)
	l2.AddTail(3)
	l2.AddTail(4)
	l2.AddTail(4)
	l2.AddTail(5)

	tests = append(tests, struct {
		l *SinglyLinkedList
		out *SinglyLinkedList
	}{
		l: l2,
		out: l1,
	})

	l3 := NewSinglyLinkedList(1)

	tests = append(tests, struct {
		l *SinglyLinkedList
		out *SinglyLinkedList
	}{
		l: l3,
		out: l3,
	})

	for _, test := range tests {
		removeDupes(test.l)
		if !reflect.DeepEqual(test.l, test.out) {
			t.Errorf("Test Failed: \n%#v\n Actual result: \n%v\n", test, test.l)
		}
	}
}