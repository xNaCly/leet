package main

import (
	"strings"
	"testing"

	"github.com/xNaCly/leet/go/helper"
)

func Test1576(t *testing.T) {
	var i = []struct {
		in  string
		exp string
	}{
		{
			in:  "?zs",
			exp: "azs",
		},
		{
			in:  "ubv?w",
			exp: "ubvaw",
		},
		{
			in:  "j?qg??b",
			exp: "jaqgacb",
		},
	}

	for _, s := range i {
		t.Run(s.in, func(t *testing.T) {
			t.Run("first", func(t *testing.T) {
				res := modifyStringFirst(s.in)
				helper.AssertFalse(strings.ContainsRune(res, '?'))
				helper.AssertEqualsGeneric(res, s.exp)
			})
			t.Run("second", func(t *testing.T) {
				res := modifyStringSecond(s.in)
				helper.AssertFalse(strings.ContainsRune(res, '?'))
				helper.AssertEqualsGeneric(res, s.exp)
			})
			t.Run("third", func(t *testing.T) {
				res := modifyStringThird(s.in)
				helper.AssertFalse(strings.ContainsRune(res, '?'))
				helper.AssertEqualsGeneric(res, s.exp)
			})
			t.Run("forth", func(t *testing.T) {
				res := modifyStringForth(s.in)
				helper.AssertFalse(strings.ContainsRune(res, '?'))
				helper.AssertEqualsGeneric(res, s.exp)
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

func BenchmarkForth(b *testing.B) {
	for size, v := range inputSizes {
		b.Run(size, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				modifyStringForth(strings.Repeat("j?qg??b", v))
			}
		})
	}
}
