package main

import (
	"fmt"
	"log"
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
	// fmt.Println(dir[i], val[i], color[i])
	fmt.Println(coords)
	part1 := shoeLaceFormula(coords, outline)
	fmt.Println("Part1", part1)
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())
	s2 := time.Now()
	coords2, outline2 := getCoord2(color)
	part2 := shoeLaceFormula(coords2, float64(outline2))
	fmt.Println(outline2)
	fmt.Println("Part2", part2)
	fmt.Println("Time in nanoseconds:", time.Since(s2).Nanoseconds())
}

func hexDecode(hex string) int64 {

	num, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return num

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

func getCoord2(color []string) ([]Coord, int64) {
	var coords []Coord
	var outline int64 = 0
	coord := Coord{0, 0}
	for _, d := range color {
		hexNum := d[2:7]
		num := float64(hexDecode(hexNum))
		dir := d[7:8]
		fmt.Println(d, hexNum, dir, num)
		switch dir {
		case "0":
			coord = Coord{coord.i, coord.j + num}
			outline += int64(num)
		case "2":
			coord = Coord{coord.i, coord.j - num}
			outline += int64(num)
		case "1":
			coord = Coord{coord.i + num, coord.j}
			outline += int64(num)
		case "3":
			coord = Coord{coord.i - num, coord.j}
			outline += int64(num)
		}
		coords = append(coords, coord)
	}
	return coords, outline
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

func shoeLaceFormula(coords []Coord, outline float64) int64 {
	// https://en.wikipedia.org/wiki/Shoelace_formula
	var polygonArea int64 = 0
	for i := 0; i < len(coords); i++ {
		cur := coords[i]
		next := coords[(i+1)%len(coords)]
		//Al fin me sirvió algebra lineal, producto cruz....
		polygonArea += int64(cur.j*next.i - cur.i*next.j)
	}

	if polygonArea < 0 {
		polygonArea = -polygonArea
	}
	polygonArea /= 2

	return polygonArea + int64(outline/2) + 1
}
