package main

import (
	"fmt"
	"os"
	"strings"
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
	//dir := make(map[string]Coord)
	grid := insertData(file)
	start := findStart(grid)
	_ = start
	findNext(grid, Point{0, 3})
	result1 := navigate(grid, start)
	fmt.Println(result1)
	// startNode := getStart(grid)
	// navigate(grid, startNode, *startNode.next)
	//fmt.Println(grid, start)
}

func navigate(grid []string, start Point) int {
	visited := make(map[Point]int)
	current := start
	next := findNext(grid, start)
	result := 0
	for len(next) > 0 {
		//fmt.Println(next)
		visited[current] = 1
		if visited[next[0]] == 0 {
			result++
			current = next[0]
			visited[next[0]] = 1
			//fmt.Println(string(grid[next[0].y][next[0].x]), result)
			next = append(next, findNext(grid, current)...)
		}
		next = next[1:]
	}
	result = (result + 1) / 2
	return result
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
	//	fmt.Println(string(current), points)
	return points
}

// func createDirMap(m map[string]Coord) {
// 	m["U"]
// }

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

// func getStart(maze Maze) Node {
// 	startNode := new(Node)
// 	coord := new(Coord)
// 	for i, pipe := range maze.pipes {
// 		j := strings.Index(pipe, "S")
// 		if j != -1 {
// 			coord.i, coord.j = i, j
// 			startNode.current = *coord
// 			nextNode := new(Node)
// 			prevNode := new(Node)
// 			nextNode.prev = startNode
// 			prevNode.next = startNode
// 			prev := new(Coord)
// 			next := new(Coord)
// 			if validEntry(maze, i+1, j, "U") {
// 				next.i = i + 1
// 				next.j = j
// 				nextNode.current = *next
// 				startNode.next = nextNode
// 			}
// 			if validEntry(maze, i-1, j, "D") {
// 				next.i = i - 1
// 				next.j = j
// 				nextNode.current = *next
// 				startNode.next = nextNode
// 			}
// 			if validEntry(maze, i, j+1, "R") {
// 				prev.i = i
// 				prev.j = j
// 				prevNode.current = *prev
// 				startNode.prev = prevNode
// 			}
// 			if validEntry(maze, i, j-1, "L") {
// 				prev.i = i
// 				prev.j = j
// 				prevNode.current = *prev
// 				startNode.prev = prevNode
// 			}
// 			return *startNode
// 		}
// 	}
// 	return *startNode
// }

// func navigate(maze Maze, startNode Node, currentNode Node) int {
// 	if currentNode.next == &startNode {
// 		return 1
// 	}
// 	prevNode := currentNode.prev
// 	fmt.Println(prevNode.current)

// next := new(Coord)
// if validEntry(maze, i+1, j, "U") {
// 	next.i = i + 1
// 	next.j = j
// 	nextNode.current = *next
// 	startNode.next = nextNode
// }
// if validEntry(maze, i-1, j, "D") {
// 	next.i = i - 1
// 	next.j = j
// 	nextNode.current = *next
// 	startNode.next = nextNode
// }
// if validEntry(maze, i, j+1, "R") {
// 	prev.i = i
// 	prev.j = j
// 	prevNode.current = *prev
// 	startNode.prev = prevNode
// }
// if validEntry(maze, i, j-1, "L") {
// 	prev.i = i
// 	prev.j = j
// 	prevNode.current = *prev
// 	startNode.prev = prevNode
// }

//navigate(maze, startNode, *startNode.next)
// if maze.pipes[i][j] == 'S' && maze.travel[i][j] == 1 {
// 	return 1
// }
// maze.travel[i][j] = 1
// //dir := [4]string{"U", "D", "L", "R"}

// if validEntry(maze, i+1, j, "U") {
// 	return navigate(maze, i+1, j) + 1
// }
// if validEntry(maze, i-1, j, "D") {
// 	return navigate(maze, i-1, j) + 1
// }
// if validEntry(maze, i, j+1, "R") {
// 	return navigate(maze, i, j+1) + 1
// }
// if validEntry(maze, i, j-1, "L") {
// 	return navigate(maze, i, j-1) + 1
// }
// fmt.Println(maze.travel)
// 	return 0
// }

// func validEntry(maze Maze, i int, j int, direction string) bool {
// 	if i < 0 || i >= len(maze.pipes) || j < 0 || j >= len(maze.pipes[0]) {
// 		return false
// 	}
// 	fmt.Println()
// 	switch direction {
// 	case "U":
// 		if maze.pipes[i][j] == '7' || maze.pipes[i][j] == 'F' || maze.pipes[i][j] == '|' {
// 			return true
// 		} else {
// 			return false
// 		}
// 	case "D":
// 		if maze.pipes[i][j] == 'L' || maze.pipes[i][j] == 'J' || maze.pipes[i][j] == '|' {
// 			return true
// 		} else {
// 			return false
// 		}
// 	case "R":
// 		if maze.pipes[i][j] == '7' || maze.pipes[i][j] == 'J' || maze.pipes[i][j] == '-' {
// 			return true
// 		} else {
// 			return false
// 		}
// 	case "L":
// 		if maze.pipes[i][j] == 'L' || maze.pipes[i][j] == 'F' || maze.pipes[i][j] == '-' {
// 			return true
// 		} else {
// 			return false
// 		}
// 	}
// 	return false
// }

// func validExit(destination string, direction string) bool {
// 	switch direction {
// 	case "U":
// 		if destination == "7" || destination == "F" || destination == "|" {
// 			return true
// 		} else {
// 			return false
// 		}
// 	case "D":
// 		if destination == "L" || destination == "J" || destination == "|" {
// 			return true
// 		} else {
// 			return false
// 		}
// 	case "R":
// 		if destination == "7" || destination == "J" || destination == "-" {
// 			return true
// 		} else {
// 			return false
// 		}
// 	case "L":
// 		if destination == "L" || destination == "F" || destination == "-" {
// 			return true
// 		} else {
// 			return false
// 		}
// 	}
// 	return false
// }
