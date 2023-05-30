package main

import (
	"strings"
	"testing"

	"github.com/xNaCly/leet/go/helper"
)

func Test0468(t *testing.T) {
	examples := map[string]string{
		"12..33.4":                              "neither",
		"2001:0db8:85a3:0:0:8A2E:0370:7334:":    "neither",
		"1.0.1.":                                "neither",
		"1e1.4.5.6":                             "neither",
		"172.16.254.1":                          "ipv4",
		"2001:0db8:85a3:0:0:8A2E:0370:7334":     "ipv6",
		"2001:0db8:85a3:00000:0:8A2E:0370:7334": "neither",
		"256.256.256.256":                       "neither",
	}
	for k, v := range examples {
		helper.AssertEquals(strings.ToLower(validIPAddress(k)), v, k)
	}
}
