package main

import (
	"golang.org/x/tour/pic"
	"math"
)

func Pic(dx, dy int) [][]uint8 {
	my_pic := make([][]uint8, dx)
	for i := 0; i < dx; i++ {
		my_pic[i] = make([]uint8, dy)
		for j := 0; j < dy; j++ {
			if math.Pow(float64(i-100), 2)+math.Pow(float64(j-100), 2) == 2500 {
				my_pic[i][j] = 255
			}
		}
	}
	return my_pic
}

func main() {
	pic.Show(Pic)
}
