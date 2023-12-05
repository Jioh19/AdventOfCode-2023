package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}
type gridNumBuilder struct {
	sb  strings.Builder
	pos coord
}

func (gnb *gridNumBuilder) setPos(c coord) {
	gnb.pos = c
}

func (gnb *gridNumBuilder) addDigit(c rune) {
	gnb.sb.WriteRune(c)
}

func (gnb gridNumBuilder) empty() bool {
	return gnb.sb.Len() == 0
}

func (gnb *gridNumBuilder) flush() (coord, gridNum) {
	pos := gnb.pos
	gn := parseNum(gnb.sb.String())
	gnb.reset()
	return pos, gn
}

func (gnb *gridNumBuilder) reset() {
	gnb.pos = coord{}
	gnb.sb.Reset()
}

type gridNum int

func (g gridNum) bounds(pos coord) (coord, coord) {
	l := len(strconv.Itoa(int(g)))
	return coord{x: pos.x - 1, y: pos.y - 1}, coord{x: pos.x + l, y: pos.y + 1}
}

type numGrid map[coord]gridNum
type symbolGrid map[coord]rune

func contains(p coord, min coord, max coord) bool {
	return p.x >= min.x && p.y >= min.y && p.x <= max.x && p.y <= max.y
}

func parseNum(s string) gridNum {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return gridNum(n)
}

func parseGrid(lines []string) (numGrid, symbolGrid) {
	nums := numGrid{}
	symbols := symbolGrid{}
	for y, line := range lines {
		var gnb gridNumBuilder
		for x, ch := range line {
			c := coord{x: x, y: y}
			switch ch {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
				// handle digit
				if gnb.empty() {
					gnb.setPos(c)
				}
				gnb.addDigit(ch)
				// process more characters as long as we're not at the EOL
				if x < len(line)-1 {
					continue
				}
			case '.':
				// empty, do nothing
			default:
				// handle symbol
				symbols[c] = ch
			}

			if !gnb.empty() {
				pos, gn := gnb.flush()
				nums[pos] = gn
			}
		}
	}
	return nums, symbols
}

func findParts(nums numGrid, symbols symbolGrid) (parts []int) {
	for npos, num := range nums {
		for spos := range symbols {
			min, max := num.bounds(npos)
			if contains(spos, min, max) {
				parts = append(parts, int(num))
			}
		}
	}
	return parts
}

func part1(lines []string) {
	result := 0
	nums, symbols := parseGrid(lines)
	for _, n := range findParts(nums, symbols) {
		result += n
	}
	fmt.Printf("Part 1: %d\n", result)
}

func part2(lines []string) {
	result := 0
	nums, symbols := parseGrid(lines)
	for spos := range symbols {
		adjacent := []int{}
		for npos, gn := range nums {
			min, max := gn.bounds(npos)
			if contains(spos, min, max) {
				adjacent = append(adjacent, int(gn))
			}
		}
		if len(adjacent) == 2 {
			result += adjacent[0] * adjacent[1]
		}
	}
	fmt.Printf("part 2: %d\n", result)
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	part1(lines)
	part2(lines)
}
