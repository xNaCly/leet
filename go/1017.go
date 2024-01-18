package main

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	buf := make([]byte, 0, 8)
	for n != 0 {
		r := n % -2
		n /= -2
		if r < 0 {
			r += 2
			n++
		}
		buf = append(buf, byte(r)+'0')
	}
	t := make([]byte, len(buf))
	j := 0
	for i := len(buf) - 1; i >= 0; i-- {
		t[j] = buf[i]
		j++
	}
	return string(t)
}
