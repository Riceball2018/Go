package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	var ys [][]uint8
	ys = make([][]uint8, dy)
	
	for y := range ys {
		var xs []uint8
		xs = make([]uint8, dx)
		ys[y] = xs
	}
	
	for y := range ys {
		for x := range ys[y] {
			ys[y][x] = (uint8(x) * uint8(y)) / 2	
		}
	}
	
	return ys
}

func main() {
	pic.Show(Pic)
}