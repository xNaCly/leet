package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	p := 0

	for i := 0; i < len(flowerbed); i++ {
		prev := 0
		next := 0
		v := flowerbed[i]

		if i-1 >= 0 {
			prev = flowerbed[i-1]
		}

		if i+1 < len(flowerbed) {
			next = flowerbed[i+1]
		}

		if v == 0 && prev == 0 && next == 0 {
			flowerbed[i] = 1
			p++
		}
	}

	return p >= n
}
