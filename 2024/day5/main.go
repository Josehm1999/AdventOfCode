package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	part1()
}

func part1() {

	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	// line := 0

	temp := strings.Split(file, "\n\n")

	rulesArr := strings.Split(temp[0], "\n")
	pagesToUpdate := strings.Split(temp[1], "\n")

	sum := 0
	// rules := make(map[int]int)
	// orderedPages := make([]int, )
	for i := 0; i < len(pagesToUpdate)-1; i++ {
		fmt.Println(pagesToUpdate[i])
		sum++
	}
	println(rulesArr[20])
	println(pagesToUpdate[5])
	println(sum)
}

func part2() {

	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	// line := 0

	temp := strings.Split(file, "\n")

	sum := 0

	for i := 0; i < len(temp)-1; i++ {
		fmt.Println(temp[i])

		isPossibleDown := false

		if (len(temp)-2)-i >= 2 {
			isPossibleDown = true
		}
		letters := strings.Split(temp[i], "")

		j := 0

		for j < len(letters) {
			if letters[j] != "M" && letters[j] != "S" {
				j++
				continue
			}
			if isPossibleDown {
				nextCharArr1 := strings.Split(temp[i+1], "")
				nextCharArr2 := strings.Split(temp[i+2], "")

				firstAxis := false
				// check for diagonal down - right
				if (len(letters)-1)-j >= 2 {

					possibleXmasDownDiagRight := letters[j] + nextCharArr1[j+1] + nextCharArr2[j+2]

					if possibleXmasDownDiagRight == "MAS" || possibleXmasDownDiagRight == "SAM" {

						firstAxis = true
					}
				}

				if !firstAxis {
					j++
					continue
				}

				// check for diagonal down - left
				if j+2 >= 2 {

					possibleXmasDownDiagLeft := letters[j+2] + nextCharArr1[j+1] + nextCharArr2[j]
					if possibleXmasDownDiagLeft == "MAS" || possibleXmasDownDiagLeft == "SAM" {

						sum++
					}
				}
			}
			j++
		}

	}

	println(sum)
}
