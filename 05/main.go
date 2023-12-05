package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Scope struct {
	source int
	dest   int
	span   int
}

type Chart struct {
	scopes []Scope
}

func main() {

	var seeds []int

	fileName := "input.txt"

	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(file), "\n")
	seeds = insertSeeds(lines)
	soil, index := insertChart(lines, 2)
	fertilizer, index := insertChart(lines, index)
	water, index := insertChart(lines, index)
	light, index := insertChart(lines, index)
	temp, index := insertChart(lines, index)
	humidity, index := insertChart(lines, index)
	location, index := insertChart(lines, index)
	min := math.MaxInt
	minSeed := 0
	for _, seed := range seeds {
		s := charting(soil, seed)
		s = charting(fertilizer, s)
		s = charting(water, s)
		s = charting(light, s)
		s = charting(temp, s)
		s = charting(humidity, s)
		s = charting(location, s)
		if min > s {
			minSeed = seed
			min = s
		}
	}
	fmt.Println(min, minSeed)
	_, _, _, _, _, _, _, _ = seeds, soil, fertilizer, water, light, temp, humidity, location
}

func insertSeeds(lines []string) []int {
	var result []int
	line := strings.Split(lines[0], ":")[1]
	numbers := strings.Split(strings.TrimSpace(line), " ")
	for i, number := range numbers {
		numbers[i] = strings.TrimSpace(number)
		val, _ := strconv.Atoi(numbers[i])
		result = append(result, val)
	}
	return result
}

func insertChart(lines []string, index int) (Chart, int) {
	chart := new(Chart)
	//fmt.Println(lines[index])
	index++
	for {
		numbers := strings.Split(strings.TrimSpace(lines[index]), " ")
		scope := new(Scope)
		var arr []int
		for i, number := range numbers {
			numbers[i] = strings.TrimSpace(number)
			val, _ := strconv.Atoi(numbers[i])
			arr = append(arr, val)
		}
		//fmt.Println(arr)
		scope.dest = arr[0]
		scope.source = arr[1]
		scope.span = arr[2]
		chart.scopes = append(chart.scopes, *scope)
		index++
		lines[index] = strings.TrimSpace(lines[index])
		if lines[index] == "" {
			break
		}
	}
	index++
	return *chart, index
}

func charting(chart Chart, seed int) int {
	var result int = seed
	for _, scope := range chart.scopes {
		//	fmt.Println("Charting: ", seed)
		if seed >= scope.source && seed <= scope.source+scope.span {
			//fmt.Println(seed, scope.source, scope.span, scope.dest)
			result = scope.dest + seed - scope.source
		}
	}
	//fmt.Println(result)
	return result
}
