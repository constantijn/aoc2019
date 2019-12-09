package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	imgDataRaw := readFile()
	start := time.Now()

	width := 25
	height := 6
	layers := 100

	imgData := [100][6][25]int{}

	for i := 0; i < len(imgDataRaw); i++ {
		layer := i / (width * height)
		y := i % (width * height) / width
		x := i % (width * height) % width
		imgData[layer][y][x] = int(imgDataRaw[i]) - 48
	}

	min0 := 100
	result := -1

	for _, layer := range imgData {
		count0 := 0
		count1 := 0
		count2 := 0
		for _, row := range layer {
			for _, pixel := range row {
				if pixel == 0 {
					count0++
				} else if pixel == 1 {
					count1++
				} else if pixel == 2 {
					count2++
				} else {
					log.Fatal("invalid pixel", pixel)
				}
			}
		}

		if count0 < min0 {
			min0 = count0
			result = count1 * count2
		}
	}

	fmt.Println("Day 8A", result, time.Since(start))
	start = time.Now()

	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			for layer := 0; layer < layers; layer++ {
				if imgData[layer][x][y] == 0 {
					fmt.Print("░")
					break
				} else if imgData[layer][x][y] == 1 {
					fmt.Print("█")
					break
				}
			}
		}
		fmt.Println("")
	}

	fmt.Println("Day 8B (see above)", time.Since(start))

}

func readFile() string {
	file, err := os.Open("./day8/day8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
