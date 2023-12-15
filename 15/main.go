package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	input := insertData(file)
	// for i := range input {
	// 	for j := range input[i] {
	// 		fmt.Printf("%s", string(input[i][j]))

	// 	}
	// }
	// fmt.Println(input)
	// fmt.Println()
	s := time.Now()
	part1 := sum(input)
	fmt.Println("Part 1:", part1)
	fmt.Println("Time in milliseconds:", time.Since(s).Milliseconds())
}

func insertData(file []byte) [][]byte {
	var input [][]byte
	var word []byte
	for _, letter := range file {
		if letter == ',' || letter == 13 || letter == '\n' {
			input = append(input, word)
			word = []byte{}
		} else {
			word = append(word, letter)
		}
	}
	return input
}

func hash(input []byte) int {
	result := 0
	for _, val := range input {

		result += int(val)
		result *= 17
		result %= 256
	}
	return result
}

func sum(inputs [][]byte) int {
	result := 0
	for _, input := range inputs {
		val := hash(input)
		//fmt.Println(i, val)
		result += val
	}
	return result
}
