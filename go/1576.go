package main

func modifyString(s string) string {
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
