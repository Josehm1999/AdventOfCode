package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2()
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

	// orderedPagesArr := make([]map[string]int, len(pagesToUpdate)-1)

	for i := 0; i < len(pagesToUpdate)-1; i++ {
		numArr := strings.Split(pagesToUpdate[i], ",")
		orderedPages := make(map[string]int)

		currentMiddleNum := ""
		for j, numToUpdate := range numArr {
			orderedPages[numToUpdate] = j
			if (len(numArr)-1)/2 == j {
				currentMiddleNum = numToUpdate
			}
		}

		isValid := true
		for _, v := range rulesArr {
			rulesLevels := strings.Split(v, "|")
			lowerLevel := rulesLevels[0]
			upperLevel := rulesLevels[1]

			// println(lowerLevel, upperLevel)
			lowerBoundIndex, existsLower := orderedPages[lowerLevel]
			upperBoundIndex, existsUpper := orderedPages[upperLevel]

			if !existsLower || !existsUpper {
				continue
			}

			if lowerBoundIndex > upperBoundIndex {
				isValid = false
			}

		}

		if isValid {
			value, _ := strconv.Atoi(currentMiddleNum)
			sum += value
		}
	}

	println(sum)
}

func part2() {

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

	// orderedPagesArr := make([]map[string]int, len(pagesToUpdate)-1)

	for i := 0; i < len(pagesToUpdate)-1; i++ {
		numArr := strings.Split(pagesToUpdate[i], ",")
		orderedPages := make(map[string]int)

		// currentMiddleNum := ""
		for j, numToUpdate := range numArr {
			orderedPages[numToUpdate] = j
			// if (len(numArr)-1)/2 == j {
			// 	currentMiddleNum = numToUpdate
			// }
		}

		isValid := true
		for _, v := range rulesArr {
			rulesLevels := strings.Split(v, "|")
			lowerLevel := rulesLevels[0]
			upperLevel := rulesLevels[1]

			// println(lowerLevel, upperLevel)
			lowerBoundIndex, existsLower := orderedPages[lowerLevel]
			upperBoundIndex, existsUpper := orderedPages[upperLevel]

			if !existsLower || !existsUpper {
				continue
			}

			if lowerBoundIndex > upperBoundIndex {
				isValid = false
			}

		}

		// if isValid {
		// 	value, _ := strconv.Atoi(currentMiddleNum)
		// 	sum += value
		// }

		if !isValid {

			// println("New Line")
			for _, v := range rulesArr {
				rulesLevels := strings.Split(v, "|")
				lowerLevel := rulesLevels[0]
				upperLevel := rulesLevels[1]

				lowerBoundIndex, existsLower := orderedPages[lowerLevel]
				upperBoundIndex, existsUpper := orderedPages[upperLevel]

				if !existsLower || !existsUpper {
					continue
				}

				if lowerBoundIndex > upperBoundIndex {
					orderedPages[lowerLevel] = upperBoundIndex
					orderedPages[upperLevel] = lowerBoundIndex
					numArr[upperBoundIndex] = lowerLevel
					numArr[lowerBoundIndex] = upperLevel
				}

				// println(lowerBoundIndex, lowerLevel)
				// println(upperBoundIndex, upperLevel)
				//
				// println("------")
			}

			test := numArr[(len(numArr)-1)/2]

			// for _, v := range numArr {
			// 	println(v)
			// }
			// println(test)
			value, _ := strconv.Atoi(test)
			// println(test)
			sum += value
		}
	}

	println(sum)
}
