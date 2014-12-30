package main

import (
	// "code.google.com/p/go-tour/pic"
	// "code.google.com/p/go-tour/wc"
	"fmt"
	//"sort"
	st "strings"
)

var (
	v_str = "a a a a big and fat fat brown fox been brown"
)

func WordCount(s []string) map[string]int {
	var (
	//	exist bool //= true
	)
	sorted_map := make(map[string]int)

	for i := 0; i < len(s); i++ {
		if _, exist := sorted_map[s[i]]; exist {
			sorted_map[s[i]] = sorted_map[s[i]] + 1
			//fmt.Println("sorted_map[i] - ", sorted_map[s[i]])
		} else {
			sorted_map[s[i]] = 1
		}
	}

	return sorted_map
	//map[string]int{"x": 1}
}

func split_string(str string) {
	fmt.Println(st.Fields(str))
}

func Pic(dx, dy int) [][]uint8 {
	slce := make([][]uint8, dx)
	for y := 0; y < dy; y++ {
		slce[y] = make([]uint8, dx)
	}
	return slce
}

func main() {
	// pic.Show(Pic)
	slice_s := st.Fields(v_str)
	//sorted_slice := sort.Sort(slice_s)

	//WordCount(slice_s)
	fmt.Printf("slice_s - %v \n", slice_s)
	fmt.Printf("sorted_map - %v", WordCount(slice_s))
	// wc.Test(WordCount)

}
