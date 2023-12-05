package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Hands struct {
	hand [][]string
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
	winner, myHand := insertHands(lines)

	for i, winnerSet := range winner.hand {
		var aux float64
		var power float64
		for _, win := range winnerSet {
			for _, mySet := range myHand.hand[i] {
				if win == mySet {
					aux = math.Pow(2, power)
					power++
					fmt.Println(winnerSet, i, aux)
					//	fmt.Println("Plus 1")
				}
			}
		}
		//fmt.Println(aux)
		count += int(aux)
	}
	fmt.Println(count)
}

func insertHands(lines []string) (*Hands, *Hands) {
	winner := new(Hands)
	myHand := new(Hands)
	for i, line := range lines {
		lines[i] = strings.Split(line, ":")[1]
		lineWinner := strings.Split(lines[i], " | ")[0]
		lineMyHand := strings.Split(lines[i], " | ")[1]
		//fmt.Println(lineMyHand)
		var arrWinner []string
		for ix := 0; ix <= len(lineWinner)-3; ix += 3 {
			//fmt.Println(i, lineWinner[ix-3:ix], len(lineWinner))
			arrWinner = append(arrWinner, strings.TrimSpace(lineWinner[ix:ix+3]))
		}
		var arrMyHand []string
		for ix := 0; ix <= len(lineMyHand)-3; ix += 3 {
			arrMyHand = append(arrMyHand, strings.TrimSpace(lineMyHand[ix:ix+3]))
			//	fmt.Println(arrMyHand, len(lineMyHand), ix)
		}
		myHand.hand = append(myHand.hand, arrMyHand)
		winner.hand = append(winner.hand, arrWinner)
	}
	//fmt.Println(len(myHand.hand[0][2]))
	return winner, myHand
}
