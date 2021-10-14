package ch01

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	tests := []struct {
		input string
		out   bool
	}{
		{"", true},
		{"uniqe", true},
		{"uniquee", false},
		{"abcdefghijklmnopqrstuvwxyz", true},
		{"abcdefghijklmnopqrstuvwxyza", false},
	}

	for _, test := range tests {
		res := isUnique(test.input)
		if res != test.out {
			t.Errorf("Test %#v, but got: %v\n", test, res)
		}
	}
}

func TestIsUniqueNoMap(t *testing.T) {
	tests := []struct {
		input string
		out   bool
	}{
		{"", true},
		{"unique", false},
		{"uniquee", false},
		{"abcdefghijklmnopqrstuvwxyz", true},
		{"abcdefghijklmnopqrstuvwxyza", false},
	}

	for _, test := range tests {
		res := isUniqueNoMap(test.input)
		if res != test.out {
			t.Errorf("Test %#v, but got: %v\n", test, res)
		}
	}
}