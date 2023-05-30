package main

import (
	"strings"
	"testing"

	"github.com/xNaCly/leet/go/helper"
)

func Test1576(t *testing.T) {

	var i = []string{
		"?zs",
		"ubv?w",
		"j?qg??b",
	}

	for _, s := range i {
		helper.AssertFalse(strings.ContainsRune(modifyString(s), '?'))
	}
}
