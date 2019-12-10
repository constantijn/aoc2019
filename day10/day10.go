package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type point struct {
	X, Y     int
	Distance float64
}

func main() {
	grid := readFile()
	start := time.Now()

	lenX := len(grid[0])
	lenY := len(grid)
	max := 0
	stationDirections := map[float64][]point{}

	for x := 0; x < lenX; x++ {
		for y := 0; y < lenY; y++ {
			if grid[y][x] != "." {
				directions := map[float64][]point{}
				for x2 := 0; x2 < lenX; x2++ {
					for y2 := 0; y2 < lenY; y2++ {
						if grid[y2][x2] != "." && !(x == x2 && y == y2) {
							direction := (math.Atan2(float64(y2-y), float64(x2-x)) / math.Pi) + 0.5
							if direction < 0 {
								direction += 2
							}
							distance := math.Sqrt(math.Pow(float64(x2-x), 2) + math.Pow(float64(y2-y), 2))
							pointArray, present := directions[direction]
							if present {
								directions[direction] = append(pointArray, point{x2, y2, distance})
							} else {
								directions[direction] = []point{point{x2, y2, distance}}
							}
						}
					}
				}
				if len(directions) > max {
					max = len(directions)
					stationDirections = directions
				}
				grid[y][x] = strconv.Itoa(len(directions))
			}
		}
	}

	fmt.Println("Dat 10A", max, time.Since(start))
	start = time.Now()

	var keys []float64

	for key, value := range stationDirections {
		sort.Slice(value, func(i, j int) bool { return value[i].Distance < value[j].Distance })
		keys = append(keys, key)
	}
	sort.Float64s(keys)

	count := 0
	for rotations := 0; true; rotations++ {
		for i := 0; i < len(keys); i++ {
			points := stationDirections[keys[i]]
			if rotations < len(points) {
				count++
			}
			if count == 200 {
				fmt.Println("Day 10B", points[rotations].X*100+points[rotations].Y, time.Since(start))
				os.Exit(0)
			}
		}
	}

}

func readFile() [][]string {
	file, err := os.Open("./day10/day10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, strings.Split(scanner.Text(), ""))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
