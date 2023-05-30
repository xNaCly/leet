package main

import (
	"testing"

	"github.com/xNaCly/leet/go/helper"
)

func Test0050(t *testing.T) {
	tests := map[float64][]float64{
		1024.0:   {2, 10},
		9.26100:  {2.1, 3},
		0.25:     {2.0, -2},
		-1.0:     {-1.0, 2147483648},
		1.0:      {-1.0, -2147483648},
		4.0:      {-2.0, 2},
		54.83508: {0.44894, -5},
		8e-05:    {0.99999, 948688},
	}
	for r, ks := range tests {
		helper.EqualsFloat(r, myPow(ks[0], int(ks[1])))
	}
}
