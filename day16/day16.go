package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var start = time.Now()

func main() {
	inputString := readFile()
	start = time.Now()
	inputPattern := []int{0, 1, 0, -1}

	inputOriginal := make([]int, len(inputString))

	for i, intString := range strings.Split(inputString, "") {
		parsedInt, err := strconv.Atoi(intString)
		if err != nil {
			log.Fatal("cannot parse int", intString)
		}
		inputOriginal[i] = parsedInt
	}

	input := make([]int, len(inputOriginal))
	copy(input, inputOriginal)

	for i := 0; i < 100; i++ {
		input = doPhaseA(input, inputPattern)
	}
	fmt.Println("Day 16A", input[:8], time.Since(start))
	start = time.Now()

	offset, err := strconv.Atoi(inputString[0:7])
	if err != nil {
		log.Fatal("can't parse offset", inputString[0:7])
	}

	input = []int{}
	for i := 0; i < 10000; i++ {
		input = append(input, inputOriginal...)
	}
	input = input[offset:]

	for i := 0; i < 100; i++ {
		input = doPhaseB(input)
	}
	fmt.Println("Day 16B", input[0:8], time.Since(start))

}

func doPhaseB(input []int) []int {
	for i := len(input) - 2; i >= 0; i-- {
		input[i] = (input[i] + input[i+1]) % 10
	}
	return input
}

func doPhaseA(input []int, inputPattern []int) []int {
	result := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		pattern := multiPattern(i+1, len(input), inputPattern)
		digit := 0
		for j := 0; j < len(input); j++ {
			digit += input[j] * pattern[j]
		}
		digit = abs(digit % 10)
		result[i] = digit
	}
	return result
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func multiPattern(num int, length int, inputPattern []int) []int {
	result := make([]int, length+1)

	for i := 0; i < len(result); i++ {
		result[i] = inputPattern[i/num%len(inputPattern)]
	}
	return result[1:]

}

func readFile() string {
	file, err := os.Open("./day16/day16.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return scanner.Text()
}
