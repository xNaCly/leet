package main

import (
	"testing"

	"github.com/xNaCly/leet/go/helper"
)

func Test2299(t *testing.T) {
	tests := map[string]bool{"IloveLe3tcode!": true, "Me+You--IsMyDream": false, "1aB!": false, "11A!A!Aa": false}
	for k, v := range tests {
		helper.AssertEquals(strongPasswordCheckerII(k), v, k)
	}
}
