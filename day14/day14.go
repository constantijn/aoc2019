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

type reaction struct {
	output       string
	outputAmount int
	inputs       map[string]int
}

func main() {
	lines := readFile()
	start := time.Now()
	reactions := parseReactions(lines)

	ore := calculate(reactions, map[string]int{"FUEL": 1})
	fmt.Println("Day 14A", ore, time.Since(start))

	start = time.Now()
	cargoSize := 1000000000000
	guess := cargoSize / ore
	increment := 1000000
	for true {
		ore := calculate(reactions, map[string]int{"FUEL": guess})
		if ore > cargoSize {
			if increment > 1 {
				guess -= increment
				increment = increment / 10
			} else {
				fmt.Println("Day 14B", guess-1, time.Since(start))
				break
			}

		}
		guess += increment
	}

}

func calculate(reactions map[string]reaction, requirements map[string]int) int {
	leftovers := map[string]int{}

	for true {
		for chemical, amount := range requirements {
			if chemical == "ORE" {
				continue
			}
			reaction := reactions[chemical]
			reactionTimes := amount / reaction.outputAmount
			if amount%reaction.outputAmount > 0 {
				reactionTimes++
				leftovers[chemical] += reaction.outputAmount*reactionTimes - amount
			}
			for inputChemical, inputAmount := range reaction.inputs {
				requiredAmount := inputAmount * reactionTimes
				if requiredAmount > leftovers[inputChemical] {
					requirements[inputChemical] += requiredAmount - leftovers[inputChemical]
					delete(leftovers, inputChemical)
				} else {
					leftovers[inputChemical] = leftovers[inputChemical] - requiredAmount
				}
			}
			delete(requirements, chemical)
		}

		if len(requirements) == 1 && requirements["ORE"] > 0 {
			break
		}
	}
	return requirements["ORE"]
}

func parseReactions(lines []string) map[string]reaction {
	reactions := map[string]reaction{}

	for _, line := range lines {
		lineParts := strings.Split(line, " => ")
		inputStrings := strings.Split(lineParts[0], ", ")
		outputParts := strings.Split(lineParts[1], " ")
		inputs := map[string]int{}
		for _, inputString := range inputStrings {
			inputParts := strings.Split(inputString, " ")
			inputs[inputParts[1]] = parseInt(inputParts[0])
		}

		reactions[outputParts[1]] = reaction{output: outputParts[1], outputAmount: parseInt(outputParts[0]), inputs: inputs}

	}
	return reactions
}

func parseInt(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("invalid int string [", input, "]")
	}
	return result
}

func readFile() []string {
	file, err := os.Open("./day14/day14.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
