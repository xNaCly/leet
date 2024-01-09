package main

import "testing"

func Test2243(t *testing.T) {
	tests := []struct {
		in  string
		k   int
		out string
	}{
		{
			in:  "11111222223",
			k:   3,
			out: "135",
		},
		{
			in:  "00000000",
			k:   3,
			out: "000",
		},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			got := digitSum(test.in, test.k)
			if got != test.out {
				t.Errorf("Expected %q, got %q", test.out, got)
			}
		})
	}
}
