package main

import (
    "fmt"
    "strings"
)

const MIN_LEN = 8
const ALPHABET_LOWER = "abcdefghijklmnopqrstuvwxyz"
const ALPHABET_UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const SYMBOLS = "!@#$%^&*()-+"
const DIGITS = "0123456789"

func strongPasswordCheckerII(password string) bool {
    if len(password) < MIN_LEN {
        return false
    }

    containsLowerCase := false
    containsUpperCase := false
    containsDigit := false
    containsSymbol := false

    for i, cc := range password {
        if i < len(password) - 1 {
            nc := rune(password[i+1])

            if cc == nc {
                return false
            }
        }

         
        if containsLowerCase == false {
            containsLowerCase = strings.ContainsRune(ALPHABET_LOWER, cc)
        }

        if containsUpperCase == false {
            containsUpperCase = strings.ContainsRune(ALPHABET_UPPER, cc)
        }

        if containsDigit == false {
            containsDigit = strings.ContainsRune(DIGITS, cc)
        }

        if containsSymbol == false {
            containsSymbol = strings.ContainsRune(SYMBOLS, cc)
        }
    }

    return containsLowerCase && containsUpperCase && containsDigit && containsSymbol
}

func main() {
    tests := []string{"IloveLe3tcode!", "Me+You--IsMyDream", "1aB!", "11A!A!Aa"}
    for _, test := range tests {
        fmt.Println(strongPasswordCheckerII(test))
    }
}
