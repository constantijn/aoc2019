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

type moon struct {
	posX, posY, posZ int
	velX, velY, velZ int
}

func main() {
	lines := readFile()
	start := time.Now()
	moons := parseMoons(lines)
	origMoons := parseMoons(lines)

	loopX, loopY, loopZ, i := 0, 0, 0, 0

	for true {
		i++
		applyGravityAll(moons)
		applyVelocity(moons)
		if loopX == 0 &&
			moons[0].velX == 0 && moons[1].velX == 0 && moons[2].velX == 0 && moons[3].velX == 0 &&
			moons[0].posX == origMoons[0].posX && moons[1].posX == origMoons[1].posX &&
			moons[2].posX == origMoons[2].posX && moons[3].posX == origMoons[3].posX {
			loopX = i
		}
		if loopY == 0 &&
			moons[0].velY == 0 && moons[1].velY == 0 && moons[2].velY == 0 && moons[3].velY == 0 &&
			moons[0].posY == origMoons[0].posY && moons[1].posY == origMoons[1].posY &&
			moons[2].posY == origMoons[2].posY && moons[3].posY == origMoons[3].posY {
			loopY = i
		}
		if loopZ == 0 &&
			moons[0].velZ == 0 && moons[1].velZ == 0 && moons[2].velZ == 0 && moons[3].velZ == 0 &&
			moons[0].posZ == origMoons[0].posZ && moons[1].posZ == origMoons[1].posZ &&
			moons[2].posZ == origMoons[2].posZ && moons[3].posZ == origMoons[3].posZ {
			loopZ = i
		}
		if i == 1000 {
			fmt.Println("Day12A", calculateEnergy(moons), time.Since(start))
		}
		if loopX != 0 && loopY != 0 && loopZ != 0 {
			fmt.Println("Day12B", lcm(loopX, loopY, loopZ), time.Since(start))
			break
		}
	}

}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func calculateEnergy(moons []moon) int {
	result := 0
	for _, moon := range moons {
		result += calculateEnergySingle(moon)
	}
	return result
}

func calculateEnergySingle(moon moon) int {
	positionEnergy := abs(moon.posX) + abs(moon.posY) + abs(moon.posZ)
	velocityEnergy := abs(moon.velX) + abs(moon.velY) + abs(moon.velZ)
	return positionEnergy * velocityEnergy
}

func abs(input int) int {
	if input < 0 {
		return -input
	}
	return input
}

func printMoons(moons []moon) {
	for _, moon := range moons {
		fmt.Printf("pos=<x=%3d, y=%3d, z=%3d>, vel=<x=%3d, y=%3d, z=%3d>\n", moon.posX, moon.posY, moon.posZ, moon.velX, moon.velY, moon.velZ)
	}
}

func applyVelocity(moons []moon) {
	for i := range moons {
		moons[i].posX += moons[i].velX
		moons[i].posY += moons[i].velY
		moons[i].posZ += moons[i].velZ
	}
}

func applyGravityAll(moons []moon) {
	if len(moons) != 4 {
		log.Fatal("Expected 4 moons")
	}
	applyGravityPair(&moons[0], &moons[1])
	applyGravityPair(&moons[0], &moons[2])
	applyGravityPair(&moons[0], &moons[3])
	applyGravityPair(&moons[1], &moons[2])
	applyGravityPair(&moons[1], &moons[3])
	applyGravityPair(&moons[2], &moons[3])
}

func applyGravityPair(moon1 *moon, moon2 *moon) {
	if moon1.posX > moon2.posX {
		moon1.velX--
		moon2.velX++
	}
	if moon1.posX < moon2.posX {
		moon1.velX++
		moon2.velX--
	}
	if moon1.posY > moon2.posY {
		moon1.velY--
		moon2.velY++
	}
	if moon1.posY < moon2.posY {
		moon1.velY++
		moon2.velY--
	}
	if moon1.posZ > moon2.posZ {
		moon1.velZ--
		moon2.velZ++

	}
	if moon1.posZ < moon2.posZ {
		moon1.velZ++
		moon2.velZ--
	}
}

func parseMoons(input []string) []moon {
	var moons []moon

	for _, line := range input {
		moonData := map[string]int{}
		for _, tuple := range strings.Split(strings.ReplaceAll(line, ">", ""), ",") {
			parts := strings.Split(tuple, "=")
			intValue, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal("can't parse int", parts[1])
			}
			moonData[strings.TrimSpace(parts[0])] = intValue
		}
		moons = append(moons, moon{posX: moonData["<x"], posY: moonData["y"], posZ: moonData["z"]})
	}

	return moons
}

func readFile() []string {
	file, err := os.Open("./day12/day12.txt")
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
