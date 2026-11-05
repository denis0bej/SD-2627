package main

import (
	"fmt"
	"sync"
)

type Msg struct {
	Value int
	Done  bool // liderul anunta ca a fost ales
}

// realizeaza operatiile interne nodului
func node(id int, in <-chan Msg, out chan<- Msg, isLeader *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	// initital send
	_max := id
	out <- Msg{_max, false}

	// loop
	for {
		// receive
		rec := <-in
		rec_id, done := rec.Value, rec.Done

		// compute
		if !done {
			if rec_id == id {
				// eu sunt lider
				*isLeader = true
			} else if rec_id > _max {
				// am primit lider nou
				_max = rec_id
			} else {
				// pastrez lider actual
				continue
			}
		}

		// send
		if *isLeader || done {
			out <- Msg{_max, true}
			return
		} else {
			out <- Msg{_max, false}
		}
	}
}

// returneaza id-ul liderului
func RingFlooding(ids []int) int {

	n := len(ids)
	channels := make([]chan Msg, n)

	for i := range channels {
		channels[i] = make(chan Msg, 1)
	}

	var wg sync.WaitGroup
	results := []bool{}

	for i, val := range ids {
		results = append(results, false)

		wg.Add(1)
		go node(val, channels[i], channels[(i+1)%n], &results[i], &wg)
	}
	wg.Wait()

	fmt.Println(results)
	for i, val := range results {
		if val == true {
			return i
		}

	}
	return -1
}

func main() {
	values := []int{7, 3, 6, 12, 6, 33, 2}

	fmt.Println(RingFlooding(values))
}
