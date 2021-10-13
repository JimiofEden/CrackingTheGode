package ch01

import (
	"testing"
)

func TestCheckPermutations(t *testing.T) {
	tests := []struct {
		s1 string
		i1 int
		out string
	}{
		{"testing this", 12, "testing%20this"},
		{"testing this   ", 12,"testing%20this"},
		{"testing this again  ", 18, "testing%20this%20again"},
		{" testing this again  ", 20, "%20testing%20this%20again%20"},
	}

	for _, test := range tests {
		res := urlify(test.s1, test.i1)
		if res != test.out {
			t.Errorf("Test Failed: \n%#v\n Actual result: \n%v\n", test, res)
		}
	}
}