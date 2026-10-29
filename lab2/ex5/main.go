package main

import "fmt"

func send(ch chan string) {
	ch <- "Salut din send!"
}

func receive(ch chan string) string {
	msg := <-ch
	return msg
}

func main() {
	ch := make(chan string)
	res := make(chan string)

	go send(ch)
	go func() {
		res <- receive(ch)
	}()

	fmt.Println(<-res)

	// Nu stiam cum altfel sa 'returnez' dintr-un goroutine
}
