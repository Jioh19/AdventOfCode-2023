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
	s := time.Now()
	springs, broken := insertData(file)
	// for _, row := range springs {
	// 	fmt.Println(row)
	// }
	// for _, row := range broken {
	// 	fmt.Println(row)
	// }
	result := getArrangements(springs, broken)
	fmt.Println("Result:", result)
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())
	s2 := time.Now()
	springs2, broken2 := insertData2(file)
	// for _, row := range springs2 {
	// 	fmt.Println(row)
	// }
	// for _, row := range broken2 {
	// 	fmt.Println(row)
	// }
	result2 := getArrangements2(springs2, broken2)
	fmt.Println("Result 2:", result2)
	fmt.Println("Time in nanoseconds:", time.Since(s2).Nanoseconds())
}

func insertData(file []byte) ([]string, [][]int) {
	var springs []string
	var broken [][]int
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(string(line))
		spring := strings.Split(line, " ")[0]
		strNum := strings.Split(strings.Split(line, " ")[1], ",")
		var numbers []int
		for _, num := range strNum {
			number, _ := strconv.Atoi(num)
			numbers = append(numbers, number)
		}
		springs = append(springs, spring)
		broken = append(broken, numbers)
	}
	return springs, broken
}

func insertData2(file []byte) ([]string, [][]int) {
	var springs []string
	var broken [][]int
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(string(line))
		spring := strings.Split(line, " ")[0]
		spring = strings.Repeat(spring+"?", 5)
		strNum := strings.Split(line, " ")[1] + ","
		strNum = strings.Repeat(strNum, 5)
		numStr := strings.Split(strNum, ",")
		var numbers []int
		for _, num := range numStr {
			number, _ := strconv.Atoi(num)
			numbers = append(numbers, number)
		}
		springs = append(springs, spring)
		broken = append(broken, numbers)
	}
	return springs, broken
}

func getArrangements(springs []string, broken [][]int) int {
	result := 0
	for i, spring := range springs {
		result += countArrangement(spring, broken[i])
		fmt.Println("Cycle:", i)
	}
	return result
}

func getArrangements2(springs []string, broken [][]int) int {
	result := 0
	for i, spring := range springs {
		result += getCache(spring, broken[i])
		fmt.Println("Cycle:", i)
	}
	return result
}

func getCache(spring string, broke []int) int {
	var cache [][]int
	for i := 0; i < len(spring); i++ {
		cache = append(cache, make([]int, len(spring)+1))
		for j := 0; j < len(spring)+1; j++ {
			cache[i][j] = -1
		}
	}
	return countArrangementCache(0, 0, spring, broke, cache)
}

func countArrangement(spring string, broke []int) int {
	if spring == "" {
		if len(broke) == 0 {
			return 1
		}
		return 0
	}
	if len(broke) == 0 {
		if strings.Contains(spring, "#") {
			return 0
		}
		return 1
	}
	result := 0
	if spring[0] == '.' || spring[0] == '?' {
		result += countArrangement(spring[1:], broke)
	}
	if spring[0] == '#' || spring[0] == '?' {
		if broke[0] <= len(spring) && !strings.Contains(spring[:broke[0]], ".") && (broke[0] == len(spring) || spring[broke[0]] != '#') {
			if broke[0] == len(spring) {
				result += countArrangement("", broke[1:])
			} else {
				result += countArrangement(spring[broke[0]+1:], broke[1:])
			}
		}
	}
	return result
}

func countArrangementCache(i, j int, spring string, broke []int, cache [][]int) int {
	if i >= len(spring) {
		if j < len(broke) {
			return 0
		}
		return 1
	}

	if cache[i][j] != -1 {
		return cache[i][j]
	}

	res := 0
	if spring[i] == '.' {
		res = countArrangementCache(i+1, j, spring, broke, cache)
	} else {
		if spring[i] == '?' {
			res += countArrangementCache(i+1, j, spring, broke, cache)
		}
		if j < len(broke) {
			count := 0
			for k := i; k < len(spring); k++ {
				if count > broke[j] || spring[k] == '.' || count == broke[j] && spring[k] == '?' {
					break
				}
				count += 1
			}
			if count == broke[j] {
				if i+count < len(spring) && spring[i+count] != '#' {
					res += countArrangementCache(i+count, j+1, spring, broke, cache)
				} else {
					res += countArrangementCache(i+count, j+1, spring, broke, cache)
				}
			}
		}
	}
	cache[i][j] = res
	if res < 0 {
		fmt.Println("i, j", cache[i][j])
	}
	return res

}
