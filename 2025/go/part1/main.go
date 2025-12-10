package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.Open("./test.txt")

	if err != nil {
		panic(err)
	}

	defer data.Close()
	scanner := bufio.NewScanner(data)

	starting_point := 50
	actual_password := 0
	for scanner.Scan() {

		line := scanner.Text()
		current_dir := string(line[0])
		number_ticks, err := strconv.Atoi(string(line[1:]))
		// fmt.Println(current_dir, number_ticks)
		if err != nil {
			panic(err)
		}

		if starting_point == 0 {
			actual_password = actual_password + 1
		}

		if number_ticks > 99 {
			number_ticks = number_ticks % 100
		}

		if current_dir == "L" {
			starting_point = starting_point - number_ticks
		}

		if current_dir == "R" {
			starting_point += number_ticks
		}
		if starting_point < 0 {
			starting_point = starting_point + 100
		}

		if starting_point > 99 {
			starting_point = starting_point - 100
		}

		// fmt.Println(starting_point)
	}

	if starting_point == 0 {
		actual_password = actual_password + 1
	}

	fmt.Println(actual_password)
}
