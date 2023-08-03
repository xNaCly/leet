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
		t.Run(s, func(t *testing.T) {
			t.Run("first", func(t *testing.T) {
				helper.AssertFalse(strings.ContainsRune(modifyStringFirst(s), '?'))
			})
			t.Run("second", func(t *testing.T) {
				helper.AssertFalse(strings.ContainsRune(modifyStringSecond(s), '?'))
			})
			t.Run("third", func(t *testing.T) {
				helper.AssertFalse(strings.ContainsRune(modifyStringThird(s), '?'))
			})
		})
	}
}

var inputSizes = map[string]int{
	"size10":     10,
	"size100":    100,
	"size1000":   1000,
	"size10000":  10000,
	"size100000": 100000,
}

func BenchmarkFirst(b *testing.B) {
	for size, v := range inputSizes {
		b.Run(size, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				modifyStringFirst(strings.Repeat("j?qg??b", v))
			}
		})
	}
}

func BenchmarkSecond(b *testing.B) {
	for size, v := range inputSizes {
		b.Run(size, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				modifyStringSecond(strings.Repeat("j?qg??b", v))
			}
		})
	}
}

func BenchmarkThird(b *testing.B) {
	for size, v := range inputSizes {
		b.Run(size, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				modifyStringThird(strings.Repeat("j?qg??b", v))
			}
		})
	}
}
