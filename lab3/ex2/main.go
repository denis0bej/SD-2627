package main

type LamportClock struct{ counter int }

// incrementeaza contorul
func (l *LamportClock) Tick() { l.counter += 1 }
func (l *LamportClock) Send() {
	l.Tick()
	// idk cum trimit unde
}
func (l *LamportClock) Receive(ts int) { l.counter = ts }

type Message struct{ sender, ts int }
type Event struct {
	process          int
	internal_counter int
	kind             string // "internal", "send", "recv"
}

// TODO: retuneaza true daca a precede b
func less(a, b Event) bool { return false }

func main() {
}
