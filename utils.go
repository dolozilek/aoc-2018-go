package main

import (
	"bufio"
	"io/ioutil"
	"os"
)

func readInput(name string) (result string) {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	result = string(b)
	return
}

func readLines(name string) (lines []string) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
