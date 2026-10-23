package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	n, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("You entered a non-integer value.\n", err)
	} else {
		var prim bool = true
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				prim = false
				break
			}
		}
		if prim {
			fmt.Printf("You entered a prime number. (%d)\n", n)
		} else {
			fmt.Printf("You entered a non-prime number. (%d)\n", n)
		}
	}

}
