package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slce := make([][]uint8, dx)
	for y := 0; y < dy; y++ {
		slce[y] = make([]uint8, dx)
	}
	return slce
}

func main() {
	pic.Show(Pic)
}
