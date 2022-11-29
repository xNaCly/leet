package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// returns true if not really a number
func _checkEdge(s string) bool {
    r, _ := regexp.Compile("\\+|-")
    s = r.ReplaceAllString(s, "")
    if strings.HasPrefix(s, "i") || strings.HasPrefix(s, "n"){
        return true
    } else if strings.Count(s, "e") > 1 {
        return true
    }
    return false
}

func isNumber(s string) bool {
	s = strings.ToLower(s)
    if _checkEdge(s) {
        return false
    } else if strings.Contains(s, "e") {
        split := strings.Split(s, "e")
        _, err := strconv.ParseInt(split[1], 10, 64)
        _, err1 := strconv.ParseFloat(split[0], 64)
        return err == nil && err1 == nil
    } else {
        _, err := strconv.ParseFloat(s, 64)
        return err == nil
    }
}

func main() {
	ints := []string{"2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789",
		"abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53", "Infinity", "inf", "NaN", "+Inf", "-inf", "0e1", "7e69e"}
	for _, integer := range ints {
		fmt.Printf("%s: %t\n", integer, isNumber(integer))
	}
}
