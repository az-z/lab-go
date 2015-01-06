package main

import (
	// "code.google.com/p/go-tour/pic"
	// "code.google.com/p/go-tour/wc"
	"fmt"
	"sort"
	// "io"
	ioutl "io/ioutil"
	st "strings"
)

var (
	v_str             = "a a a a big and fat fat brown fox been brown"
	slice_s, slice_s1 []string
)

func WordCount(s []string) map[string]int {
	var (
	// ind int
	)
	sorted_map := make(map[string]int)

	for i := 0; i < len(s); i++ {
		if _, exist := sorted_map[s[i]]; exist {
			sorted_map[s[i]] = sorted_map[s[i]] + 1
			// fmt.Printf("sorted_map[%v] - %v \n", i, sorted_map[s[i]])
		} else {
			sorted_map[s[i]] = 1
		}
		// fmt.Println(sorted_map)
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
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(s string) *string {
	dat, err := ioutl.ReadFile(s)
	check(err)
	str_ := string(dat)
	return &str_
}

func main() {
	// pic.Show(Pic)
	// v_str = *(readFile("/home/az/work/GO/data/10.txt.utf-8"))
	// v_str = *(readFile("/home/az/work/GO/data/https://archive.org/stream/Scudry-ArtamneOuLeGrandCyrusTroisimePartie1654/Scudry-ArtamneOuLeGrandCyrusTroisimePartie1654_djvu.txt"))
	slice_s = st.Fields(v_str)

	sort.Strings(sort.StringSlice(slice_s))
	fmt.Printf("sorted slice_s - %v \n", slice_s)
	fmt.Printf("sorted_map - %v", WordCount(slice_s))
	// wc.Test(WordCount)

}
