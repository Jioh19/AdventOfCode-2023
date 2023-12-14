package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Charts struct {
	grids [][]string
}

func main() {
	fileName := "test.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	//var charts Charts
	charts := insertData(file)
	s := time.Now()
	part1 := 0
	for _, lines := range charts.grids {
		part1 += getHorizontal(lines) * 100
		part1 += getVerticalSmudge(lines)
	}
	fmt.Println("Part1:", part1)
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())
}

func insertData(file []byte) Charts {
	var charts = new(Charts)
	data := strings.Split(string(file), "\n\r\n")
	for _, datum := range data {
		var grids [][]string
		lines := (strings.Split(strings.TrimSpace(datum), "\n"))
		for i, line := range lines {
			lines[i] = strings.TrimSpace(line)
		}
		grids = append(grids, lines)
		charts.grids = append(charts.grids, grids...)
	}
	return *charts
}

func getHorizontal(grid []string) int {
	var bet []int
	for i := 0; i < len(grid)-1; i++ {
		if grid[i] == grid[i+1] {
			bet = append(bet, i)
		}
	}
	total := 0
	//fmt.Println(bet)
	if len(bet) > 0 {
	DIFFERENT:
		for _, idx := range bet {
			for i, j := idx, idx+1; i >= 0 && j < len(grid); i, j = i-1, j+1 {
				if grid[i] != grid[j] {
					continue DIFFERENT
				}
			}
			total += idx + 1
		}
	}
	return total
}

func getHorizontalSmudge(grid []string) int {
	var bet []int
	for i := 0; i < len(grid)-1; i++ {
		var sumXor byte
		for j := 0; j < len(grid[i]); j++ {
			sumXor ^= grid[i][j] ^ grid[i+1][j]
			fmt.Println(sumXor)
		}
		fmt.Println(sumXor, i)

		// if grid[i] == grid[i+1] {
		// 	bet = append(bet, i)
		// }
	}
	total := 0
	//fmt.Println(bet)
	if len(bet) > 0 {
	DIFFERENT:
		for _, idx := range bet {
			for i, j := idx, idx+1; i >= 0 && j < len(grid); i, j = i-1, j+1 {
				if grid[i] != grid[j] {
					continue DIFFERENT
				}
			}
			total += idx + 1
		}
	}
	return total
}

func getVertical(grid []string) int {
	var bet []int
	for i := 0; i < len(grid[0])-1; i++ {
		counter := 0
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == grid[j][i+1] {
				//fmt.Println(string(grid[j][i]), j, i)
				counter++
			}
			if counter == len(grid)-1 {
				bet = append(bet, i)
			}
		}
	}

	total := 0
	//fmt.Println(bet)
	if len(bet) > 0 {
	DIFFERENT:
		for _, idx := range bet {
			for i, j := idx, idx+1; i >= 0 && j < len(grid[0]); i, j = i-1, j+1 {
				for k := 0; k < len(grid); k++ {
					if grid[k][i] != grid[k][j] {
						continue DIFFERENT
					}
				}
			}
			total += idx + 1
			//fmt.Println(total)
		}
	}
	return total
}

func getVerticalSmudge(grid []string) int {
	var bet []int
	var smudge []bool

	for i := 0; i < len(grid[0])-1; i++ {
		smudge = append(smudge, false)
		counter := 0
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == grid[j][i+1] {
				//fmt.Println(string(grid[j][i]), j, i)
				counter++
			} else if smudge[i] {
				smudge[i] = true
				counter++
			}
			if counter == len(grid)-1 {
				bet = append(bet, i)
			}
		}
	}

	total := 0
	//fmt.Println(bet)
	if len(bet) > 0 {
	DIFFERENT:
		for _, idx := range bet {
			for i, j := idx, idx+1; i >= 0 && j < len(grid[0]); i, j = i-1, j+1 {
				for k := 0; k < len(grid); k++ {
					if grid[k][i] != grid[k][j] {
						if smudge[i] {
							continue DIFFERENT
						}
						smudge[i] = true
					}
				}
			}
			total += idx + 1
			//fmt.Println(total)
		}
	}
	return total
}
