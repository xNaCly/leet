package main

import "testing"

func Test1017(t *testing.T) {
	tests := []struct {
		in  int
		out string
	}{
		{2, "110"},
		{3, "111"},
		{4, "100"},
	}
	for _, test := range tests {
		out := baseNeg2(test.in)
		if test.out != out {
			t.Errorf("Failure, wanted %q, got %q", test.out, out)
		}
	}
}
