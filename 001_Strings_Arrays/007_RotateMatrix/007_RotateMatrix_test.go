package ch01

import (
	"fmt"
	"testing"
	"reflect"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		s [][]int
		out [][]int
	}{
		{[][]int{{1}}, [][]int{{1}}},
		{[][]int{{1,2},{3,4}}, [][]int{{3,1},{4,2}}},
		{[][]int{{1,2,3},{4,5,6},{7,8,9}}, [][]int{{7,4,1},{8,5,2},{9,6,3}}},
	}

	for _, test := range tests {
		res := rotateMatrix(test.s)
		fmt.Println(res)
		if !reflect.DeepEqual(res, test.out) {
			t.Errorf("Test Failed: \n%#v\n Actual result: \n%v\n", test, res)
		}
	}
}