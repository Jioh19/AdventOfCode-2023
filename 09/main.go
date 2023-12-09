package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type History struct {
	data []int
}
type Report struct {
	histories []History
}

func main() {
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	r := insertData(file)

	fmt.Println("Part 1:", calcResult(r))
	fmt.Println("Part 2:", calcResultPrev(r))
}

func insertData(file []byte) Report {
	r := new(Report)
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	for _, line := range lines {
		//fmt.Println(line)
		numbers := strings.Fields(strings.TrimSpace(string(line)))
		h := new(History)
		for _, number := range numbers {
			num, _ := strconv.Atoi(number)
			h.data = append(h.data, num)
		}
		r.histories = append(r.histories, *h)
	}
	return *r
}

func calcNext(h History) int {
	if h.data[len(h.data)-1] == 0 && h.data[0] == 0 {
		return 0
	}
	res := new(History)
	//var result = 0
	for i := 0; i < len(h.data)-1; i++ {
		diff := h.data[i+1] - h.data[i]
		res.data = append(res.data, diff)
	}
	//fmt.Println(h.data)
	//fmt.Println(res.data)
	return calcNext(*res) + h.data[len(h.data)-1]
}

func calcPrev(h History) int {
	if h.data[len(h.data)-1] == 0 && h.data[0] == 0 {
		return 0
	}
	res := new(History)
	//var result = 0
	for i := 0; i < len(h.data)-1; i++ {
		diff := h.data[i+1] - h.data[i]
		res.data = append(res.data, diff)
	}
	//fmt.Println(h.data)
	//fmt.Println(res.data)
	return h.data[0] - calcPrev(*res)
}

func calcResult(r Report) int {
	total := 0
	for _, history := range r.histories {
		total += calcNext(history)
	}
	return total
}

func calcResultPrev(r Report) int {
	total := 0
	for _, history := range r.histories {
		total += calcPrev(history)
	}
	return total
}
