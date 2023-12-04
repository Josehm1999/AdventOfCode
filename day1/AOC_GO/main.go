package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("../input_day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		left, right := extractNumbers(line)

		sum += left*10 + right
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func extractNumbers(s string) (int, int) {
	runes := []rune(s)
	left_number := -1
	right_number := -1

	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			if left_number == -1 {
				left_number, _ = strconv.Atoi(string(runes[i]))
				right_number, _ = strconv.Atoi(string(runes[i]))
			} else {
				right_number, _ = strconv.Atoi(string(runes[i]))
			}
		}
	}
	fmt.Println(left_number, right_number)
	return left_number, right_number
}
