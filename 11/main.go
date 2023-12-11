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
	galaxies2 := findGalaxies(grid)
	grid, expY := expandY(grid)
	grid, expX := expandX(grid)
	galaxies := findGalaxies(grid)
	part1 := getAllDiscantes(galaxies)
	// for _, row := range grid {
	// 	fmt.Println(row, len(grid))
	// }
	// fmt.Println(galaxies)
	fmt.Println("Part 1:", part1)
	fmt.Println(expY, expX)
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())
	s2 := time.Now()
	part2 := getAllDiscantes2(galaxies2, expY, expX)
	fmt.Println("Part 2:", part2)
	fmt.Println("Time in nanoseconds:", time.Since(s2).Nanoseconds())
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

func expandY(grid []string) ([]string, []int) {
	empty := strings.Repeat(".", len(grid[0]))
	var expY []int
GALAXY:
	for y := len(grid) - 1; y >= 0; y-- {
		for _, col := range grid[y] {
			if col == '#' {
				continue GALAXY
			}
		}
		expY = append(expY, y)
		aux := grid[:y]
		aux = append(aux, empty)
		grid = append(aux, grid[y:]...)
	}
	return grid, expY
}

func expandX(grid []string) ([]string, []int) {
	var expX []int
GALAXY:
	for x := len(grid[0]) - 1; x >= 0; x-- {
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == '#' {
				continue GALAXY
			}
		}
		expX = append(expX, x)
		for y := 0; y < len(grid); y++ {
			aux := ""
			aux += grid[y][:x] + "." + grid[y][x:]
			grid[y] = aux
		}
	}
	return grid, expX
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
	fmt.Println(galaxies)
	return galaxies
}

func getDistance(a Point, b Point) int {
	dx := math.Abs(float64(a.x - b.x))
	dy := math.Abs(float64(a.y - b.y))
	distance := int(dx + dy)
	return distance
}

func getDistance2(a Point, b Point, expX []int, expY []int) int {
	if a.x < b.x {
		a.x, b.x = b.x, a.x
	}
	dx := a.x - b.x
	for _, lenX := range expX {
		if b.x <= lenX && a.x >= lenX {
			dx += 999999
			//fmt.Println(b.x, lenX, a.x)
		}
	}
	if a.y < b.y {
		a.y, b.y = b.y, a.y
	}
	dy := a.y - b.y
	for _, lenY := range expY {
		if b.y <= lenY && a.y >= lenY {
			dy += 999999
			//fmt.Println(b.y, lenY, a.y)
		}
	}
	distance := dx + dy
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

func getAllDiscantes2(galaxies []Point, expY []int, expX []int) int {
	total := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			total += getDistance2(galaxies[i], galaxies[j], expY, expX)

		}
	}
	return total
}
