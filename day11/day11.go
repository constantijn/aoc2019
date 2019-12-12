package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var running = true

func main() {
	stack := readFile()

	input := make(chan int)
	output := make(chan int)

	visited := map[string]bool{}
	count := 0

	go func() {
		stackcopy := make([]int, 1000000)
		copy(stackcopy, stack)
		solve(stackcopy, input, output)
		<-input
	}()

	grid := [100][100]int{}
	currentX, currentY := 50, 50
	direction := "UP"

	//input <- 0 # Part A
	input <- 1
	for running {
		color := <-output
		turn := <-output

		visited[fmt.Sprint(currentX, ",", currentY)] = true
		count++
		grid[currentX][currentY] = color

		switch x := fmt.Sprint(direction, turn); x {
		case "UP0":
			direction = "LEFT"
		case "UP1":
			direction = "RIGHT"
		case "LEFT0":
			direction = "DOWN"
		case "LEFT1":
			direction = "UP"
		case "DOWN0":
			direction = "RIGHT"
		case "DOWN1":
			direction = "LEFT"
		case "RIGHT0":
			direction = "UP"
		case "RIGHT1":
			direction = "DOWN"
		default:
			log.Fatal("unknown turn ", x)
		}

		switch direction {
		case "UP":
			currentY++
		case "LEFT":
			currentX--
		case "DOWN":
			currentY--
		case "RIGHT":
			currentX++
		}

		if running {
			input <- grid[currentX][currentY]
		}
	}

	fmt.Println("Day11B", len(visited))

	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			if grid[x][y] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Println()
	}

}

func solve(stack []int, input chan int, output chan int) {
	counter := 0
	relativeBase := 0

	for true {
		instruction := stack[counter]
		opcode := instruction % 100
		mode1 := instruction / 100 % 10
		mode2 := instruction / 1000 % 10
		mode3 := instruction / 10000 % 10

		if opcode == 99 {
			running = false
			fmt.Println("Done processing")
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
			stack[destination] = <-input
			counter += 2
		} else if opcode == 4 {
			output <- operandValue(stack, counter+1, mode1, relativeBase)
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
	file, err := os.Open("./day11/day11.txt")
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
