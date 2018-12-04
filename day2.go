package main

import "fmt"

type Counter struct {
	twoCount   int
	threeCount int
}

func (counter *Counter) sum() int {
	return counter.twoCount * counter.threeCount
}

func day2() {
	fmt.Println("Day 2")
	lines := readLines("inputs/day2.txt")

	counter := Counter{}

	//part 1
	for _, line := range lines {
		counts := make(map[string]int)
		for _, intLetter := range line {
			letter := string(intLetter)
			count, present := counts[letter]
			if present {
				counts[letter] = count + 1
			} else {
				counts[letter] = 1
			}
		}

		isTwo, isThree := false, false

		fmt.Println(counts)
		for _, v := range counts {
			switch v {
			case 2:
				isTwo = true
			case 3:
				isThree = true
			}
		}

		if isTwo {
			counter.twoCount += 1
		}

		if isThree {
			counter.threeCount += 1
		}

		fmt.Println(counter)
	}
	fmt.Println(counter.sum())

	//part 2
	for _, line1 := range lines {
		for _, line2 := range lines {
			//fmt.Printf("Comparing %s:%s\n", line1, line2)
			differences := 0

			for i, letter1 := range line1 {
				if letter1 != int32(line2[i]) {
					differences++
				}
				if differences > 1 {
					break
				}
			}

			if differences == 1 {
				fmt.Printf("Found %s:%s\n", line1, line2)
			}
		}
	}
}
