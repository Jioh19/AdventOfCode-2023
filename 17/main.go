package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type Dir struct {
	i int
	j int
}

func main() {
	fileName := "test.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	grid := insertData(file)
	for i, line := range grid {
		fmt.Println(i, string(line))
	}
	s := time.Now()
	part1 := part1(grid)
	fmt.Println(part1)
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())
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

func part1(grid [][]byte) int {

	m := make(map[string]int)
	res1, ok1 := travel(grid, 0, 0, Dir{1, 0}, 3, m, 0, 1)
	res2, ok2 := travel(grid, 0, 0, Dir{0, 1}, 3, m, 0, 1)
	fmt.Println(ok1, ok2)
	fmt.Println(m)
	if res1 < res2 {
		return res1
	}
	return res2
}

func travel(grid [][]byte, i, j int, dir Dir, f int, m map[string]int, di, dj int) (int, bool) {
	i = dir.i
	j = dir.j

	key := fmt.Sprintf("%d,%d,%d,%d", i, j, dir.i, dir.j)
	if value, ok := m[key]; ok {
		return value, ok
	}
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return 0, false
	}

	val, _ := strconv.Atoi(string(grid[i][j]))
	if i == di && j == dj {
		return val, true
	}

	result := math.MaxInt
	var ok bool
	if f > 0 {
		min, ok := travel(grid, i, j, dir, f-1, m, di, dj)
		if ok && result > min {
			result = min
		}
	}
	if dir.i == 0 {
		min, ok := travel(grid, i, j, Dir{1, 0}, 3, m, di, dj)
		if ok && result > min {
			result = min
		}
		min, ok = travel(grid, i, j, Dir{-1, 0}, 3, m, di, dj)
		if ok && result > min {
			result = min
		}
	}
	if dir.j == 0 {
		min, ok := travel(grid, i, j, Dir{0, 1}, 3, m, di, dj)
		if ok && result > min {
			result = min
		}
		min, ok = travel(grid, i, j, Dir{0, -1}, 3, m, di, dj)
		if ok && result > min {
			result = min
		}

	}
	result += val
	m[key] = result
	return result, ok
}
