package main

import (
	"fmt"
)

func main() {
	myMap := map[string]int{
		"Andrei": 190,
		"Maria":  165,
		"Ion":    182,
		"Elena":  172,
		"George": 178,
		"Ana":    165,
	}

	// Expected: Ana <-> Maria, Elena, George, Ion, Andrei

	values := []int{}
	keys := []string{}
	for k, v := range myMap {
		values = append(values, v)
		keys = append(keys, k)
	}

	sorted := false
	for {
		sorted = true
		for i := range len(keys) - 1 {
			if values[i] > values[i+1] {
				values[i], values[i+1] = values[i+1], values[i]
				keys[i], keys[i+1] = keys[i+1], keys[i]
				sorted = false
			}
		}
		if sorted == true {
			break
		}
	}

	fmt.Println(keys)
}
