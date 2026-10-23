package main

import (
	"fmt"
)

func main() {
	var v [5]uint32 = [5]uint32{8, 1, 3, 2, 7}

	for i := range 5 {
		fmt.Printf("%d ", v[i])
	}
	fmt.Println()

	sorted := false
	for {
		sorted = true
		for i := range 4 {
			if v[i] > v[i+1] {
				v[i], v[i+1] = v[i+1], v[i]
				sorted = false
			}
		}
		if sorted == true {
			break
		}
	}

	for i := range 5 {
		fmt.Printf("%d ", v[i])
	}

}
