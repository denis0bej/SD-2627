package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X float64
	Y float64
}

const (
	viewDistance  = 20.0
	roundingError = 0.000001
	maxRobots     = 50
	maxIterations = 1000
)

func robot(pos Point, positions []Point, out chan<- Point) {
	neighbors := scan(pos, positions)
	var sum Point

	for _, neighbor := range neighbors {
		sum.X += neighbor.X
		sum.Y += neighbor.Y
	}

	count := float64(len(neighbors))
	out <- Point{X: sum.X / count, Y: sum.Y / count}
}

func dist(p1 Point, p2 Point) float64 {
	x := p1.X - p2.X
	y := p1.Y - p2.Y
	return math.Sqrt(x*x + y*y)
}

func scan(origin Point, positions []Point) []Point {
	neighbors := make([]Point, 0, len(positions))

	for _, position := range positions {
		if dist(origin, position) <= viewDistance {
			neighbors = append(neighbors, position)
		}
	}

	return neighbors
}

func recalibrate(points *[]Point) {
	set := make([]Point, 0, len(*points))

	for _, point := range *points {
		duplicate := false
		for _, existing := range set {
			if dist(point, existing) <= roundingError {
				duplicate = true
				break
			}
		}
		if !duplicate {
			set = append(set, point)
		}
	}

	*points = set
}

func main() {
	file, err := os.Open("lab5/ex1/input.txt")
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		fmt.Println("File read error:", scanner.Err())
		return
	}
	n, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Println("Not convertable to int:", err)
		return
	}

	positions := make([]Point, n)
	for i := range positions {
		if !scanner.Scan() {
			fmt.Printf("Invalid pos robot %d: %v\n", i+1, scanner.Err())
			return
		}
		words := strings.Fields(scanner.Text())

		positions[i].X, err = strconv.ParseFloat(words[0], 64)
		if err != nil {
			fmt.Printf("Invalid X robot %d: %v\n", i+1, err)
			return
		}
		positions[i].Y, err = strconv.ParseFloat(words[1], 64)
		if err != nil {
			fmt.Printf("Invalid Y robot %d: %v\n", i+1, err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("File read error:", err)
		return
	}

	for iteration := 0; iteration < maxIterations; iteration++ {
		numRobots := len(positions)
		if numRobots == 1 {
			break
		}

		resultsCh := make(chan Point, numRobots)

		for _, position := range positions {
			go robot(position, positions, resultsCh)
		}

		newPositions := make([]Point, numRobots)
		for i := 0; i < numRobots; i++ {
			newPositions[i] = <-resultsCh
		}

		positions = newPositions
		recalibrate(&positions)
	}

	if len(positions) > 0 {
		fmt.Println("Pozitia finala:", positions[0])
	}
}
