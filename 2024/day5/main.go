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

	temp := strings.Split(file, "\n")

	sum := 0

	for i := 0; i < len(temp)-1; i++ {
		fmt.Println(temp[i])

		isPossibleUp := false
		isPossibleDown := false

		if (len(temp)-2)-i > 2 {
			isPossibleDown = true
		}
		if i > 2 {
			isPossibleUp = true
		}

		letters := strings.Split(temp[i], "")
		for j, v := range letters {
			if v != "X" {
				continue
			}

			// check for right
			if (len(letters)-1)-j >= 3 {
				// println((len(letters) - 1) - i)
				if strings.Join(letters[j:j+4], "") == "XMAS" {
					println(strings.Join(letters[j:j+4], ""), "RIGHT")
					sum++
				}
			}

			// check for left
			if j >= 3 {
				if strings.Join(letters[j-3:j+1], "") == "SAMX" {
					println(strings.Join(letters[j-3:j+1], ""), "LEFT")
					sum++
				}
			}

			// check for down

			if isPossibleDown {

				nextCharArr1 := strings.Split(temp[i+1], "")
				nextCharArr2 := strings.Split(temp[i+2], "")
				nextCharArr3 := strings.Split(temp[i+3], "")

				possibleXmasDown := letters[j] + nextCharArr1[j] + nextCharArr2[j] + nextCharArr3[j]

				if possibleXmasDown == "XMAS" {
					println(possibleXmasDown, "DOWN")
					sum++
				}
				// check for diagonal down - left
				if j >= 3 {

					possibleXmasDownDiagLeft := letters[j] + nextCharArr1[j-1] + nextCharArr2[j-2] + nextCharArr3[j-3]
					if possibleXmasDownDiagLeft == "XMAS" {

						println(possibleXmasDownDiagLeft, "DDL")
						sum++
					}
				}
				// check for diagonal down - right
				if (len(letters)-1)-j >= 3 {

					possibleXmasDownDiagRight := letters[j] + nextCharArr1[j+1] + nextCharArr2[j+2] + nextCharArr3[j+3]

					if possibleXmasDownDiagRight == "XMAS" {
						println(possibleXmasDownDiagRight, "DDR")
						sum++
					}
				}
			}

			// check for up - Up
			if isPossibleUp {
				prevCharArr1 := strings.Split(temp[i-1], "")
				prevCharArr2 := strings.Split(temp[i-2], "")
				prevCharArr3 := strings.Split(temp[i-3], "")

				possibleXmasUp := letters[j] + prevCharArr1[j] + prevCharArr2[j] + prevCharArr3[j]

				if possibleXmasUp == "XMAS" {
					println(possibleXmasUp, "UP")
					sum++
				}

				// check for diagonal up - left
				if j >= 3 {

					possibleXmasUpDiagLeft := letters[j] + prevCharArr1[j-1] + prevCharArr2[j-2] + prevCharArr3[j-3]

					if possibleXmasUpDiagLeft == "XMAS" {
						println(possibleXmasUpDiagLeft, "DUL")
						sum++
					}
				}

				// check for diagonal up -  right
				if (len(letters)-1)-j >= 3 {

					possibleXmasUpDiagRight := letters[j] + prevCharArr1[j+1] + prevCharArr2[j+2] + prevCharArr3[j+3]

					if possibleXmasUpDiagRight == "XMAS" {
						println(possibleXmasUpDiagRight, "DUR")
						sum++
					}
				}

			}

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
