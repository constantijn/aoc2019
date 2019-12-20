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
	start := time.Now()

	input := make(chan int)
	output := make(chan int)

	grid := [50][50]string{}
	stackcopy := make([]int, 500)

	count := 0
	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			go func() {
				copy(stackcopy, stack)
				solve(stackcopy, input, output)
			}()
			input <- x
			input <- y
			result := <-output
			if result == 0 {
				grid[x][y] = "."
			} else {
				grid[x][y] = "X"
			}
			count += result
		}
	}
	//print(grid)
	fmt.Println("Day 19A", count, time.Since(start))
	start = time.Now()
	grid2 := [1500][1500]string{}

	count = 0
	for y := 0; y < 1500; y++ {
		for x := 0; x < 1500; x++ {
			go func() {
				copy(stackcopy, stack)
				solve(stackcopy, input, output)
			}()
			input <- x
			input <- y
			result := <-output
			if result == 0 {
				grid2[x][y] = "."
			} else {
				grid2[x][y] = "X"
			}
			count += result
		}
	}

	//print2(grid2)

	for x := 0; x < 1399; x++ {
		for y := 0; y < 1399; y++ {
			if grid2[x][y] == "X" && grid2[x+99][y] == "X" && grid2[x][y+99] == "X" && grid2[x+100][y] == "." && grid2[x][y+100] == "." {
				fmt.Println("Day 19B", x*10000+y, time.Since(start))
				os.Exit(0)
			}
		}
	}
}

func print2(grid [1500][1500]string) {
	for i := 0; i < 1500; i++ {
		for j := 0; j < 1500; j++ {
			fmt.Print(grid[j][i])
		}
		fmt.Println()
	}
}

func print(grid [50][50]string) {

	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			fmt.Print(grid[j][i])
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

		//fmt.Println("instruction", instruction)
		if opcode == 99 {
			//fmt.Println("Done processing")
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
			//fmt.Println("reading input")
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
	file, err := os.Open("./day19/day19.txt")
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
