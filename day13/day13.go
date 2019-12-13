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
var paddleInput = 0
var grid = [40][40]string{}
var score = 0
var paddleX = -1
var ballX = 0

func main() {
	stack := readFile()

	//input := make(chan int)
	//output := make(chan int)
	//
	//go func() {
	//	stackcopy := make([]int, 1000000)
	//	copy(stackcopy, stack)
	//	solve(stackcopy, input, output)
	//	<-input
	//}()
	//
	//blockCount := 0
	//
	//for running {
	//	x := <-output
	//	y := <-output
	//	tile := <-output
	//
	//	switch tile {
	//	case 0:
	//		grid[x][y] = " "
	//	case 1:
	//		grid[x][y] = "█"
	//	case 2:
	//		grid[x][y] = "░"
	//		blockCount++
	//	case 3:
	//		grid[x][y] = "═"
	//	case 4:
	//		grid[x][y] = "*"
	//	default:
	//		log.Fatal("Unknown tile ", tile)
	//	}
	//}
	//
	//fmt.Println("Day13A", blockCount)

	//for x := 0; x < 20; x++ {
	//	for y := 0; y < 40; y++ {
	//		 fmt.Print(grid[y][x])
	//	}
	//	fmt.Println()
	//}

	input2 := make(chan int)
	output2 := make(chan int)

	go func() {
		stackcopy := make([]int, 100000)
		copy(stackcopy, stack)
		stackcopy[0] = 2
		solve(stackcopy, input2, output2)
		<-input2
	}()

	debugCount := 0

	running = true
	for running {
		x := <-output2
		y := <-output2
		tile := <-output2
		//fmt.Println("a", x,y, tile, debugCount)
		debugCount++

		if x == -1 && y == 0 {
			score = tile
		} else {
			switch tile {
			case 0:
				grid[x][y] = " "
			case 1:
				grid[x][y] = "█"
			case 2:
				grid[x][y] = "░"
			case 3:
				grid[x][y] = "═"
				paddleX = x
			case 4:
				grid[x][y] = "o"
				if x-paddleX < 0 {
					paddleInput = -1
				} else if x-paddleX > 0 {
					paddleInput = 1
				} else {
					paddleInput = 0
				}
			default:
				log.Fatal("Unknown tile ", tile)
			}
		}

		if tile == 4 && paddleX > 0 || tile == 3 && paddleX == 0 {
		}

	}

	//fmt.Println("Day13A", blockCount)
}

func printScreen(grid [40][40]string, score int, ballX int, paddleX int, paddleInput int) {
	output := ""
	for x := 0; x < 20; x++ {
		for y := 0; y < 40; y++ {
			output += grid[y][x]
		}
		output += "\n"
	}
	fmt.Print(output)
	fmt.Println("Score:", score, "BallX:", ballX, " PaddleX:", paddleX, " PaddleInput", paddleInput)
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
			close(output)
			fmt.Println("Done processing", score)
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
			printScreen(grid, score, ballX, paddleX, paddleInput)
			stack[destination] = paddleInput
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
	file, err := os.Open("./day13/day13.txt")
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
