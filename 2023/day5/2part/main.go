package main

import (
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

type SeedValues struct {
	src    int
	dest   int
	length int
}

type SeedRanges struct {
	start  int
	length int
}

func main() {
	file, err := os.ReadFile("../input_day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	contents := strings.Split(string(file[:]), "\n\n")
	seeds := strings.Split(strings.Split(contents[0], ": ")[1], " ")
	var seed_ranges_arr []SeedRanges

	for i := 0; i < len(seeds)-1; i = i + 2 {
		start, _ := strconv.Atoi(seeds[i])
		length, _ := strconv.Atoi(seeds[i+1])
		seed_ranges_arr = append(seed_ranges_arr, SeedRanges{start, length})
	}
	fmt.Println(seed_ranges_arr)
	_, new_content := contents[0], contents[1:]

	lowest_seed := 2147483647
	fmt.Println(new_content)

	// initial_range := SeedValues{
	// 	src:    2147483647,
	// 	dest:   2147483647,
	// 	length: 214748364,
	// }
	// for _, line := range new_content {
	// 	for _, seed_range := range seed_ranges_arr {
 //            for
	// 		curr_range := findMinimunRange(line, initial_range)
 //
	// 	}
	// 	fmt.Println(line)
	// }
	for _, seed_range := range seed_ranges_arr {
		for curr_seed := seed_range.start; curr_seed < seed_range.start+seed_range.length-1; curr_seed++ {
			curr_seed_int := curr_seed
			for i := 0; i < len(new_content); i++ {
				for _, val := range parseMap(new_content[i]) {
					if curr_seed_int >= val.src && curr_seed_int <= val.src+val.length-1 {
						curr_seed_int = curr_seed_int + val.dest - val.src
						break
					}
				}
			}

			if curr_seed_int < lowest_seed {
				lowest_seed = curr_seed_int
			}
		}
	}

	fmt.Println(lowest_seed)
}

func createRange(line string) SeedValues {
	items := strings.Split(line, " ")
	dest, _ := strconv.Atoi(items[0])
	src, _ := strconv.Atoi(items[1])
	length, _ := strconv.Atoi(items[2])
	seedValue := SeedValues{src, dest, length}
	return seedValue
}

func parseMap(data string) []SeedValues {
	var seed_values_arr []SeedValues
	lines := strings.Split(data, "\n")
	// src_to_dest := strings.Split(strings.Split(lines[0], " ")[0], "-")
	_, new_line := lines[0], lines[1:]
	// src := src_to_dest[0]
	// dest := src_to_dest[2]

	for _, val := range new_line {
		if val != "" {
			seed_values_arr = append(seed_values_arr, createRange(val))
		}
	}
	// fmt.Println(src, dest)
	return seed_values_arr
}

func findMinimunRange(line string, curr_range SeedValues) SeedValues {
	arr_values := parseMap(line)
	for _, value := range arr_values {
		if value.src+value.length < curr_range.src+curr_range.length {
			curr_range = value
		}
	}
	return curr_range
}
