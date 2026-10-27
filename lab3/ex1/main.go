package main

import (
	"fmt"
	"math/rand"
	"time"
)

// stocheaza deviatia fata de timpul real
type Clock struct{ offset time.Duration }

func (c *Clock) Now() time.Time { return time.Now().Add(c.offset) }

// TODO:
// Calculeaza decalajul theta si round -trip -ul delta
func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	return ((t2.Sub(t1)) + (t3.Sub(t4))) / 2, t4.Sub(t1) - t3.Sub(t2)
}
func measure(c_clock, s_clock *Clock) (theta, delta time.Duration) {
	// Request
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.
		Millisecond)
	t2 := s_clock.Now()
	// Procesare intre 5 si 20 ms
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.
		Millisecond)
	t3 := s_clock.Now()
	// Response
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.
		Millisecond)
	t4 := c_clock.Now()
	return estimate(t1, t2, t3, t4)
}

func main() {

	c := &Clock{offset: 0}
	s := &Clock{offset: 1 * time.Second}
	fmt.Print("(a) ")
	t, d := measure(c, s)
	fmt.Println(t.Milliseconds(), d.Milliseconds())

	fmt.Println("(b)")

	var tmin, dmin time.Duration

	for i := range 10 {
		fmt.Printf("(%v) ", i+1)
		t, d := measure(c, s)
		if i == 0 {
			tmin, dmin = t, d
		}
		if t < tmin {
			tmin = t
		}
		if d < dmin {
			dmin = d
		}
		fmt.Println(t.Milliseconds(), d.Milliseconds())
	}

	fmt.Printf("(min) %v %v", tmin.Milliseconds(), dmin.Milliseconds())

}
