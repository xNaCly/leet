package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		t := strings.Split(str, "")
		sort.Strings(t)
		k := strings.Join(t, "")
		if _, ok := m[k]; !ok {
			m[k] = []string{}
		}
		m[k] = append(m[k], str)
	}

	r := make([][]string, len(m))
	i := 0
	for _, a := range m {
		r[i] = a
		i++
	}
	return r
}
