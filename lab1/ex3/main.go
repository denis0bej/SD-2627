package main

import (
	"fmt"
)

func main() {
	v := [9]int{1, 2, 3, 4, 5, 6, 7, 6, 7}
	slice := []int{}

	fmt.Println(v)

	for _, x := range v {
		if x%2 == 0 {
			slice = append(slice, x)
		}
	}

	fmt.Println(slice)

}
