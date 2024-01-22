package main

import "testing"

func Test0771(t *testing.T) {
	tests := []struct {
		j   string
		s   string
		out int
	}{}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			got := numJewelsInStones(test.j, test.s)
			if got != test.out {
				t.Errorf("Wanted %d, got %d", test.out, got)
			}
		})
	}
}
