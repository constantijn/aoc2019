package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x int
	y int
	distance int
}

func main() {
	start := time.Now()
	wires := readFile()
	intersections := intersect(getPoints(wires[0]), getPoints(wires[1]))

	closest := math.MaxInt32
	for _, intersection := range intersections {
		manhattan := abs(intersection.x) + abs(intersection.y)
		closest = min(closest, manhattan)
	}
	fmt.Println("Day 3A:", closest)

	shortest := math.MaxInt32
	for _, intersection := range intersections {
		shortest = min(shortest, intersection.distance )
	}
	fmt.Println("Day 3B:", shortest)
	fmt.Println(time.Since(start))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func intersect(left []point, right []point) []point {
	var result []point

	for _, a := range left {
		for _, b := range right {
			if a.x == b.x && a.y == b.y {
				result = append(result, point {x:a.x, y:a.y, distance:a.distance + b.distance})
			}
		}
	}

	return result
}



func getPoints(wire []string) []point {
	x := 0
	y := 0
	distance := 0

	var points []point

	for _, instruction := range wire {
		direction := string(instruction[0])
		amount, err := strconv.Atoi(string(instruction[1:]))
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < amount; i++ {
			distance ++
			switch direction {
			case "R":
				x++
			case "L":
				x--
			case "U":
				y++
			case "D":
				y--
			default:
				log.Fatal("unknown direction " + direction)
			}
			points = append(points, point{x:x, y:y, distance:distance})
		}
		//println(instruction, direction, amount)
	}
	return points
}

func readFile() [][]string {
	file, err := os.Open("./day3/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, strings.Split(scanner.Text(), ","))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
