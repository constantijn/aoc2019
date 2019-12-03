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

func main() {
	stack := readFile()
	stackcopy := make([]int, len(stack))
	copy(stackcopy, stack)

	fmt.Println("2A", solve(stackcopy, 12, 2))

	start := time.Now()

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			copy(stackcopy, stack)
			result := solve(stackcopy, i,j)
			if result == 19690720 {
				fmt.Println("2B", 100*i + j, time.Since(start))
				os.Exit(0)
			}
		}
	}
}

func solve(stack []int, noun int, verb int) int {
	counter := 0
	stack[1] = noun
	stack[2] = verb

	for true {
		opcode := stack[counter]
		if opcode == 99 {
			break
		}
		operand1 := stack[stack[counter+1]]
		operand2 := stack[stack[counter+2]]
		destination := stack[counter+3]


		if opcode == 1 {
			stack[destination] = operand1 + operand2
		} else if opcode == 2 {
			stack[destination] = operand1 * operand2
		} else {
			log.Fatal("Unknown opcode", opcode)
		}
		counter += 4
	}

	return stack[0]
}

func readFile() []int {
	file, err := os.Open("./day2/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result []int
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for _, linePart := range strings.Split(scanner.Text(), ",") {
		n, err := strconv.Atoi(linePart)

		if err != nil {
			log.Fatal("Can't parse to int: " + scanner.Text() )
		}

		result = append(result, n)
	}
	return result
}
