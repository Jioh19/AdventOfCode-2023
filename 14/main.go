package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	grid := insertData(file)
	s := time.Now()
	tilted := tilt(grid)
	// for i := 0; i < len(tilted); i++ {
	// 	fmt.Printf("\n")
	// 	for j := 0; j < len(tilted); j++ {
	// 		fmt.Printf("%s", string(tilted[i][j]))

	// 	}
	// }
	part1 := count(tilted)
	fmt.Println("part 1:", part1)
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())
	//fmt.Println(grid)
}

func insertData(file []byte) []string {
	stringify := strings.TrimSpace(string(file))
	lines := strings.Split(stringify, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	return lines
}

func tilt(grid []string) [][]byte {
	var tilted [][]byte
	for i := 0; i < len(grid); i++ {
		tilted = append(tilted, make([]byte, len(grid[i])))
	}
	for k := 0; k < len(grid[0]); k++ {
		limit := 0
		for i := 0; i < len(grid); i++ {
			if grid[i][k] == 'O' {
				if limit == i {
					tilted[i][k] = 'O'
				} else {
					tilted[limit][k] = 'O'
					tilted[i][k] = '.'
				}
				limit++
			} else if grid[i][k] == '.' {
				tilted[i][k] = '.'
			} else if grid[i][k] == '#' {
				if i < len(grid)-1 {
					limit = i + 1
				}
				tilted[i][k] = '#'
			}

		}
	}
	return tilted
}

func count(tilted [][]byte) int {
	result := 0
	for i := 0; i < len(tilted); i++ {
		for j := 0; j < len(tilted); j++ {
			if tilted[i][j] == 'O' {
				result += len(tilted) - i
			}
		}
	}
	return result
}
