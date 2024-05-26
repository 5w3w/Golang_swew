package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		s := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			s[x] = uint8((x + y)/2)
		}
		slice[y] = s
	}
	return slice
}

	func main() {
		pic.Show(Pic)
	}
