// o sa presupun ca e two sum (daca e n sum n-am idee cum)
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	vals := [6]int{2, 7, 11, 8, 4, 9}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	s, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("You entered a non-integer value.\n", err)
		os.Exit(1)
	}

	compl := map[int]int{}

	found := false

	for i, v := range vals {
		diff := s - v
		i_compl, ok := compl[v]
		if ok {
			fmt.Println(i_compl, i)
			found = true
		} else {
			compl[diff] = i
		}
	}

	if !found {
		fmt.Printf("Found no pair in %v that summed equals %v.\n", vals, s)
	}

}
