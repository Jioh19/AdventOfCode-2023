package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Coord struct {
	i float64
	j float64
}

func main() {
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	s := time.Now()
	dir, val, color := insertData(file)
	_, _, _ = dir, val, color
	coords := getCoord(dir, val)
	outline := fillGrid(dir, val, color)
	for i := 0; i < len(coords); i++ {
		//fmt.Println(dir[i], val[i], color[i])
		fmt.Println(coords[i])
	}
	// fmt.Println(dir[i], val[i], color[i])
	// fmt.Println(coords[i])
	part1 := shoeLaceFormula(coords, outline)
	fmt.Println("Part1", part1)
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())
}

func insertData(file []byte) ([]string, []float64, []string) {
	var dir []string
	var val []float64
	var color []string
	split := strings.Split(strings.TrimSpace(string(file)), "\n")
	for _, line := range split {

		dir = append(dir, strings.Fields(line)[0])
		num, _ := strconv.Atoi(strings.Fields(line)[1])
		val = append(val, float64(num))
		color = append(color, strings.Fields(line)[2])
	}
	return dir, val, color
}
func makeGrid(dir []string, val []float64) [][]byte {
	maxi, maxj := 0, 0
	sumi, sumj := 1, 1
	for i, num := range val {
		switch dir[i] {
		case "R":
			sumj += int(num)
		case "L":
			sumj -= int(num)
		case "D":
			sumi += int(num)
		case "U":
			sumi -= int(num)
		}
		if sumi < 0 {
			fmt.Println("Sumi is negative", i)
		}
		if sumj < 0 {
			fmt.Println("Sumj is negative", i)
		}
		if maxi < sumi {
			maxi = sumi
		}
		if maxj < sumj {
			maxj = sumj
		}
	}
	fmt.Println(maxi, maxj)
	grid := make([][]byte, maxi)
	for i := range grid {
		grid[i] = make([]byte, maxj)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	return grid
}

func fillGrid(dir []string, val []float64, color []string) float64 {
	outline := 0.0
	posi, posj := 0, 0
	for i, num := range val {
		for j := 0; j < int(num); j++ {
			switch dir[i] {
			case "R":
				posj++
				outline++
			case "L":
				posj--
				outline++
			case "D":
				posi++
				outline++
			case "U":
				posi--
				outline++
			}
			//grid[posi][posj] = '#'
		}
	}
	return outline
}

func getCoord(dir []string, val []float64) []Coord {
	var coords []Coord
	coord := Coord{0, 0}
	for i, d := range dir {
		switch d {
		case "R":
			coord = Coord{coord.i, coord.j + val[i]}
		case "L":
			coord = Coord{coord.i, coord.j - val[i]}
		case "D":
			coord = Coord{coord.i + val[i], coord.j}
		case "U":
			coord = Coord{coord.i - val[i], coord.j}
		}
		coords = append(coords, coord)
	}
	return coords
}

func shoeLaceFormula(coords []Coord, outline float64) float64 {
	// https://en.wikipedia.org/wiki/Shoelace_formula
	polygonArea := 0.0
	for i := 0; i < len(coords); i++ {
		cur := coords[i]
		next := coords[(i+1)%len(coords)]
		//Al fin me sirvió algebra lineal, producto cruz....
		polygonArea += cur.j*next.i - cur.i*next.j
	}

	if polygonArea < 0 {
		polygonArea = -polygonArea
	}
	polygonArea /= 2

	return polygonArea + outline/2 + 1
}
