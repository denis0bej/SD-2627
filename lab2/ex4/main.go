package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func f(ch chan string) {
	n := rand.Int32N(10)
	time.Sleep(time.Millisecond * time.Duration(n*250))
	ch <- "Im done!"
}

func main() {
	var ch chan string = make(chan string, 1)

	go f(ch)

	fmt.Println(<-ch)
}
