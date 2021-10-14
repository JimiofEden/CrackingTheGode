package ch01

import (
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		s1 string
		out bool
	}{
		{"testing this", false},
		{"taco cat", true},
		{"ttaacco", true},
		{"warsaw was raw", true},
	}

	for _, test := range tests {
		res := palindromePermutationN(test.s1)
		if res != test.out {
			t.Errorf("Test Failed: \n%#v\n Actual result: \n%v\n", test, res)
		}
	}

	for _, test := range tests {
		res := palindromePermutationNLogN(test.s1)
		if res != test.out {
			t.Errorf("Test Failed: \n%#v\n Actual result: \n%v\n", test, res)
		}
	}
}