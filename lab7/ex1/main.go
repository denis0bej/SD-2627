package main

import (
	"fmt"
	"math/rand/v2"
)

// Message reprezinta un mesaj transmis intre noduri in Runda 1 sau 2
type Message struct {
	SenderID int
	Round    int // 1 pentru starea curenta , 2 pentru valoarea Tie
	Value    int // 0 sau 1
}

type Node struct {
	id          int
	x           int // starea curenta x_i(t)
	s           int
	n           int  // numarul de noduri din sistem
	isByzantine bool // comportamentul la iteratia curenta
	inbox       chan Message
	targets     []*Node
}

func Send(msg Message, target Node) {
	target.inbox <- msg
}

func (n Node) Recv(msg Message) {
	// TODO
}

func (n Node) Broadcast(msg Message) {
	if n.isByzantine {
		msg.Value = (msg.Value + 1) % 2
	}
	for _, t := range n.targets {
		Send(msg, *t)
	}
}

func PrintNode(n Node) {
	fmt.Printf("id: %v\nx: %v\nisBysantine: %v\ninbox: %v\ntargets: ", n.id, n.x, n.isByzantine, n.inbox)
	for _, t := range n.targets {
		fmt.Print(t.id, " ")
	}
	fmt.Println("\n ")
}

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	fmt.Println(n, s)

	// INIT

	nodes := make([]Node, n)
	byzantines := make([]bool, n)

	{
		nr_byz := 0
		for nr_byz < s {
			victim := rand.IntN(n)

			if byzantines[victim] == true {
				continue
			}
			byzantines[victim] = true
			nr_byz += 1
		}
	}

	for i := range nodes {
		nodes[i].id, nodes[i].s, nodes[i].n, nodes[i].isByzantine, nodes[i].inbox = i, s, n, byzantines[i], make(chan Message)
		nodes[i].targets = make([]*Node, 0)
		for j := range nodes {
			if j == i {
				continue
			}
			nodes[i].targets = append(nodes[i].targets, &nodes[j])
		}
		fmt.Scan(&nodes[i].x)
	}

	// for i := range nodes {
	// 	PrintNode(nodes[i])
	// }

	// ALGORITM

	for t := range s + 2 {

		// RUNDA 1

		// Broadcast
		for i := range nodes {
			go nodes[i].Broadcast(Message{nodes[i].id, t, nodes[i].x})
		}
		// Listen
		// TODO

		// RUNDA 2

	}
}
