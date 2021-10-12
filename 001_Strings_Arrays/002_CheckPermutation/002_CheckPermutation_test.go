package ch01

import (
	"testing"
)

func TestCheckPermutations(t *testing.T) {
	tests := []struct {
		s1 string
		s2 string
		out   bool
	}{
		{"ab", "ba", true},
		{"a", "b", false},
		{"", "", true},
		{"a", "ab", false},
		{"ab", "a", false},
		{"tokyo", "kyoto", true},
	}

	for _, test := range tests {
		res := checkPermutation(test.s1, test.s2)
		if res != test.out {
			t.Errorf("Test Failed: \n%#v\n Actual result: \n%v\n", test, res)
		}
	}
}