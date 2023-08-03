package main

import "strings"

func modifyStringForth(s string) string {
	ls := len(s)
	b := make([]byte, ls)
	for i, r := range []byte(s) {
		if r == '?' {
			var rr byte
			for (i > 0 && rr+97 == b[i-1]) ||
				(ls > i+1 && rr+97 == s[i+1]) {
				rr++
			}
			r = rr + 97
		}
		b[i] = r
	}
	return string(b)
}

func modifyStringThird(s string) string {
	ls := len(s)
	b := make([]byte, 0)
	for i, r := range s {
		by := byte(r)
		lb := len(b)
		if r == '?' {
			var rr byte
			for (lb > 0 && rr+97 == b[lb-1]) ||
				(ls > i+1 && rr+97 == s[i+1]) {
				rr++
			}
			by = rr + 97
		}
		b = append(b, by)
	}
	return string(b)
}

func modifyStringSecond(s string) string {
	b := strings.Builder{}
	for i, r := range s {
		if r == '?' {
			rr := 0
			for (b.Len() > 0 && rr+97 == int(b.String()[b.Len()-1])) ||
				(len(s) > i+1 && rr+97 == int(s[i+1])) {
				rr++
			}
			b.WriteRune(rune(rr + 97))
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func modifyStringFirst(s string) string {
	b := strings.Builder{}
	for i, r := range s {
		if r == '?' {
			rr := 0
			for (b.Len() > 0 && alphabet[rr] == b.String()[b.Len()-1]) ||
				(len(s) > i+1 && alphabet[rr] == s[i+1]) {
				rr++
			}
			b.WriteByte(alphabet[rr])
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}
