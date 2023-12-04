package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	schema []string
	length int
	height int
}

func main() {
	fileName := "test.txt"

	board := new(Board)
	var total int
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	board.schema = strings.Split(string(file), "\n")
	//for i, line := range board.schema {
	// 	fmt.Println(i, line)
	// }
	board.length = len(board.schema[0]) - 1
	board.height = len(board.schema)
	//fmt.Println(board.length, board.height)
	//fmt.Println(findSymbol(board, 6, 2))
	for y := range board.schema {
		for x := 0; x < board.length; x++ {
			found, amount, jump := findSymbol(board, x, y)
			x += jump
			if found {
				total += amount
				fmt.Println(amount, x, y, total)
			} else {
				//total += amount
			}
		}
	}
	fmt.Println(total)
}

func checkSymbol(board *Board, x int, y int) bool {
	if x >= board.length || x < 0 || y >= board.height || y < 0 {
		return false
	}
	if !(board.schema[y][x] >= '0' && board.schema[y][x] <= '9') && board.schema[y][x] != '.' {
		return true
	}
	return false
}

func findSymbol(board *Board, x int, y int) (bool, int, int) {
	var found bool
	var aux string
	var amount int
	var jump int
	found = found || checkSymbol(board, x-1, y-1)
	found = found || checkSymbol(board, x-1, y)
	found = found || checkSymbol(board, x-1, y+1)
	for board.schema[y][x] >= '0' && board.schema[y][x] <= '9' {
		//fmt.Println(x, y)
		jump++
		found = found || checkSymbol(board, x, y-1)
		found = found || checkSymbol(board, x, y)
		found = found || checkSymbol(board, x, y+1)
		aux += string(board.schema[y][x])
		x++
		if x == board.length {
			break
		}
	}
	found = found || checkSymbol(board, x, y-1)
	found = found || checkSymbol(board, x, y)
	found = found || checkSymbol(board, x, y+1)
	amount, _ = strconv.Atoi(aux)
	return found, amount, jump
}
