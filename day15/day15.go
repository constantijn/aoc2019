package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x, y int
}

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

	grid := [41][41]string{}
	x, y, fixX, fixY := 21, 21, -1, -1
	grid[x][y] = "."

	const NORTH, SOUTH, WEST, EAST = 1, 2, 3, 4

	i := 0

	for true {
		i++
		direction := rand.Intn(4) + 1
		//fmt.Println("direction", direction)
		input <- direction
		result := <-output
		//fmt.Println("output", x , y, result)

		if result == 2 {
			fixX = x
			fixY = y
		}

		switch direction {
		case NORTH:
			{
				if result >= 1 {
					y--
					grid[x][y] = "."
				} else {
					grid[x][y-1] = "█"
				}
			}
		case SOUTH:
			{
				if result >= 1 {
					y++
					grid[x][y] = "."
				} else {
					grid[x][y+1] = "█"
				}
			}
		case WEST:
			{
				if result >= 1 {
					x--
					grid[x][y] = "."
				} else {
					grid[x-1][y] = "█"
				}
			}
		case EAST:
			{
				if result >= 1 {
					x++
					grid[x][y] = "."
				} else {
					grid[x+1][y] = "█"
				}
			}
		}

		if i > 999999 {
			print(grid, fixX, fixY)
			fmt.Println("15A done (manually count path)", time.Since(start))
			break
		}
	}

	start = time.Now()
	minutes := 0
	newPoints := []point{point{fixX, fixY}}
	for true {
		newPointSet := map[point]bool{}
		for _, spot := range newPoints {
			if grid[spot.x+1][spot.y] == "." {
				newPointSet[point{spot.x + 1, spot.y}] = true
			}
			if grid[spot.x-1][spot.y] == "." {
				newPointSet[point{spot.x - 1, spot.y}] = true
			}
			if grid[spot.x][spot.y+1] == "." {
				newPointSet[point{spot.x, spot.y + 1}] = true
			}
			if grid[spot.x][spot.y-1] == "." {
				newPointSet[point{spot.x, spot.y - 1}] = true
			}
		}
		newPoints = []point{}
		for newPoint := range newPointSet {
			newPoints = append(newPoints, newPoint)
			grid[newPoint.x][newPoint.y] = "O"
		}
		minutes++

		if len(newPoints) == 0 {
			print(grid, fixX, fixY)
			fmt.Println("Day 15B", minutes, time.Since(start))
			break
		}

	}

}

func print(grid [41][41]string, x int, y int) {

	for i := 0; i < 41; i++ {
		for j := 0; j < 41; j++ {
			if j == x && i == y {
				fmt.Print("D")
				continue
			}
			if j == 21 && i == 21 {
				fmt.Print("S")
				continue
			}
			switch grid[j][i] {
			case "":
				fmt.Print(" ")
			default:
				fmt.Print(grid[j][i])
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
	file, err := os.Open("./day15/day15.txt")
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
