package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const fileName = "input.txt"
	var checker = [3]int{12, 13, 14}
	var counter int
	var power int
	//fmt.Println(fileName)

	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	//row := 0
	lines := strings.Split(string(file), "\n")
	for i, line := range lines {
		lines[i] = strings.Split(line, ":")[1]
		_, valid := valid(lines[i], checker)
		//	fmt.Println(i, num, valid)
		if valid {
			counter += i + 1
		}
		num2 := valid2(lines[i])
		power += num2
	}
	// for _, val := range lines {
	// 	if val == '\n' || val == ',' || val == ';' {

	// 	}
	// }

	fmt.Println("Part 1:", counter)
	fmt.Println("Part 2:", power)
	//fmt.Println(valid(lines[9], checker))
}

func valid(line string, checker [3]int) (int, bool) {
	var (
		valid bool = true
		//val   [3]int
		//aux   string
	)
	sets := strings.Split(line, ";")
	for _, set := range sets {
		// fmt.Println(set)
		// fmt.Println(eval(set))
		for i := range checker {
			if checker[i] < eval(set)[i] {
				valid = false
			}
		}

	}
	// num, _ = strconv.Atoi(aux)
	// fmt.Println(num + 1)
	return len(sets), valid
}

func valid2(line string) int {

	var result [3]int
	sets := strings.Split(line, ";")
	for _, set := range sets {
		// fmt.Println(set)
		// fmt.Println(eval(set))
		for i := range result {
			if result[i] < eval(set)[i] {
				result[i] = eval(set)[i]
			}
		}

	}

	// num, _ = strconv.Atoi(aux)
	// fmt.Println(num + 1)
	return result[0] * result[1] * result[2]
}

func eval(set string) [3]int {

	var val [3]int

	//var valid bool
	numbers := map[string]int{
		"red":   0,
		"green": 1,
		"blue":  2,
	}
	color := strings.Split(set, ",")

	for _, ball := range color {
		aux := ""
		for _, letter := range ball {
			if letter >= '0' && letter <= '9' {
				aux = aux + string(letter)
			}
		}
		num, _ := strconv.Atoi(aux)
		//fmt.Println(num, ball)
		//if sets[5:]
		for word, index := range numbers {

			if strings.Contains(ball, word) {
				//fmt.Println(word, index, num)
				val[index] = num
			}
		}
	}
	return val
}
