package ch01

import (
	"testing"
)

func TestCheckStringCompression(t *testing.T) {
	tests := []struct {
		s string
		out string
	}{
		{"aabcccccaaa",  "a2b1c5a3"},
		{" aabcccccaaa",  "a2b1c5a3"},
		{"aabcccccaaa ",  "a2b1c5a3"},
		{" aabcccccaaa ",  "a2b1c5a3"},
	}

	for _, test := range tests {
		res := stringCompression(test.s)
		if res != test.out {
			t.Errorf("Test Failed: \n%#v\n Actual result: \n%v\n", test, res)
		}
	}
}