package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	m, p := insertData(file)

	part1 := allFlows(m, p)
	s := time.Now()
	fmt.Println("Part 1", part1)
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())
}

func allFlows(m map[string][]string, p []map[string]int) int {
	result := 0
	for _, part := range p {
		result += matchFlow(m, "in", part)
	}
	return result
}

func matchFlow(m map[string][]string, key string, p map[string]int) int {
	if key == "A" {
		return p["x"] + p["m"] + p["a"] + p["s"]
	}
	if key == "R" {
		return 0
	}
	for i, flow := range m[key] {
		var part string
		var nextKey string
		var value int
		if strings.Contains(flow, "<") {
			part = strings.Split(flow, "<")[0]
			flow = strings.Split(flow, "<")[1]
			strVal, _ := strconv.Atoi(strings.Split(flow, ":")[0])
			value = strVal
			nextKey = strings.Split(flow, ":")[1]
			if p[part] < value {
				return matchFlow(m, nextKey, p)
			}
		} else if strings.Contains(flow, ">") {
			part = strings.Split(flow, ">")[0]
			flow = strings.Split(flow, ">")[1]
			strVal, _ := strconv.Atoi(strings.Split(flow, ":")[0])
			value = strVal
			nextKey = strings.Split(flow, ":")[1]
			if p[part] > value {
				return matchFlow(m, nextKey, p)
			}
		} else {
			nextKey = flow
			return matchFlow(m, nextKey, p)
		}
		fmt.Println(i, flow, part, nextKey, value)
	}
	return 0
}

func insertData(file []byte) (map[string][]string, []map[string]int) {
	m := make(map[string][]string)
	parts := strings.Split(strings.TrimSpace(string(file)), "\n\n")
	workFlows := strings.Split(parts[0], "\n")
	machineParts := strings.Split(strings.TrimSpace(parts[1]), "\n")
	var p []map[string]int
	for _, line := range workFlows {
		key := strings.Split(line, "{")[0]
		rule := strings.Split(line, "{")[1]
		rule = strings.Split(rule, "}")[0]
		rules := strings.Split(rule, ",")
		m[key] = rules
	}
	for _, line := range machineParts {
		mAux := make(map[string]int)
		partAux := strings.Split(line, "{")[1]
		partAux = strings.Split(partAux, "}")[0]
		mParts := strings.Split(partAux, ",")
		for _, part := range mParts {
			key := strings.Split(part, "=")[0]
			strNum := strings.Split(part, "=")[1]
			num, _ := strconv.Atoi(strNum)
			mAux[key] = num
		}
		//fmt.Println(mParts)
		p = append(p, mAux)
	}
	return m, p
}
