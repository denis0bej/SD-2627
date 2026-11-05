package main

import (
	"fmt"
	"sync"
)

type Msg struct {
	Value float64
}

// realizeaza operatiile interne nodului
func node(initialVote float64, maxRounds int, ins []chan Msg, outs []chan Msg, finalVote *int, wg *sync.WaitGroup) {
	defer wg.Done()

	val := initialVote
	rounds := 0

	// initial send
	broadcast(Msg{val}, outs)

	// loop
	for {
		// receive
		messages := listen(ins)
		rounds += 1

		// compute
		val = analyze(val, messages)
		if rounds >= maxRounds {
			if val > 0.5 {
				*finalVote = 1
			} else {
				*finalVote = 0
			}
			return
		}

		// send
		broadcast(Msg{val}, outs)
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
		messages = append(messages, <-source)
	}

	return messages
}

func analyze(currentVal float64, messages []Msg) float64 {
	sum := currentVal
	for _, msg := range messages {
		sum += msg.Value
	}
	return sum / float64(len(messages)+1)
}

func Flooding(votes []float64) int {
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

	n := len(votes)
	maxRounds := n * 10

	var wg sync.WaitGroup
	results := make([]int, n)

	for i := range n {
		wg.Add(1)
		go node(votes[i], maxRounds, ins[i], outs[i], &results[i], &wg)
	}
	wg.Wait()

	fmt.Println("Deciziile nodurilor:", results)
	return results[0]
}

func main() {
	votes := []float64{1, 0, 1, 0, 1, 0, 1}

	fmt.Println("Vot majoritar:", Flooding(votes))
}
