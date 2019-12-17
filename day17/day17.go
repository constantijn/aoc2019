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

	go func() {
		stackcopy := make([]int, 1000000)
		copy(stackcopy, stack)
		solve(stackcopy, input, output)
	}()

	grid := [47][47]string{}
	x, y := 0, 0

	for true {
		intChar, ok := <-output
		if !ok {
			break
		}
		stringChar := string(intChar)

		if stringChar == "\n" {
			x = 0
			y++
		} else {
			grid[x][y] = stringChar
			x++
		}
	}

	alignment := 0

	for y := 1; y < 45; y++ {
		for x := 1; x < 45; x++ {
			if grid[x][y] == "#" && grid[x+1][y] == "#" && grid[x-1][y] == "#" && grid[x][y+1] == "#" && grid[x][y-1] == "#" {
				alignment += x * y
			}
		}
	}

	fmt.Println("Day 17A", alignment, time.Since(start))

	start = time.Now()
	input2 := make(chan int)
	output2 := make(chan int)

	go func() {
		stackcopy := make([]int, 1000000)
		copy(stackcopy, stack)
		stackcopy[0] = 2
		solve(stackcopy, input2, output2)
	}()

	const mainRoutine = "A,B,A,B,A,C,B,C,A,C"
	const aRoutine = "L,10,L,12,R,6"
	const bRoutine = "R,10,L,4,L,4,L,12"
	const cRoutine = "L,10,R,10,R,6,L,4"
	const video = "n"

	go func() {
		transmit(mainRoutine, input2)
		transmit(aRoutine, input2)
		transmit(bRoutine, input2)
		transmit(cRoutine, input2)
		transmit(video, input2)
	}()

	lastInt := -1
	for true {
		intChar, ok := <-output2
		if !ok {
			break
		}
		lastInt = intChar
		//stringChar := string(intChar)
		//fmt.Print(stringChar)
	}

	fmt.Println("Day 17B", lastInt, time.Since(start))

}

func transmit(data string, channel chan int) {
	if len(data) > 20 {
		log.Fatal("Data string too long. len: ", len(data), " Data: ", data)
	}
	for i := 0; i < len(data); i++ {
		channel <- int(data[i])
	}
	channel <- 10 // newline
}

func print(grid [47][47]string) {

	for i := 0; i < 47; i++ {
		for j := 0; j < 47; j++ {
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

		if opcode == 99 {
			close(output)
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
	file, err := os.Open("./day17/day17.txt")
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
