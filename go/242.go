package main


func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}
	ms := make(map[rune]uint)
	mt := make(map[rune]uint)

	for _, cc := range s {
		if _, ok := ms[cc]; ok {
			ms[cc]++
		} else {
			ms[cc] = 1
		}
	}

	for _, cc := range t {
		if _, ok := mt[cc]; ok {
			mt[cc]++
		} else {
			mt[cc] = 1
		}
	}

	r := true
	for key, value := range ms {
		if v, ok := mt[key]; ok {
			if v != value {
				return false
			}
		} else {
			return false
		}
	}
	return r
}
