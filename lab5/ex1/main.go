package main

import (
	"fmt"
)

type Point struct {
	X float64
	Y float64
}

func main(){
	var n int
    fmt.Scan(&n)

    pos := make([]Point, n)
    for i := range pos {
        fmt.Scan(&pos[i].X, &pos[i].Y)
    }


}