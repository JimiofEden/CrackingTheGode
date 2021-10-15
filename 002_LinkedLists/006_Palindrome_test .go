package ch02

import (
	"testing"
	"reflect"
)

func TestCheckPalindrome(t *testing.T) {
	tests := []struct {
		l *Node
		out bool
	}{}


	l1 := NewNode(1)
	l1.AddTail(2)
	l1.AddTail(3)

	tests = append(tests, struct {
		l *Node
		out bool
	}{
		l:   l1,
		out: false,
	})

	l2 := NewNode(1)
	l2.AddTail(2)
	l2.AddTail(1)

	tests = append(tests, struct {
		l *Node
		out bool
	}{
		l: 	 l2,
		out: true,
	})

	l3 := NewNode(1)

	tests = append(tests, struct {
		l *Node
		out bool
	}{
		l:   l3,
		out: true,
	})

	var l4 *Node

	tests = append(tests, struct {
		l *Node
		out bool
	}{
		l:   l4,
		out: false,
	})

	l5 := NewNode(1)
	l5.AddTail(2)
	l5.AddTail(1)
	l5.AddTail(1)

	tests = append(tests, struct {
		l *Node
		out bool
	}{
		l:   l5,
		out: false,
	})

	l6 := NewNode(1)
	l6.AddTail(2)
	l6.AddTail(2)
	l6.AddTail(1)

	tests = append(tests, struct {
		l *Node
		out bool
	}{
		l:   l5,
		out: true,
	})

	for i, test := range tests {
		res := palindrome(test.l)
		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("Test %v Failed Expected: \n%#v\n Actual result: \n%v\n", i+1, test.out, res)
		}
	}
}