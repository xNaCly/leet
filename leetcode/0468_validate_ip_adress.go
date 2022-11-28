package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func containsChars(word string) bool {
	return regexp.MustCompile(`[a-zA-z]`).MatchString(word)
}

func validIPAddress(queryIP string) string {
	isIp4 := false
	isIp6 := false

	if strings.Contains(queryIP, ".") {
		if !(strings.HasPrefix(queryIP, ".") || strings.HasSuffix(queryIP, ".")) {
			splitQuery := strings.Split(queryIP, ".")
			for _, e := range splitQuery {
				if len(splitQuery) != 4 || containsChars(e) || len(e) == 0 || e[0] == ' ' {
					isIp4 = false
					break
				} else {
					// convert e to integer
					num, err := strconv.Atoi(e)

					if err == nil && (0 <= num && num <= 255) && !(len(e) > 1 && []rune(e)[0] == rune('0')) {
						isIp4 = true
					} else {
						isIp4 = false
						break
					}
				}
			}
		}
	} else {
		if !(strings.HasPrefix(queryIP, ":") || strings.HasSuffix(queryIP, ":")) {
			splitQuery := strings.Split(queryIP, ":")
			for _, e := range splitQuery {
				if len(splitQuery) != 8 || len([]rune(e)) > 4 {
					isIp6 = false
					break
				}

				// check if char is valid
				num, err := strconv.ParseUint(e, 16, 64)

				if err != nil || num > 65535 {
					isIp6 = false
					break
				}

				isIp6 = true
			}
		}
	}

	if isIp4 == true {
		return "IPv4"
	} else if isIp6 == true {
		return "IPv6"
	} else {
		return "Neither"
	}
}

func main() {
	examples := []string{
		"12..33.4",                              // neither
		"2001:0db8:85a3:0:0:8A2E:0370:7334:",    // neither
		"1.0.1.",                                // neither
		"1e1.4.5.6",                             // neither
		"172.16.254.1",                          // ipv4
		"2001:0db8:85a3:0:0:8A2E:0370:7334",     // ipv6
		"2001:0db8:85a3:00000:0:8A2E:0370:7334", // neither
		"256.256.256.256"}                       // neither
	for _, e := range examples {
		fmt.Println(validIPAddress(e))
	}
}
