package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	data := make([]uint8, dx)

	for i, _ := range pic {
		for j := range data {
			data[j] = uint8(i * j)
		}
		pic[i] = data
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
