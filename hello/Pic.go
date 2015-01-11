package main

import (
	reg "regexp"
	// "code.google.com/p/go-tour/pic"
	// "code.google.com/p/go-tour/wc"
	"fmt"
	"sort"
	// "io"
	ioutl "io/ioutil"
	st "strings"
)

var (
	v_str             = "ab ab ab ab, big and fat fat brown fox been brown"
	slice_s, slice_s1 []string
)

func WordCount(s []string) (map[string]int, string) {
	var (
		max_key   string
		max_value int = 0
	)
	sorted_map := make(map[string]int)

	for i := range s {
		if _, exist := sorted_map[s[i]]; exist {
			sorted_map[s[i]] = sorted_map[s[i]] + 1
			// fmt.Printf("%v - %v\n", sorted_map[s[i]], s[i])
			if sorted_map[s[i]] > max_value {
				// fmt.Printf("max_value - %v, max_key - %v \n", max_value, max_key)
				max_key = s[i]
				max_value = sorted_map[s[i]]
			} // fmt.Printf("sorted_map[%v] - %v \n", i, sorted_map[s[i]])
		} else {
			sorted_map[s[i]] = 1
		}
		// fmt.Println(sorted_map)
	}

	return sorted_map, max_key
	//map[string]int{"x": 1}
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
func f_eval(incoming rune) bool {
	compiled, _ := reg.Compile(`[[:alpha:]]`)
	matched := compiled.MatchString(string(incoming))
	// matched := compiled.MatchReader(incoming)
	/*fmt.Printf("err - %v , incoming - %v , compiled -%v , matched - %v \n", err, string(rune(incoming)),
	compiled, matched)
	*/return !matched
}

func main() {
	// pic.Show(Pic)
	v_str = *(readFile("/home/az/work/GO/data/10.txt.utf-8"))
	// v_str = *(readFile("/home/az/work/GO/data/https://archive.org/stream/Scudry-ArtamneOuLeGrandCyrusTroisimePartie1654/Scudry-ArtamneOuLeGrandCyrusTroisimePartie1654_djvu.txt"))
	slice_s = st.Fields(st.ToLower(v_str))
	for i := range slice_s {
		// fmt.Println(slice_s[i])
		slice_s[i] = st.TrimFunc(slice_s[i], f_eval)
	}
	sort.Strings(sort.StringSlice(slice_s))
	// fmt.Printf("sorted slice_s - %v \n", slice_s)
	// fmt.Printf("sorted_map - %v", WordCount(slice_s))
	ordered_map, the_key := WordCount(slice_s)

	fmt.Printf("The most frequently used word is \"%v\"- %v \n", the_key, ordered_map[the_key])
	fmt.Println("total count of words - ", len(slice_s))
	// wc.Test(WordCount)

}
