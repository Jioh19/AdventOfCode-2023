package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := insertData(file)
	solutions(result[0], result[1])
}

func insertData(file []byte) [][]int {
	lines := strings.Split(string(file), "\n")
	var result [][]int
	for i := 0; i < 2; i++ {
		line := strings.Split(lines[i], ":")[1]
		numbers := strings.Fields(line)
		var numArr []int
		for _, number := range numbers {
			val, _ := strconv.Atoi(number)
			numArr = append(numArr, val)
		}
		result = append(result, numArr)
	}
	return result
}

func solutions(times []int, distances []int) {
	fmt.Println("times", times)
	fmt.Println("distances", distances)
	result := 1
	for i, time := range times {
		counter := 0
		for press := 0; press < time; press++ {
			if press*(time-press) > distances[i] {
				counter++
			}
		}
		if counter > 0 {
			result *= counter
		}
	}
	fmt.Println(result)
}
