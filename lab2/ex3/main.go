package main

import (
	"fmt"
)

func sender(ch chan<- string) {
	ch <- "Message"
}

func main() {
	ch := make(chan string)
	cnt := 2
	for range cnt {
		go sender(ch)
	}

	for range cnt {
		fmt.Println(<-ch)
	}
}
