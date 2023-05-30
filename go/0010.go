package main

import (
	"fmt"
	"regexp"
)

func isMatch(s string, p string) bool {
	r, _ := regexp.Compile(fmt.Sprintf("^%s$", p))
	return r.MatchString(s)
}
