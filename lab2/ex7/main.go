package main

import "fmt"

func robot(id int, traseu [7][9]int, done chan bool, sendch chan<- [2]int, receivech <-chan [2]int, resultch chan string) {
	x, y := 1, 1
	traseu[x][y] = 0
	sendch <- [2]int{x, y}
	for traseu[x][y] != 2 {

		select {
		case <-done:
			fmt.Printf("ROBOT[%v]: CITESC POZITIILE DE LA WINNER:\n", id)
			for {
				select {
				case pos, ok := <-receivech:
					if !ok {
						return
					}
					fmt.Printf(" Robot[%v] a citit %v\n", id, pos)
				default:
					// Gata
					resultch <- fmt.Sprintf("ROBOT[%v]: DONE! (am pierdut)", id)
					return
				}
			}
		default:
			pozitii := [4][2]int{
				{x - 1, y},
				{x, y - 1},
				{x + 1, y},
				{x, y + 1},
			}
			for _, poz := range pozitii { // verificare daca am ajuns la final (pe pozitia cu val 2)
				_x, _y := poz[0], poz[1]

				if traseu[_x][_y] == 2 {
					x = _x
					y = _y
					traseu[x][y] = 0
					sendch <- [2]int{x, y}
					done <- true
					fmt.Printf("ROBOT[%v]: A FOST CITIT MESAJUL DE DONE, TRIMIT RESULT LA MAIN()\n", id)
					resultch <- fmt.Sprintf("ROBOT[%v]: DONE! (am castigat)", id)
					return
				}
			}

			for _, poz := range pozitii { // plimbare
				_x, _y := poz[0], poz[1]

				if traseu[_x][_y] == 1 {
					x = _x
					y = _y
					traseu[x][y] = 0
					sendch <- [2]int{x, y}
					break
				}
			}
		}
	}
}

func main() {
	traseu1 := [7][9]int{ // nu am chef de edge cases deci bordura de 0 si coordonatele incep de la {1,1}
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 2, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	traseu2 := [7][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 2, 0},
		{0, 1, 0, 1, 1, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1, 0},
		{0, 1, 1, 1, 0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	donech := make(chan bool)
	robotch1 := make(chan [2]int, 40)
	robotch2 := make(chan [2]int, 40)
	responsech := make(chan string, 1)

	go robot(1, traseu1, donech, robotch1, robotch2, responsech)
	go robot(2, traseu2, donech, robotch2, robotch1, responsech)

	for range 2 {
		fmt.Println(<-responsech)
	}
}
