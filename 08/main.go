package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	right string
	left  string
}

// Calculate the greatest common divisor (GCD) using the Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Calculate the least common multiple (LCM) of two integers
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

// Calculate the LCM of a slice of integers
func lcmOfSlice(numbers []int) int {
	result := 1
	for _, num := range numbers {
		result = lcm(result, num)
	}
	return result
}

func main() {
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	m := make(map[string]Node)
	movement := insertData(file, m)
	part2 := navigate2(movement, m, selectStart(m))
	fmt.Println("Resultado de la parte 2:", lcmOfSlice(part2))
}

func insertData(file []byte, m map[string]Node) string {
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	for _, line := range lines[2:] {
		key := strings.Split(strings.TrimSpace(string(line)), "=")[0]
		key = strings.TrimSpace(key)
		data := strings.Split(strings.TrimSpace(string(line)), "=")[1]
		left := strings.Trim(strings.TrimSpace(strings.Split(data, ",")[0]), "(")
		right := strings.Trim(strings.Split(data, ",")[1], ")")
		right = strings.TrimSpace(right)
		node := new(Node)
		node.left = left
		node.right = right
		m[key] = *node

	}
	return lines[0]
}

func navigate(movement string, m map[string]Node) int {
	result := 0
	i := 0
	key := "AAA"
	for key != "ZZZ" {
		if movement[i] == 'L' {
			key = m[key].left
		}
		if movement[i] == 'R' {
			key = m[key].right
		}
		i++
		if i == len(movement) {
			i = 0
		}
		result++
	}
	return result
}

func selectStart(m map[string]Node) []string {
	var mA []string
	for k := range m {
		if k[2] == 'A' {
			mA = append(mA, k)
		}
	}
	return mA
}

func checkStatus(status []string) bool {
	for _, stat := range status {
		if stat[2] != 'Z' {
			return false
		}
	}
	return true
}

func navigate2(movement string, m map[string]Node, start []string) []int {
	status := start
	var arr []int

	for j := 0; j < len(start); j++ {
		result := 0
		i := 0
		for status[j][2] != 'Z' {
			if movement[i] == 'L' {
				status[j] = m[status[j]].left
			}
			if movement[i] == 'R' {
				status[j] = m[status[j]].right
			}
			i++
			if i == len(movement) {
				i = 0
			}
			result++
		}
		arr = append(arr, result)
	}
	fmt.Println(arr)
	return arr
}
