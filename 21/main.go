package main

import (
	"fmt"
	"os"
	"time"
)

type Coord struct {
	i int
	j int
}

func main() {
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	input := insertData(file)
	start := getStart(input)
	s := time.Now()
	part1(input, 25, start.i, start.j, 1, 1)
	for i, row := range input {
		fmt.Printf("%d", i)
		for _, col := range row {
			fmt.Printf("%s", string(col))

		}
		fmt.Println()
	}

	part1 := count(input)
	fmt.Println("Part 1:", part1)
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())
}

func insertData(file []byte) [][]byte {
	input := [][]byte{}
	word := []byte{}
	for _, letter := range file {
		if letter != 13 {
			if letter == '\n' {
				input = append(input, word)
				word = []byte{}
			} else {
				word = append(word, letter)
			}
		}
	}
	return input
}

func getStart(input [][]byte) Coord {
	start := Coord{}
	for i, row := range input {
		for j, letter := range row {
			if letter == 'S' {
				start.i = i
				start.j = j
			}
		}
	}
	return start
}

func part1(input [][]byte, steps, i, j, mi, mj int) {
	if i < 0 || i >= len(input) {
		return
	}
	if j < 0 || j >= len(input[i]) {
		return
	}
	if input[i][j] == '#' {
		return
	}
	if input[i][j] != '#' {
		input[i][j] = 'O'
	}

	if steps <= 0 {
		return
	}
	steps--
	fmt.Println(steps)
	part1(input, steps, i, j+mj*2, mi, mj)
	part1(input, steps, i+mi*2, j, mi, mj)
	part1(input, steps, i+mi, j+mj, mi, mj)
}

func recurse(input [][]byte, steps, i, j int) {
	if i < 0 || i >= len(input) {
		return
	}
	if j < 0 || j >= len(input[i]) {
		return
	}
	if input[i][j] == '#' {
		return
	}
	if steps <= 0 {
		if input[i][j] != '#' {
			input[i][j] = 'O'
			return
		} else {
			return
		}
	}
	steps--
	recurse(input, steps, i-1, j)
	recurse(input, steps, i+1, j)
	recurse(input, steps, i, j-1)
	recurse(input, steps, i, j+1)

	return
}

func count(input [][]byte) int {
	result := 0
	for _, row := range input {
		for _, col := range row {
			if col == 'O' {
				result++
			}
		}
	}
	return result
}
