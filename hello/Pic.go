package main

import (
	"code.google.com/p/go-tour/pic"
	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func Pic(dx, dy int) [][]uint8 {
	slce := make([][]uint8, dx)
	for y := 0; y < dy; y++ {
		slce[y] = make([]uint8, dx)
	}
	return slce
}

func main() {
	pic.Show(Pic)
	wc.Test(WordCount)

}
