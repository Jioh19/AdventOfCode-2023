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
	winner, myHand := insertHands(lines)

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
	fmt.Println(myHand.hand[1])
	return winner, myHand
}
