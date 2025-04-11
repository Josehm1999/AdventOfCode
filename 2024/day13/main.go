package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	// part2()
}

type Button struct {
	row int
	col int
}

func part1() {

	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	temp := strings.Split(file, "\n\n")

	for _, crawl_machine := range temp {
		// fmt.Println(crawl_machine)

		individual_lines := strings.Split(crawl_machine, "\n")
		// fmt.Println(individual_lines)

		button_a_string := individual_lines[0]
		button_b_string := individual_lines[1]
		prize_string := individual_lines[2]

		var button_a_arr [2]int
		var button_b_arr [2]int
		var prize_arr [2]int
		for i, v := range strings.Split(button_a_string, ",") {
			plus_symbol_idx := strings.Index(v, "+")
			a_value, _ := strconv.Atoi(v[plus_symbol_idx+1:])
			button_a_arr[i] = a_value
		}
		for i, v := range strings.Split(button_b_string, ",") {
			plus_symbol_idx := strings.Index(v, "+")
			b_value, _ := strconv.Atoi(v[plus_symbol_idx+1:])
			button_b_arr[i] = b_value
		}
		for i, v := range strings.Split(prize_string, ",") {
			equal_symbol_idx := strings.Index(v, "=")
			prize_value, _ := strconv.Atoi(v[equal_symbol_idx+1:])
			prize_arr[i] = prize_value
		}

		// fmt.Println(button_a_arr[0]%button_b_arr[0], button_b_arr[0])
		fmt.Println(button_a_arr, button_b_arr, prize_arr)
	}
}

func part2() {
	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	temp := strings.Split(file, "\n")
	fmt.Println(temp)
}
