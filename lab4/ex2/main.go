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
func node(id int, n int, ins []chan Msg, outs []chan Msg, isLeader *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	// initital send
	rounds := 0
	_max := id
	broadcast(Msg{_max, false}, outs)

	// loop
	for {
		// receive
		rec_id, done := analyze(listen(ins))
		rounds += 1

		// compute
		if !done {
			if rec_id > _max {
				// am primit lider nou
				_max = rec_id
			}
			if rounds >= n && _max == id {
				// dupa n runde eu sunt maximul din sistem
				*isLeader = true
			}
		}

		// send
		if *isLeader || done {
			broadcast(Msg{_max, true}, outs)
			for _, out := range outs {
				close(out)
			}
			return
		} else {
			broadcast(Msg{_max, false}, outs)
		}
	}
}

func broadcast(msg Msg, targets []chan Msg) {
	for _, target := range targets {
		target <- msg
	}
}
func listen(sources []chan Msg) []Msg {
	messages := []Msg{}

	for _, source := range sources {
		if msg, ok := <-source; ok {
			messages = append(messages, msg)
		} else {
			messages = append(messages, Msg{Value: -1, Done: true})
		}
	}

	return messages
}
func analyze(messages []Msg) (int, bool) {
	done := false
	max := -1
	for _, msg := range messages {
		if msg.Done == true {
			done = true
		}
		if msg.Value > max {
			max = msg.Value
		}
	}
	return max, done
}

// returneaza id-ul liderului
func Flooding(ids []int) int {
	edges := [][]int{{2, 3}, {0}, {6}, {5}, {0}, {1}, {4}}
	ins := [][]chan Msg{{}, {}, {}, {}, {}, {}, {}}
	outs := [][]chan Msg{{}, {}, {}, {}, {}, {}, {}}

	for i, e := range edges {
		for _, v := range e {
			ch := make(chan Msg, 49)
			outs[i] = append(outs[i], ch)
			ins[v] = append(ins[v], ch)
		}
	}
	n := len(ids)

	var wg sync.WaitGroup
	results := []bool{}

	for i := range n {
		results = append(results, false)

		wg.Add(1)
		go node(ids[i], n, ins[i], outs[i], &results[i], &wg)
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
	values := []int{1, 2, 10, 4, 5, 6, 7}

	fmt.Println(Flooding(values))
}
