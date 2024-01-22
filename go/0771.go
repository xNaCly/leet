package main

// ver1
func numJewelsInStones1(jewels string, stones string) int {
	j := map[rune]struct{}{}
	for _, r := range jewels {
		j[r] = struct{}{}
	}
	c := 0
	for _, r := range stones {
		if _, ok := j[r]; ok {
			c++
		}
	}
	return c
}

func numJewelsInStones(jewels string, stones string) int {
	c := 0
outer:
	for _, r := range stones {
		for _, j := range jewels {
			if r == j {
				c++
				continue outer
			}
		}
	}
	return c
}
