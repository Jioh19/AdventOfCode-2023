package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

type Point struct {
	x int
	y int
}

func main() {
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := time.Now()
	grid := insertData(file)
	// for _, row := range grid {
	// 	fmt.Println(row, len(grid))
	// }
	grid = expandY(grid)
	grid = expandX(grid)
	galaxies := findGalaxies(grid)
	part1 := getAllDiscantes(galaxies)
	// for _, row := range grid {
	// 	fmt.Println(row, len(grid))
	// }
	// fmt.Println(galaxies)
	fmt.Println("Part 1:", part1)
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())
}

func insertData(file []byte) []string {
	var grid []string
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(string(line))
		grid = append(grid, line)
	}
	return grid
}

func expandY(grid []string) []string {
	empty := strings.Repeat(".", len(grid[0]))
GALAXY:
	for y := len(grid) - 1; y >= 0; y-- {
		for _, col := range grid[y] {
			if col == '#' {
				continue GALAXY
			}
		}
		aux := grid[:y]
		aux = append(aux, empty)
		grid = append(aux, grid[y:]...)
	}
	return grid
}

func expandX(grid []string) []string {
GALAXY:
	for x := len(grid[0]) - 1; x >= 0; x-- {
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == '#' {
				continue GALAXY
			}
		}
		for y := 0; y < len(grid); y++ {
			aux := ""
			aux += grid[y][:x] + "." + grid[y][x:]
			grid[y] = aux
		}
	}
	return grid
}

func findGalaxies(grid []string) []Point {
	var galaxies []Point
	for x, row := range grid {
		for y, col := range row {
			if col == '#' {
				galaxies = append(galaxies, Point{x, y})
			}
		}
	}
	//fmt.Println(galaxies)
	return galaxies
}

func getDistance(a Point, b Point) int {
	dx := math.Abs(float64(a.x - b.x))
	dy := math.Abs(float64(a.y - b.y))
	distance := int(dx + dy)
	return distance
}

func getAllDiscantes(galaxies []Point) int {
	total := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			total += getDistance(galaxies[i], galaxies[j])

		}
	}
	return total
}
