package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	lines := readFile()

	start1 := time.Now()
	orbits := map[string]string{}

	for _, line := range lines {
		orbits[line[1]] = line[0]
	}

	count := 0
	for key, _ := range orbits {
		for true {
			value := orbits[key]
			count++
			if value == "COM" {
				break
			} else {
				key = value
			}
		}
	}

	fmt.Println("Day 6A", count, time.Since(start1))

	start2 := time.Now()
	youPath := pathFromCom(orbits, "YOU")
	sanPath := pathFromCom(orbits, "SAN")

	for i := 0; i < len(youPath); i++ {
		if youPath[i] != sanPath[i] {
			result := (len(youPath) + len(sanPath)) - 2*i
			fmt.Println("Day 6B", result, time.Since(start2))
			break
		}
	}

}

func pathFromCom(orbits map[string]string, start string) []string {
	var path []string

	key := start
	for true {
		value := orbits[key]
		path = append(path, value)
		if value == "COM" {
			break
		} else {
			key = value
		}
	}

	for i := len(path)/2 - 1; i >= 0; i-- {
		opp := len(path) - 1 - i
		path[i], path[opp] = path[opp], path[i]
	}

	return path

}

func readFile() [][]string {
	file, err := os.Open("./day6/day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, strings.Split(scanner.Text(), ")"))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
