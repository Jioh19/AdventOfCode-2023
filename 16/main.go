package main

import (
	"fmt"
	"os"
	"time"
)

type Dir struct {
	i int
	j int
}

type PointDir struct {
	i  int
	j  int
	di int
	dj int
}

func main() {
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	grid := insertData(file)
	// for i, line := range grid {
	// 	fmt.Println(i, string(line))
	// }
	s := time.Now()
	part1 := count(part1(grid))
	fmt.Println(part1)
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())
	s2 := time.Now()
	part2 := part2(grid)
	fmt.Println(part2)
	fmt.Println("Time in nanoseconds:", time.Since(s2).Nanoseconds())
}

func insertData(file []byte) [][]byte {
	var input [][]byte
	var word []byte
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

func part1(grid [][]byte) [][]byte {
	energized := make([][]byte, len(grid))
	for i := range energized {
		energized[i] = make([]byte, len(grid[i]))
		for j := range energized[i] {
			energized[i][j] = '.'
		}
	}
	m := make(map[PointDir]bool)

	energize(grid, energized, 0, -1, Dir{0, 1}, m)
	for i, line := range energized {
		fmt.Println(i, string(line))
	}

	return energized
}

func part2(grid [][]byte) int {
	energized := make([][]byte, len(grid))
	for i := range energized {
		energized[i] = make([]byte, len(grid[i]))
		for j := range energized[i] {
			energized[i][j] = '.'
		}
	}
	m := make(map[PointDir]bool)
	max := 0
	for i := range grid {
		energize(grid, energized, -1, i, Dir{1, 0}, m)
		aux := count(energized)
		for i := range energized {
			for j := range energized[i] {
				energized[i][j] = '.'
			}
		}
		m = make(map[PointDir]bool)
		if max < aux {
			max = aux
		}
	}
	for i := range grid[0] {
		energize(grid, energized, i, -1, Dir{0, 1}, m)
		aux := count(energized)
		for i := range energized {
			for j := range energized[i] {
				energized[i][j] = '.'
			}
		}
		m = make(map[PointDir]bool)
		if max < aux {
			max = aux
		}
	}
	for i := range grid {
		energize(grid, energized, len(grid), i, Dir{-1, 0}, m)
		aux := count(energized)
		for i := range energized {
			for j := range energized[i] {
				energized[i][j] = '.'
			}
		}
		m = make(map[PointDir]bool)
		if max < aux {
			max = aux
		}
	}
	for i := range grid[0] {
		energize(grid, energized, i, len(grid[0]), Dir{0, -1}, m)
		aux := count(energized)
		for i := range energized {
			for j := range energized[i] {
				energized[i][j] = '.'
			}
		}
		m = make(map[PointDir]bool)
		if max < aux {
			max = aux
		}
	}

	return max
}

func energize(grid [][]byte, energized [][]byte, i, j int, dir Dir, m map[PointDir]bool) {
	i += dir.i
	j += dir.j
	if ok := m[PointDir{i, j, dir.i, dir.j}]; ok {
		return
	}
	if i < 0 || i >= len(grid[0]) {
		return
	}
	if j < 0 || j >= len(grid) {
		return
	}
	m[PointDir{i, j, dir.i, dir.j}] = true
	if energized[i][j] == '.' {
		energized[i][j] = '#'
	}

	switch {
	case grid[i][j] == '|':
		if dir.j != 0 {
			energize(grid, energized, i, j, Dir{-1, 0}, m)
			energize(grid, energized, i, j, Dir{1, 0}, m)
		} else {
			energize(grid, energized, i, j, dir, m)
		}
	case grid[i][j] == '-':
		if dir.i != 0 {
			energize(grid, energized, i, j, Dir{0, -1}, m)
			energize(grid, energized, i, j, Dir{0, 1}, m)
		} else {
			energize(grid, energized, i, j, dir, m)
		}
	case grid[i][j] == '\\':
		switch {
		case dir.j == -1:
			energize(grid, energized, i, j, Dir{-1, 0}, m)
		case dir.j == 1:
			energize(grid, energized, i, j, Dir{1, 0}, m)
		case dir.i == -1:
			energize(grid, energized, i, j, Dir{0, -1}, m)
		case dir.i == 1:
			energize(grid, energized, i, j, Dir{0, 1}, m)
		}
	case grid[i][j] == '/':
		switch {
		case dir.j == -1:
			energize(grid, energized, i, j, Dir{1, 0}, m)
		case dir.j == 1:
			energize(grid, energized, i, j, Dir{-1, 0}, m)
		case dir.i == -1:
			energize(grid, energized, i, j, Dir{0, 1}, m)
		case dir.i == 1:
			energize(grid, energized, i, j, Dir{0, -1}, m)
		}
	default:
		energize(grid, energized, i, j, dir, m)
	}
}

func count(energize [][]byte) int {
	result := 0
	for _, row := range energize {
		for _, col := range row {
			if col == '#' {
				result++
			}
		}
	}
	return result
}
