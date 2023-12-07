package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Game struct {
	hand string
	bet  int
	//	high  int
	value int
}

// ByValue is a type for sorting Game slice by Value.
type ByValue []Game

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].value < a[j].value }

// ByHand is a type for sorting Game slice by Hand.
type ByHand []Game

func (a ByHand) Len() int           { return len(a) }
func (a ByHand) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByHand) Less(i, j int) bool { return a[i].hand < a[j].hand }

// MultiSorter is a type that composes multiple sort.Interface implementations.
type MultiSorter struct {
	games   []Game
	sorters []sort.Interface
}

// NewMultiSorter initializes a MultiSorter with a slice of Game and multiple sorters.
func NewMultiSorter(games []Game, sorters ...sort.Interface) *MultiSorter {
	return &MultiSorter{
		games:   games,
		sorters: sorters,
	}
}

// Len returns the length of the underlying slice.
func (ms *MultiSorter) Len() int {
	return len(ms.games)
}

// Swap swaps the elements with indexes i and j.
func (ms *MultiSorter) Swap(i, j int) {
	ms.games[i], ms.games[j] = ms.games[j], ms.games[i]
}

// Less compares elements with indexes i and j using the provided sorters.
func (ms *MultiSorter) Less(i, j int) bool {
	for _, sorter := range ms.sorters {
		if sorter.Less(i, j) {
			return true
		} else if sorter.Less(j, i) {
			return false
		}
	}
	return false
}

func main() {
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := insertData(file)
	_ = result
	fmt.Println(result)
}

func insertData(file []byte) []Game {
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	var games []Game
	for _, line := range lines {
		numbers := strings.Fields(line)
		hand := numbers[0]
		bet, err := strconv.Atoi(numbers[1])
		//		high := getHighCard(hand)
		value := getValue(hand)
		if err != nil {
			continue
		}
		game := Game{hand, bet, value}
		games = append(games, game)
	}
	multiSorter := NewMultiSorter(
		games,
		ByValue(games),
		ByHand(games),
	)
	sort.Sort(multiSorter)
	//games = sorter(games)
	return games
}

func getValue(hand string) int {
	first := 0
	second := 0
	hands := hand
	for len(hands) > 0 {
		if count := strings.Count(hands, string(hands[0])); count > first {
			first, second = count, first
		} else if count > second {
			second = count
		}
		hands = strings.Replace(hands, string(hands[0]), "", 5)
	}
	switch {
	case first == 5:
		return 7
	case first == 4:
		return 6
	case first == 3 && second == 2:
		return 5
	case first == 3 && second == 1:
		return 4
	case first == 2 && second == 2:
		return 3
	case first == 2 && second == 1:
		return 2
	default:
		return 1
	}
}

// func sorter(games []Game) []Game {
// 	sort.SliceStable(games, func(i, j int) bool {
// 		return games[i].value > games[j].value
// 	})
// 	return games
// }
