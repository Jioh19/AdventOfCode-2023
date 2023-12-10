package main

import (
	"fmt"
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
	grid := insertData(file)
	start := findStart(grid)
	_ = start
	findNext(grid, Point{0, 3})
	s := time.Now()
	result1, route := navigate(grid, start)
	fmt.Println("Part 1:", result1)
	//fmt.Println(route)
	result2 := shoeLaceFormula(grid, start, route)
	fmt.Println("Part 2:", result2)
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())
}

func shoeLaceFormula(grid []string, start Point, route []Point) any {
	// https://en.wikipedia.org/wiki/Shoelace_formula
	polygonArea := 0
	for i := 0; i < len(route); i++ {
		cur := route[i]
		next := route[(i+1)%len(route)]
		//Al fin me sirvió algebra lineal, producto cruz....
		polygonArea += cur.x*next.y - cur.y*next.x
	}

	if polygonArea < 0 {
		polygonArea = -polygonArea
	}
	polygonArea /= 2

	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	return polygonArea - len(route)/2 + 1
}

func navigate(grid []string, start Point) (int, []Point) {
	visited := make(map[Point]int)
	current := start
	next := findNext(grid, start)
	next = next[0:1]
	result := 0
	var route []Point
	route = append(route, current)
	for len(next) > 0 {
		visited[current] = 1
		if visited[next[0]] == 0 {
			result++
			current = next[0]
			route = append(route, current)
			visited[next[0]] = 1
			next = append(next, findNext(grid, current)...)
		}
		next = next[1:]
	}
	result = (result + 1) / 2
	return result, route
}

func findStart(grid []string) Point {
	for y, row := range grid {
		for x, col := range row {
			if col == 'S' {
				return Point{x, y}
			}
		}
	}
	return Point{}
}

func findNext(grid []string, pos Point) []Point {
	current := grid[pos.y][pos.x]
	var points []Point
	switch current {
	case 'S':
		if pos.y+1 < len(grid) {
			next := grid[pos.y+1][pos.x]
			if next == '|' || next == 'L' || next == 'J' {
				points = append(points, Point{pos.x, pos.y + 1})
			}
		}
		if pos.y-1 >= 0 {
			next := grid[pos.y-1][pos.x]
			if next == '|' || next == 'F' || next == '7' {
				points = append(points, Point{pos.x, pos.y - 1})
			}
		}
		if pos.x+1 < len(grid[pos.y]) {
			next := grid[pos.y][pos.x+1]
			if next == '-' || next == 'J' || next == '7' {
				points = append(points, Point{pos.x + 1, pos.y})
			}
		}
		if pos.x-1 >= 0 {
			next := grid[pos.y][pos.x-1]
			if next == '-' || next == 'L' || next == 'F' {
				points = append(points, Point{pos.x - 1, pos.y})
			}
		}
	case 'F':
		points = append(points, Point{pos.x, pos.y + 1})
		points = append(points, Point{pos.x + 1, pos.y})
	case '7':
		points = append(points, Point{pos.x, pos.y + 1})
		points = append(points, Point{pos.x - 1, pos.y})
	case '-':
		points = append(points, Point{pos.x + 1, pos.y})
		points = append(points, Point{pos.x - 1, pos.y})
	case 'J':
		points = append(points, Point{pos.x, pos.y - 1})
		points = append(points, Point{pos.x - 1, pos.y})
	case 'L':
		points = append(points, Point{pos.x, pos.y - 1})
		points = append(points, Point{pos.x + 1, pos.y})
	case '|':
		points = append(points, Point{pos.x, pos.y + 1})
		points = append(points, Point{pos.x, pos.y - 1})
	}
	return points
}

func insertData(file []byte) []string {
	var grid []string
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	for _, line := range lines {
		pipe := strings.TrimSpace(string(line))
		//	fmt.Println(len(pipe), pipe)
		grid = append(grid, pipe)
	}
	return grid
}
