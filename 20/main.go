package main

import (
	"fmt"
	"os"
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
	fmt.Println("Time in nanoseconds:", time.Since(s).Nanoseconds())
}
