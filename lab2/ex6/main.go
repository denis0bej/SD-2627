package main

import (
	"fmt"
	"strconv"
	"time"
)

func robot(id int, ch chan string, distance int) {
	for range distance {

	}
	ch <- "Robot " + strconv.Itoa(id) + " a ajuns primul!"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go robot(67, ch1, 1000)
	go robot(2, ch2, 1000)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Cam lenti amandoi...")
	}

	/*
		Proof ca merge
		dbejan@LOQ-DENIS:~/SD-2627$ go run lab2/ex6/main.go
		Robot 2 a ajuns primul!
		dbejan@LOQ-DENIS:~/SD-2627$ go run lab2/ex6/main.go
		Robot 2 a ajuns primul!
		dbejan@LOQ-DENIS:~/SD-2627$ go run lab2/ex6/main.go
		Robot 67 a ajuns primul!
	*/
}
