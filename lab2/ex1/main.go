package main

import "fmt"

func helloMe(channel chan string, me string) {
	channel <- "Salut " + me + "!\n"
}

func main() {
	ch := make(chan string)
	go helloMe(ch, "Denis")

	fmt.Print(<-ch)
}
