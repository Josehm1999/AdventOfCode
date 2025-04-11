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

	total_tokens := 0
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
			prize_arr[i] = prize_value + 10000000000000
		}

		is_possible := true
		for i := range button_a_arr {
			// gcm := find_gcm_two_ints(button_a_arr[i], button_b_arr[i])
			max_possible_value := button_a_arr[i]*100 + button_b_arr[i]*100

			if max_possible_value < prize_arr[i] {
				is_possible = false
			}
		}

		if is_possible {

			fmt.Println(button_a_arr, button_b_arr, prize_arr)
			inversa := button_a_arr[0]*button_b_arr[1] - (button_a_arr[1] * button_b_arr[0])

			a_button_times_pressed := ((button_b_arr[1] * prize_arr[0]) - (button_b_arr[0] * prize_arr[1])) / inversa
			b_button_times_pressed := ((button_a_arr[0] * prize_arr[1]) - (button_a_arr[1] * prize_arr[0])) / inversa

			if a_button_times_pressed > 0 && b_button_times_pressed > 0 {

				// fmt.Println(((button_b_arr[1] * prize_arr[0]) - (button_b_arr[0] * prize_arr[1])), ((button_a_arr[0] * prize_arr[1]) - (button_a_arr[1] * prize_arr[0])), inversa)
				// fmt.Println(a_button_times_pressed, b_button_times_pressed)

				// fmt.Println(a_button_times_pressed*button_a_arr[0]+button_b_arr[0]*b_button_times_pressed, prize_arr[0])
				// fmt.Println(a_button_times_pressed*button_a_arr[1]+button_b_arr[1]*b_button_times_pressed, prize_arr[1])
				if a_button_times_pressed*button_a_arr[0]+button_b_arr[0]*b_button_times_pressed == prize_arr[0] && a_button_times_pressed*button_a_arr[1]+button_b_arr[1]*b_button_times_pressed == prize_arr[1] {

					total_tmp := a_button_times_pressed*3 + b_button_times_pressed

					total_tokens += total_tmp
				}

			}
		}
	}
	fmt.Println(total_tokens)
}

// func find_gcm_two_ints(x int, y int) int {
// 	for y > 0 {
// 		tmp := y
// 		y = x % y
// 		x = tmp
// 	}
//
// 	return x
// }

func part2() {
	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	temp := strings.Split(file, "\n")
	fmt.Println(temp)
}
