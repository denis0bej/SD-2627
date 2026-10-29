package main

import "fmt"

func this_listens(channel chan string) {
	msg := <-channel
	fmt.Println(msg + "\nYES BRO, SHUT UP.")
}

func main() {
	ch := make(chan string, 1)

	ch <- "Are you there?"
	this_listens(ch)
}
