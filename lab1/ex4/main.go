package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	slice := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("before: ", slice)

	fmt.Print("Index of target: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	target, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("You entered a non-integer value.\n", err)
		os.Exit(1)
	}

	slice = append(slice[:target], slice[target+1:]...)

	fmt.Println("after: ", slice)
}
