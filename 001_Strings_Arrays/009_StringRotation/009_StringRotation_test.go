package ch01

import (
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		s1 string
		s2 string
		out   bool
	}{
		{"ab", "ba", true},
		{"a", "b", false},
		{"", "", false},
		{"a", "ab", false},
		{"ab", "a", false},
		{"tokyo", "kyoto", true},
		{"waterbottle", "erbottlewat", true},
	}

	for _, test := range tests {
		res := stringRotation(test.s1, test.s2)
		if res != test.out {
			t.Errorf("Test Failed: \n%#v\n Actual result: \n%v\n", test, res)
		}
	}
}