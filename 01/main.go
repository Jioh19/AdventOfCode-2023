package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	const fileName = "input.txt"
	var numbers [][2]int

	//fmt.Println(fileName)

	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := string(file)
	//fmt.Println(len(lines))
	index, aux := 0, 0

	for i, value := range lines {

		if num, len := numString(lines[i:]); num != -1 {
			fmt.Println(i, string(value), index, num, len)
			if aux == 0 {
				values := [2]int{num, num}
				numbers = append(numbers, values)
				aux++
			} else {
				numbers[index][1] = num
			}

		}

		if value >= '0' && value <= '9' {
			//	fmt.Println(i, value)
			if aux == 0 {
				values := [2]int{int(value - 48), int(value - 48)}
				numbers = append(numbers, values)
				aux++
			} else {
				numbers[index][1] = int(value - 48)
			}
		}
		if value == '\n' {
			index++
			aux = 0
		}
	}

	var result int
	for i, value := range numbers {
		fmt.Println(i, int(value[0]), int(value[1]))
		result += int(value[0])*10 + int(value[1])
	}

	numString("tsaonederful")
	fmt.Println(result)
}

func numString(lines string) (int, int) {
	numbers := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	//if lines[5:]
	for word, num := range numbers {
		if len(lines) > len(word) {
			if strings.Contains(lines[:len(word)], word) {
				return num, len(word)
			}
		}
	}
	return -1, -1
}
