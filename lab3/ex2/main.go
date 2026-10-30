package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type LamportClock struct{ counter int }

// incrementeaza contorul
func (l *LamportClock) Tick() {
	l.counter += 1
}
func (l *LamportClock) Send() {
	l.Tick()
}
func (l *LamportClock) Receive(ts int) {
	l.counter = max(l.counter, ts)
	l.Tick()
}

type Message struct{ sender, ts int }
type Event struct {
	process          int
	internal_counter int
	kind             string // "internal", "send", "recv"
}

// TODO: retuneaza true daca a precede b
func less(a, b Event) bool {
	if a.internal_counter != b.internal_counter {
		return a.internal_counter < b.internal_counter
	}
	return a.process < b.process
}

func routine(id int, targetch chan Message, logch chan Event) {

	clock := LamportClock{0}
	for range 20 { // lets say each process has (max) 10 events

		// Handle possible incoming message
		select {
		case message := <-targetch: // RECEIVE
			clock.Receive(message.ts)
			logch <- Event{id, clock.counter, "recv"}
		case <-time.After(time.Millisecond * 100):
		}

		// Work
		choice := rand.Float32()
		switch {
		case choice > 0.3: // SEND
			clock.Send()
			targetch <- Message{id, clock.counter}
			logch <- Event{id, clock.counter, "send"}
		case choice >= 0.2: // INTERNAL
			//time.Sleep(time.Millisecond * time.Duration(rand.Intn(15)*10))
			clock.Tick()
			logch <- Event{id, clock.counter, "internal"}
		default: // NO EVENT
		}
	}

}

func main() {
	ch12, ch23, ch31 := make(chan Message, 1), make(chan Message, 1), make(chan Message, 1)
	chlog := make(chan Event)
	go routine(1, ch12, chlog)
	go routine(2, ch23, chlog)
	go routine(3, ch31, chlog)

	logs := []Event{}

	running := true

	for running {
		select {
		case log := <-chlog:
			// fmt.Println("Main:", log)
			logs = append(logs, log)
		case <-time.After(time.Second * 1):
			fmt.Println("Main has timed out! [1s]")
			running = false
		}
	}

	fmt.Printf("Unsorted event list is:\n%v\n", logs)

	sort.Slice(logs, func(i, j int) bool {
		return less(logs[i], logs[j])
	})

	fmt.Printf("Sorted event list is:\n%v\n", logs)

}
