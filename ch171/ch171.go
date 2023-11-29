package ch171

import (
	"math"
)

func TitleToNumber(columnTitle string) int {
	b := []byte(columnTitle)
	var y float64 = 0
	sum := 0
	for i := len(b) - 1; i >= 0; i-- {
		num := b[i] - 'A' + 1
		sum += int(num) * int(math.Pow(26, y))
		y++
	}
	return sum
}
