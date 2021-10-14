package ch01

import (
	"testing"
	"reflect"
)

func TestCheckZeroMatrix(t *testing.T) {
	tests := []struct {
		s [][]int
		out [][]int
	}{
		{[][]int{{1}}, [][]int{{1}}},
		{[][]int{{0}}, [][]int{{0}}},
		{[][]int{{1,2},{3,4}}, [][]int{{1,2},{3,4}}},
		{[][]int{{0,2},{3,4}}, [][]int{{0,0},{0,4}}},
		{[][]int{{1,0},{3,4}}, [][]int{{0,0},{3,0}}},
		{[][]int{{1,2},{0,4}}, [][]int{{0,2},{0,0}}},
		{[][]int{{1,2},{3,0}}, [][]int{{1,0},{0,0}}},
		{[][]int{{1,2,3},{4,5,6},{7,8,9}}, [][]int{{1,2,3},{4,5,6},{7,8,9}}},
		{[][]int{{0,2,3},{4,5,6},{7,8,9}}, [][]int{{0,0,0},{0,5,6},{0,8,9}}},
		{[][]int{{1,2,3},{4,0,6},{7,8,9}}, [][]int{{1,0,3},{0,0,0},{7,0,9}}},
		{[][]int{{1,2,0},{4,5,6},{0,8,9}}, [][]int{{0,0,0},{0,5,0},{0,0,0}}},
		{[][]int{{0,0,0},{4,5,6},{7,8,9}}, [][]int{{0,0,0},{0,0,0},{0,0,0}}},
		{[][]int{{0,2,3},{4,0,6},{7,8,0}}, [][]int{{0,0,0},{0,0,0},{0,0,0}}},
	}

	for _, test := range tests {
		res := zeroMatrix(test.s)
		if !reflect.DeepEqual(res, test.out) {
			t.Errorf("Test Failed: \n%#v\n Actual result: \n%v\n", test, res)
		}
	}
}