package main

import (
	"fmt"
	"log"
	"regexp"
)

func isMatch(s string, p string) bool {
    r, _ := regexp.Compile(fmt.Sprintf("^%s$", p))
	return r.MatchString(s)
}

func main() {
	in := [][]string{{"aa", "a"}, {"aa", "a*"}, {"ab", ".*"}}
	out := []bool{false, true, true}
	for i := 0; i < len(in); i++ {
		f := isMatch(in[i][0], in[i][1])
		if f != out[i] {
			log.Fatalf("error at i: %d | got='%t', wanted='%t' | in='%s:%s'\n", i, f, out[i], in[i][0], in[i][1])
		}
	}
}
