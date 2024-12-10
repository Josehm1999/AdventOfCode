package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)

	isDoConditionActive := true
	for scanner.Scan() {
		line := scanner.Text()
		letters := strings.Split(line, "")
		fmt.Println(letters)

		for i, _ := range letters {
			if i >= len(letters)-2 {
				break
			}

			possibleDoConditional := letters[i] + letters[i+1]

			if possibleDoConditional == "do" {

				if strings.Join(letters[i:i+4], "") == "do()" {
					isDoConditionActive = true
					println(strings.Join(letters[i:i+4], ""))
				}

				if strings.Join(letters[i:i+7], "") == "don't()" {
					isDoConditionActive = false
					println(strings.Join(letters[i:i+7], ""))
				}

			}

			if isDoConditionActive {

				possibleMulMatch := letters[i] + letters[i+1] + letters[i+2]

				if possibleMulMatch != "mul" {
					continue
				}

				possibleLeftParen := letters[i+3]

				if possibleLeftParen != "(" {
					continue
				}

				possibleCommaPosition := 0
				for j := i + 4; j < i+8; j++ {
					if letters[j] == "," {
						possibleCommaPosition = j
						break
					}
				}

				if possibleCommaPosition == 0 {
					continue
				}

				possibleLeftNumber, error := strconv.Atoi(strings.Join(letters[i+4:possibleCommaPosition], ""))

				if error != nil {
					continue
				}

				possibleRightParenPosition := 0
				for k := possibleCommaPosition + 1; k < possibleCommaPosition+5; k++ {

					if letters[k] == ")" {
						possibleRightParenPosition = k
						break
					}
				}

				// fmt.Println("Is mul", letters[i])
				if possibleRightParenPosition == 0 {
					continue
				}

				possibleRightNumber, error := strconv.Atoi(strings.Join(letters[possibleCommaPosition+1:possibleRightParenPosition], ""))

				if error != nil {
					continue
				}

				// fmt.Println(possibleLeftNumber * possibleRightNumber)
				sum += possibleLeftNumber * possibleRightNumber
				fmt.Println(possibleMulMatch + possibleLeftParen + strings.Join(letters[i+4:possibleCommaPosition], "") + letters[possibleCommaPosition] + strings.Join(letters[possibleCommaPosition+1:possibleRightParenPosition], ""))
			}
		}
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
