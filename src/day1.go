package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//println(caculateFuel1(12))
	//println(caculateFuel1(14))
	//println(caculateFuel1(1969))
	//println(caculateFuel1(100756))

	result := 0
	lines :=  getLines()
	for _, line := range lines {
		result += caculateFuel1(line)
	}
	fmt.Println("1A:", result)

	//println(caculateFuel2(14))
	//println(caculateFuel2(1969))
	//println(caculateFuel2(100756))

	result = 0
	for _, line := range lines {
		result += caculateFuel2(line)
	}
	fmt.Println("1B:", result)

}

func getLines() []int {
	file, err := os.Open("./input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		n, err := strconv.ParseInt(scanner.Text(), 10, 32)

		if err != nil {
			log.Fatal("Can't parse to int: " + scanner.Text() )
		}

		result = append(result, int(n))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result

}

func caculateFuel1(mass int) int {
	return (mass / 3) - 2
}

func caculateFuel2(mass int) int {
	result := 0
	fuel := caculateFuel1(mass)
	if fuel > 0 {
		result += fuel + caculateFuel2(fuel)
	}
	return result
}
