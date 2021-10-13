package ch01

import (
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		s1 string
		s2 string
		out bool
	}{
		{"abcd", "abcd",  true},
		{"abcd", "abc",  true},
		{"abd", "abcd",  false},
		{"abcd", "abbd",  true},
		{"abca", "abcd",  true},
		{"abcda", "abc",  false},
		{"abc", "abcda",  false},
		{"abcd", "acbd",  false},
		{"", "a",  true},
		{"", "ab",  false},
		{"a", "",  true},
		{"ab", "",  false},
		{"ab", "c",  false},
		{"c", "ab",  false},
	}

	for _, test := range tests {
		res := oneAway(test.s1, test.s2)
		if res != test.out {
			t.Errorf("Test Failed: \n%#v\n Actual result: \n%v\n", test, res)
		}
	}
}