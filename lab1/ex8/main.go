package main

import "fmt"

func main() {
	a_d := map[string]int{
		"Ana Popescu":      1,
		"Mihai Ionescu":    2,
		"Bogdan Radu":      1,
		"Elena Dumitrescu": 3,
		"Cristian Vasile":  2,
		"Ioana Enache":     1,
	}

	d_a := map[int][]string{}

	for k, v := range a_d {
		_, ok := a_d[k]
		if !ok { // Not yet in d_a
			d_a[v] = []string{k}
		} else {
			d_a[v] = append(d_a[v], k)
		}
	}

	for k, v := range d_a {
		fmt.Printf("\nDepartamentul %v:\n", k)
		for _, dep := range v {
			fmt.Printf("   - %v\n", dep)
		}
	}

	fmt.Println()

}
