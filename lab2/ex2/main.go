package main

import "fmt"

func listen_to_me(channel chan string) {
	msg := <-channel
	fmt.Println(msg + "\nYES BRO, SHUT UP.")
}

func main() {
	ch := make(chan string, 1)

	ch <- "Are you there?"
	listen_to_me(ch)
}
