package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Lens struct {
	name  string
	focus int
}

func main() {
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	input := insertData(file)

	s := time.Now()
	part1 := sum(input)
	fmt.Println("Part 1:", part1)
	fmt.Println("Time in milliseconds:", time.Since(s).Milliseconds())
	s2 := time.Now()
	part2 := sum2(input)
	fmt.Println("Part 2:", part2)
	fmt.Println("Time in milliseconds:", time.Since(s2).Milliseconds())
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
		result += val
	}
	return result
}

func sum2(inputs [][]byte) int {
	result := 0
	m := make(map[int][]Lens)
NEXT:
	for _, input := range inputs {
		if len(input) > 0 {
			arrInput := strings.Split(string(input), "=")
			if len(arrInput) == 2 {
				split := arrInput[0]

				focus := arrInput[1]
				lens, _ := strconv.Atoi(string(focus))
				val := hash([]byte(split))
				for i, len := range m[val] {
					if split == len.name {
						m[val][i].focus = lens
						continue NEXT
					}
				}
				m[val] = append(m[val], Lens{string(split), lens})
			} else {
				split := strings.Split(string(input), "-")[0]
				val := hash([]byte(split))
				for i, len := range m[val] {
					if split == len.name {
						m[val] = append(m[val][:i], m[val][i+1:]...)
					}
				}
			}
		}
	}
	result = calc(m)
	return result
}

func calc(m map[int][]Lens) int {
	result := 0
	for i, box := range m {
		for j, lens := range box {
			result += (i + 1) * (j + 1) * lens.focus
		}
	}
	return result
}
