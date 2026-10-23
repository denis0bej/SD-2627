package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("ex7/input.txt")
	if err != nil {
		fmt.Println("Failed to open file! ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dict := map[string]int{}

	for scanner.Scan() {
		var line string = scanner.Text()
		words := strings.Split(line, " ")

		for _, word := range words {
			if word != "" {
				_, ok := dict[word]
				if !ok { // no dict entry
					dict[word] = 0
				}
				dict[word] = dict[word] + 1
			}
		}
	}
	fmt.Println("Word:\t\tCount:")
	for k, v := range dict {
		fmt.Printf("%v\t\t%v\n", k, v)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("File read error: ", err)
	}
}
