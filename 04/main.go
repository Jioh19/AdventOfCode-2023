package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Hands struct {
	hand [][]int
}

func main() {
	fileName := "input.txt"

	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	var count int

	lines := strings.Split(string(file), "\n")
	winner, myHand, winnings := insertHands(lines)

	for i, winnerSet := range winner.hand {
		var aux float64
		var power float64
		for _, win := range winnerSet {
			for _, mySet := range myHand.hand[i] {
				if win == mySet {
					aux = math.Pow(2, power)
					power++
					//fmt.Println(winnerSet, i, aux)
					//	fmt.Println("Plus 1")
				}
			}
		}
		fmt.Println(power)
		for w := 1; w < int(power)+1; w++ {
			if w+i >= len(winnings) {
				break
			}
			winnings[i+w] += winnings[i]
		}
		count += int(aux)
	}
	fmt.Println(count)
	fmt.Println(winnings)
	fmt.Println(countWinnings(winnings))
}

func insertHands(lines []string) (*Hands, *Hands, []int) {
	winner := new(Hands)
	myHand := new(Hands)
	var winnings []int
	for i, line := range lines {
		winnings = append(winnings, 1)
		lines[i] = strings.Split(line, ":")[1]
		lineWinner := strings.Split(lines[i], " | ")[0]
		lineMyHand := strings.Split(lines[i], " | ")[1]
		//fmt.Println(lineMyHand)
		var arrWinner []string
		var arrIntH []int
		var arrIntW []int
		for ix := 0; ix <= len(lineWinner)-3; ix += 3 {
			//fmt.Println(i, lineWinner[ix-3:ix], len(lineWinner))
			arrWinner = append(arrWinner, strings.TrimSpace(lineWinner[ix:ix+3]))
			val, _ := strconv.Atoi(strings.TrimSpace(lineWinner[ix : ix+3]))
			arrIntW = append(arrIntW, val)
		}
		//var arrMyHand []string

		arrMyHand := strings.Split(lineMyHand, " ")
		for i, myHand := range arrMyHand {
			arrMyHand[i] = strings.TrimSpace(myHand)
			val, _ := strconv.Atoi(arrMyHand[i])
			arrIntH = append(arrIntH, val)
		}
		// for ix := 0; ix <= len(lineMyHand)-3; ix += 3 {
		// 	arrMyHand = append(arrMyHand, strings.TrimSpace(lineMyHand[ix:ix+3]))
		// 	val, _ := strconv.Atoi(strings.TrimSpace(lineMyHand[ix : ix+3]))
		// 	arrIntH = append(arrIntH, val)
		// 	//fmt.Println(arrMyHand, len(lineMyHand), ix)
		// }
		myHand.hand = append(myHand.hand, arrIntH)
		winner.hand = append(winner.hand, arrIntW)
		//	fmt.Println(arrIntW)
	}
	//fmt.Println(winner.hand[1])
	fmt.Println(winnings)
	return winner, myHand, winnings
}

func countWinnings(winnings []int) int {
	var counter int
	for _, value := range winnings {
		counter += value
	}
	return counter
}
