package main

import (
	"reflect"
	"testing"
)

func Test0066(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{in: []int{1, 2, 3}, out: []int{1, 2, 4}},
		{in: []int{4, 3, 2, 1}, out: []int{4, 3, 2, 2}},
		{in: []int{9}, out: []int{1, 0}},
		{in: []int{4, 9}, out: []int{5, 0}},
		{in: []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}, out: []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7}},
	}
	for _, test := range tests {
		got := plusOne(test.in)
		if !reflect.DeepEqual(got, test.out) {
			t.Errorf("Wanted %v, got %d", test.out, got)
		}
	}
}
