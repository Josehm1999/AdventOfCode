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

type convertion_ranges struct {
	destination_range int
	source_range      int
	range_length      int
}

func main() {
	file, err := os.Open("../input_day5.txt")
	if err != nil {
		log.Fatal(err)
	}

	// sum := 0
	// counter := 0
	scanner := bufio.NewScanner(file)
	var arr_convertion_rg []convertion_ranges
	var maze [][]convertion_ranges

	counter := 0
	var arr_strings []string
	var arr_seeds []int
	for scanner.Scan() {
		line := scanner.Text()
		if counter == 0 {
			arr_seeds = extractSeeds(line)
			counter++
		}

		// if len(strings.Split(line, ":")) > 1 {
		// 	map_name = strings.Split(strings.Split(line, ":")[0], " ")[0]
		// }
		arr_strings = append(arr_strings, line)
	}

	for i, line := range arr_strings {
		_, err := strconv.Atoi(strings.Split(line, " ")[0])
		if err == nil {
			arr_convertion_rg = append(arr_convertion_rg, extractConvertionRanges(line))
		}

		if strings.Trim(line, " ") == "" || len(arr_strings)-1 == i {
			if len(arr_convertion_rg) >= 1 {
				maze = append(maze, arr_convertion_rg)
			}
			arr_convertion_rg = nil
		}
	}

	lowest := 0
	for _, num_seed := range arr_seeds {
		fmt.Println("Change Number of Seed", num_seed)
		for _, value := range maze {
			// fmt.Println("Change Number of Maps", value, num_seed)
			for _, map_values := range value {
				dest := map_values.destination_range
				src := map_values.source_range
				length := map_values.range_length
				fmt.Println(num_seed, dest, src, length)

				if (num_seed >= src) && (num_seed <= src+length-1) {
					num_seed += (dest - src)
				}
			}
			if num_seed < lowest {
				lowest = num_seed
			}
		}

	}

	fmt.Println(lowest)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func extractConvertionRanges(s string) convertion_ranges {
	split_at_space := strings.Split(s, " ")
	start_range := 0
	destination_range := 0
	length := 0
	for i, v := range split_at_space {
		num, err := strconv.Atoi(v)
		if err == nil {
			if i == 0 {
				destination_range = num
			}

			if i == 1 {
				start_range = num
			}

			if i == 2 {
				length = num
			}
		}
	}

	return convertion_ranges{source_range: start_range, destination_range: destination_range, range_length: length}
}
func extractSeeds(s string) []int {
	split_at_colon := strings.Split(s, ":")
	numbers := strings.Trim(split_at_colon[1], " ")
	arr_num := strings.Split(numbers, " ")
	var seeds []int
	for _, v := range arr_num {
		if num, err := strconv.Atoi(string(v)); err == nil {
			seeds = append(seeds, num)
		}
	}
	return seeds
}
