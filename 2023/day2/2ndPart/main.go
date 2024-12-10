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
	file, err := os.Open("../input_day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		cube_info_by_line := extractsNumOfCubesByColor(line)
		curr_red_cube := 0
		curr_green_cube := 0
		curr_blue_cube := 0

		for _, value := range cube_info_by_line {

			switch value.color {
			case "red":
				if curr_red_cube <= value.number {
					curr_red_cube = value.number
				}
			case "green":
				if curr_green_cube <= value.number {
					curr_green_cube = value.number
				}
			case "blue":
				if curr_blue_cube <= value.number {
					curr_blue_cube = value.number
				}
			}
		}
		sum += curr_red_cube * curr_blue_cube * curr_green_cube
	}

	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func extractIds(s string) int {
	position_of_separator := strings.Index(s, ":")
	current_id, err := strconv.Atoi(s[5:position_of_separator])
	if err != nil {
		fmt.Println("Cannot extract Id")
	}
	return current_id
}

type cube_data struct {
	index  int
	color  string
	number int
}

func extractsNumOfCubesByColor(s string) []cube_data {
	position_of_separator := strings.Index(s, ":")
	right_side_with_cubes := string(s[position_of_separator+1:])
	var cubes_by_line []cube_data
	splitted_by_dotted_comma := strings.Split(right_side_with_cubes, ";")
	for index, value_dotted_coma := range splitted_by_dotted_comma {
		for _, value_comma := range strings.Split(value_dotted_coma, ",") {

			trimmed_value := strings.Trim(value_comma, " ")
			number_value_cube, _ := strconv.Atoi(strings.Split(trimmed_value, " ")[0])
			cube_info := cube_data{
				index:  index,
				color:  strings.Split(trimmed_value, " ")[1],
				number: number_value_cube,
			}

			cubes_by_line = append(cubes_by_line, cube_info)
		}
	}

	return cubes_by_line
}
