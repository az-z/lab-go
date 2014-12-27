package main

import (
	"fmt"
	//	"time"
)

func f_test(i int) int {
	a := i

}
func main() {

	a := []int{10, 20, 30}
	b := []int{1, 2, 3}
	c := append(a, b...)
	fmt.Println("c - ", c)
	f_test(100)
	i = 2
	x = []int{3, 5, 7}
	for i, x[i] = range x {
		// set i, x[2] = 0, x[0]
		fmt.Println("i - ", i)
		fmt.Println("x[i] - ", x[i])
	}

	fmt.Println("When's Saturday?")
	// today := time.Now().Weekday()
	// switch time.Saturday {
	// case today + 0:
	// 	fmt.Println("Today.")
	// case today + 1:
	// 	fmt.Println("Tomorrow.")
	// case today + 2:
	// 	fmt.Println("In two days.")
	// default:
	// 	fmt.Println("Too far away.")
	// }

	// for i := 0; i < 3; {
	// 	fmt.Println("i - ", i)
	// 	i = i + 1
	// }
}
