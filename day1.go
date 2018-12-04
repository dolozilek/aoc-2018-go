package main

import (
	"fmt"
	"strconv"
)

func day1() {
	fmt.Println("Day 1")
	lines := readLines("inputs/day1.txt")

	//part 1
	sum := 0
	for _, line := range lines {
		lineInt, err := strconv.Atoi(line)
		panicErr(err)
		sum += lineInt
	}
	fmt.Println(sum)

	//part 2
	i := 0
	linesCount := len(lines)
	found := false
	sum = 0
	frequencies := make(map[int]bool)
	for !found {
		index := i % linesCount
		lineInt, err := strconv.Atoi(lines[index])
		panicErr(err)
		sum += lineInt
		_, present := frequencies[sum]
		if present {
			fmt.Printf("Found: %v", sum)
			found = true
		} else {
			frequencies[sum] = true
		}
		i++
	}

}
