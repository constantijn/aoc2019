package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	start1 := time.Now()
	possibilities1 := 0
	for i:= 248345; i< 746315; i++ {
		if potentialPassword1(i) {
			possibilities1++
		}
	}
	fmt.Println("Day 4A", possibilities1, time.Since(start1))

	//fmt.Println("112233 T", potentialPassword2(112233))
	//fmt.Println("123444 F", potentialPassword2(123444))
	//fmt.Println("111122 T", potentialPassword2(111122))
	//fmt.Println("123455 T", potentialPassword2(123455))
	//fmt.Println("112345 T", potentialPassword2(112345))
	//fmt.Println("123334 F", potentialPassword2(123334))
	//fmt.Println("144445 F", potentialPassword2(144445))

	start2 := time.Now()
	possibilities2 := 0
	for i:= 248345; i< 746315; i++ {
		if potentialPassword2(i) {
			possibilities2++
		}
	}
	fmt.Println("Day 4B", possibilities2, time.Since(start2))


}

func potentialPassword1(code int) bool {
	codeString := strconv.Itoa(code)

	doubleFound := false

	for i := 1; i< len(codeString) ; i++ {
		if codeString[i] < codeString[i-1] {
			return false
		}
		if codeString[i] == codeString[i-1] {
			doubleFound = true
		}
	}
	return doubleFound
}

func potentialPassword2(code int) bool {
	codeString := strconv.Itoa(code)

	doubleFound := false

	lastChar := codeString[0]
	lastCharCount := 1

	for i := 1; i < len(codeString) ; i++ {
		if codeString[i] < codeString[i-1] {
			return false
		}

		currentChar := codeString[i]

		if currentChar == lastChar {
			lastCharCount ++
			if i == len(codeString) -1 && lastCharCount == 2 {
				doubleFound = true
			}
		} else if lastCharCount == 2 {
			doubleFound = true
		} else {
			lastCharCount = 1
		}

		lastChar = currentChar

	}
	return doubleFound



}