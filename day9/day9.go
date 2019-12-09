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
	stackcopy := make([]int, len(stack)+100000)

	copy(stackcopy, stack)
	start := time.Now()
	fmt.Println("Day 9A", solve(stackcopy, 1), time.Since(start))
	copy(stackcopy, stack)
	start = time.Now()
	fmt.Println("Day 9B", solve(stackcopy, 2), time.Since(start))

}

func solve(stack []int, input int) []int {
	counter := 0
	var result []int
	relativeBase := 0

	for true {
		instruction := stack[counter]
		opcode := instruction % 100
		mode1 := instruction / 100 % 10
		mode2 := instruction / 1000 % 10
		mode3 := instruction / 10000 % 10

		if opcode == 99 {
			break
		}

		if opcode == 1 {
			operand1 := operandValue(stack, counter+1, mode1, relativeBase)
			operand2 := operandValue(stack, counter+2, mode2, relativeBase)
			destination := destinationValue(stack, counter+3, mode3, relativeBase)
			stack[destination] = operand1 + operand2
			counter += 4
		} else if opcode == 2 {
			operand1 := operandValue(stack, counter+1, mode1, relativeBase)
			operand2 := operandValue(stack, counter+2, mode2, relativeBase)
			destination := destinationValue(stack, counter+3, mode3, relativeBase)
			stack[destination] = operand1 * operand2
			counter += 4
		} else if opcode == 3 {
			destination := destinationValue(stack, counter+1, mode1, relativeBase)
			stack[destination] = input
			counter += 2
		} else if opcode == 4 {
			output := operandValue(stack, counter+1, mode1, relativeBase)
			result = append(result, output)
			counter += 2
		} else if opcode == 5 {
			operand1 := operandValue(stack, counter+1, mode1, relativeBase)
			operand2 := operandValue(stack, counter+2, mode2, relativeBase)
			if operand1 != 0 {
				counter = operand2
			} else {
				counter += 3
			}
		} else if opcode == 6 {
			operand1 := operandValue(stack, counter+1, mode1, relativeBase)
			operand2 := operandValue(stack, counter+2, mode2, relativeBase)
			if operand1 == 0 {
				counter = operand2
			} else {
				counter += 3
			}
		} else if opcode == 7 {
			operand1 := operandValue(stack, counter+1, mode1, relativeBase)
			operand2 := operandValue(stack, counter+2, mode2, relativeBase)
			destination := destinationValue(stack, counter+3, mode3, relativeBase)
			if operand1 < operand2 {
				stack[destination] = 1
			} else {
				stack[destination] = 0
			}
			counter += 4
		} else if opcode == 8 {
			operand1 := operandValue(stack, counter+1, mode1, relativeBase)
			operand2 := operandValue(stack, counter+2, mode2, relativeBase)
			destination := destinationValue(stack, counter+3, mode3, relativeBase)
			if operand1 == operand2 {
				stack[destination] = 1
			} else {
				stack[destination] = 0
			}
			counter += 4
		} else if opcode == 9 {
			operand1 := operandValue(stack, counter+1, mode1, relativeBase)
			relativeBase += operand1
			counter += 2
		} else {
			log.Fatal("Unknown opcode ", opcode)
		}
	}

	return result
}

func destinationValue(stack []int, index int, mode int, relativeBase int) int {
	result := -1
	if mode == 0 {
		result = stack[index]
	} else if mode == 2 {
		result = stack[index] + relativeBase
	} else {
		log.Fatal("Invalid mode for destination ", mode)
	}
	return result
}

func operandValue(stack []int, index int, mode int, relativeBase int) int {
	result := -1
	if mode == 0 {
		result = stack[stack[index]]
	} else if mode == 1 {
		result = stack[index]
	} else if mode == 2 {
		result = stack[stack[index]+relativeBase]
	} else {
		log.Fatal("Invalid mode ", mode)
	}
	return result
}

func readFile() []int {
	file, err := os.Open("./day9/day9.txt")
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
			log.Fatal("Can't parse to int: " + scanner.Text())
		}

		result = append(result, n)
	}
	return result
}
