package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := m[n]; !ok {
			m[n] = struct{}{}
		} else {
			return true
		}
	}
	return false
}
