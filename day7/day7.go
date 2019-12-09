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

	start := time.Now()
	maxResult := 0
	var maxSettings []int

	for _, phaseSettings := range getPossiblePhaseSettings([]int{0, 1, 2, 3, 4}) {
		channel1 := make(chan int)
		channel2 := make(chan int)
		channel3 := make(chan int)
		channel4 := make(chan int)
		channel5 := make(chan int)
		channelDone := make(chan int)

		go func() {
			stackcopy := make([]int, len(stack))
			copy(stackcopy, stack)
			solve(stackcopy, channel1, channel2, channelDone)
		}()
		go func() {
			stackcopy := make([]int, len(stack))
			copy(stackcopy, stack)
			solve(stackcopy, channel2, channel3, channelDone)
		}()
		go func() {
			stackcopy := make([]int, len(stack))
			copy(stackcopy, stack)
			solve(stackcopy, channel3, channel4, channelDone)
		}()
		go func() {
			stackcopy := make([]int, len(stack))
			copy(stackcopy, stack)
			solve(stackcopy, channel4, channel5, channelDone)
		}()
		go func() {
			stackcopy := make([]int, len(stack))
			copy(stackcopy, stack)
			solve(stackcopy, channel5, channel1, channelDone)
		}()

		channel1 <- phaseSettings[0]
		channel2 <- phaseSettings[1]
		channel3 <- phaseSettings[2]
		channel4 <- phaseSettings[3]
		channel5 <- phaseSettings[4]
		channel1 <- 0

		result := <-channel1

		if result > maxResult {
			maxResult = result
			maxSettings = phaseSettings
		}
	}

	fmt.Println("Day7A", maxResult, maxSettings, time.Since(start))

	start = time.Now()
	maxResult = 0

	for _, phaseSettings := range getPossiblePhaseSettings([]int{5, 6, 7, 8, 9}) {
		channel1 := make(chan int)
		channel2 := make(chan int)
		channel3 := make(chan int)
		channel4 := make(chan int)
		channel5 := make(chan int)
		channelDone := make(chan int)

		go func() {
			stackcopy := make([]int, len(stack))
			copy(stackcopy, stack)
			solve(stackcopy, channel1, channel2, channelDone)
		}()
		go func() {
			stackcopy := make([]int, len(stack))
			copy(stackcopy, stack)
			solve(stackcopy, channel2, channel3, channelDone)
		}()
		go func() {
			stackcopy := make([]int, len(stack))
			copy(stackcopy, stack)
			solve(stackcopy, channel3, channel4, channelDone)
		}()
		go func() {
			stackcopy := make([]int, len(stack))
			copy(stackcopy, stack)
			solve(stackcopy, channel4, channel5, channelDone)
		}()
		go func() {
			stackcopy := make([]int, len(stack))
			copy(stackcopy, stack)
			solve(stackcopy, channel5, channel1, channelDone)
		}()

		channel1 <- phaseSettings[0]
		channel2 <- phaseSettings[1]
		channel3 <- phaseSettings[2]
		channel4 <- phaseSettings[3]
		channel5 <- phaseSettings[4]
		channel1 <- 0

		<-channelDone
		<-channelDone
		<-channelDone
		<-channelDone

		result := <-channel1

		if result > maxResult {
			maxResult = result
			maxSettings = phaseSettings
		}
	}

	fmt.Println("Day7B", maxResult, maxSettings, time.Since(start))
}

func getPossiblePhaseSettings(options []int) [][]int {
	if len(options) == 1 {
		return [][]int{options}
	}

	var results [][]int
	for i := 0; i < len(options); i++ {
		optionsCopy := make([]int, len(options))
		copy(optionsCopy, options)
		optionsCopy[i], optionsCopy[0] = optionsCopy[0], optionsCopy[i]
		head := optionsCopy[0:1]
		tail := optionsCopy[1:]
		for _, x := range getPossiblePhaseSettings(tail) {
			result := make([]int, len(options))
			copy(result, append(head, x...))
			results = append(results, result)
		}
	}

	return results
}

func solve(stack []int, input chan int, output chan int, done chan int) {
	counter := 0

	for true {
		instruction := stack[counter]
		opcode := instruction % 100
		mode1 := instruction / 100 % 10
		mode2 := instruction / 1000 % 10

		if opcode == 99 {
			done <- 0
			break
		}

		if opcode == 1 {
			operand1 := operandValue(stack, counter+1, mode1)
			operand2 := operandValue(stack, counter+2, mode2)
			destination := stack[counter+3]
			stack[destination] = operand1 + operand2
			counter += 4
		} else if opcode == 2 {
			operand1 := operandValue(stack, counter+1, mode1)
			operand2 := operandValue(stack, counter+2, mode2)
			destination := stack[counter+3]
			stack[destination] = operand1 * operand2
			counter += 4
		} else if opcode == 3 {
			destination := stack[counter+1]
			stack[destination] = <-input
			counter += 2
		} else if opcode == 4 {
			output <- operandValue(stack, counter+1, mode1)
			counter += 2
		} else if opcode == 5 {
			operand1 := operandValue(stack, counter+1, mode1)
			operand2 := operandValue(stack, counter+2, mode2)
			if operand1 != 0 {
				counter = operand2
			} else {
				counter += 3
			}
		} else if opcode == 6 {
			operand1 := operandValue(stack, counter+1, mode1)
			operand2 := operandValue(stack, counter+2, mode2)
			if operand1 == 0 {
				counter = operand2
			} else {
				counter += 3
			}
		} else if opcode == 7 {
			operand1 := operandValue(stack, counter+1, mode1)
			operand2 := operandValue(stack, counter+2, mode2)
			destination := stack[counter+3]
			if operand1 < operand2 {
				stack[destination] = 1
			} else {
				stack[destination] = 0
			}
			counter += 4
		} else if opcode == 8 {
			operand1 := operandValue(stack, counter+1, mode1)
			operand2 := operandValue(stack, counter+2, mode2)
			destination := stack[counter+3]
			if operand1 == operand2 {
				stack[destination] = 1
			} else {
				stack[destination] = 0
			}
			counter += 4
		} else {
			log.Fatal("Unknown opcode ", opcode)
		}
	}

}

func operandValue(stack []int, index int, mode int) int {
	result := -1
	if mode == 0 {
		result = stack[stack[index]]
	} else if mode == 1 {
		result = stack[index]
	} else {
		log.Fatal("Invalid mode ", mode)
	}
	return result
}

func readFile() []int {
	file, err := os.Open("./day7/day7.txt")
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
