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

func main() {
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	m := make(map[string]Node)
	movement := insertData(file, m)
	// fmt.Println(movement)
	// fmt.Println(m)
	fmt.Println(navigate(movement, m))
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
		//fmt.Println(left, key, right)
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
