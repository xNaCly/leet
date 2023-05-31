package main

func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}
	ms := make(map[rune]uint)
	mt := make(map[rune]uint)

	sr := []rune(s)
	tr := []rune(t)

	for i := 0; i < len(s); i++ {
		ccs := sr[i]
		cct := tr[i]
		if _, ok := ms[ccs]; ok {
			ms[ccs]++
		} else {
			ms[ccs] = 1
		}
		if _, ok := mt[cct]; ok {
			mt[cct]++
		} else {
			mt[cct] = 1
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

func isAnagramOld(s string, t string) bool {
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
