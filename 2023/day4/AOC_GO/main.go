package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	file, err := os.Open("../input_day4.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	scanner := bufio.NewScanner(file)
	var winning_numbers []int
	var my_numbers []int
	for scanner.Scan() {
		line := scanner.Text()
		winning_numbers, my_numbers = extractNumbers(line)

		num_ocurr := 0
		for i, v := range winning_numbers {
			for _, my_v := range my_numbers {
				if v == my_v {
					num_ocurr++
				}
			}
			if i == 4 {
				if num_ocurr == 1 {
					sum += 1
				} else {
					sum += int(math.Pow(2, float64(num_ocurr-1)))
				}
				fmt.Println(v, sum, num_ocurr)
			}
		}
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func extractNumbers(s string) ([]int, []int) {
	// runes := []rune(s)
	var wn []int
	var myn []int

	split_at_colon := strings.Split(s, ":")
	// fmt.Println(split_at_colon[1])

	split_at_vertical_bar := strings.Split(split_at_colon[1], "|")
	split_at_space_wn := strings.Split(strings.Trim(split_at_vertical_bar[0], " "), " ")
	split_at_space_myn := strings.Split(strings.Trim(split_at_vertical_bar[1], " "), " ")

	for _, v := range split_at_space_wn {
		curr_number, err := strconv.Atoi(v)
		if err == nil {
			wn = append(wn, curr_number)
		}
	}

	for _, v := range split_at_space_myn {
		curr_number, err := strconv.Atoi(v)
		if err == nil {
			myn = append(myn, curr_number)
		}
	}
	return wn, myn
}
