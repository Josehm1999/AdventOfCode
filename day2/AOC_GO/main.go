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
	file, err := os.Open("../input_day2.txt")
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
	matchnum := -1
	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			if left_number == -1 {
				left_number, _ = strconv.Atoi(string(runes[i]))
				right_number, _ = strconv.Atoi(string(runes[i]))
			} else {
				right_number, _ = strconv.Atoi(string(runes[i]))
			}
		} else {

			if len(runes) >= 3 && i <= len(runes)-3  {
				matchnum = matchNumbers(string(runes[i : i+3]))
			}
			if len(runes) >= 4 && i <= len(runes)-4 && matchnum == -1 {
				matchnum = matchNumbers(string(runes[i : i+4]))
			}

			if len(runes) >= 5 && i <= len(runes)-5 && matchnum == -1 {
				matchnum = matchNumbers(string(runes[i : i+5]))
			}

			if matchnum != -1 {
				if left_number == -1 {
					left_number = matchnum
					right_number = matchnum
				} else {
					right_number = matchnum
				}
				matchnum = -1
			}
		}

	}
	fmt.Println(left_number, right_number)
	return left_number, right_number
}

func matchNumbers(s string) int {
	m := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	for k, v := range m {
		if k == s {
			return v
		}
	}
	return -1
}
func extractCharacters(s string) int {
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		return matchNumbers(string(runes[i:]))
	}
	return -1
}
