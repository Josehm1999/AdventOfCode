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
}

func part1() {

	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	temp := strings.Split(file, "\n")

	values_arr := strings.Split(temp[0], "")

	var disk_map_uncompressed []string
	for i, val := range values_arr {
		parsed_val, _ := strconv.Atoi(val)
		if i%2 == 0 {
			fmt.Println("This is multiple of two", i, parsed_val)
		} else {
			temp_free_space := make([]string, parsed_val)

			for tpf := range temp_free_space {
				temp_free_space[tpf] = "."
			}
			fmt.Println("This is the amount of free space", i, parsed_val)
			disk_map_uncompressed = append(disk_map_uncompressed, temp_free_space...)
		}
	}

	fmt.Println(slidingWindow(5, values_arr), disk_map_uncompressed)

}

func slidingWindow(size int, input []string) [][]string {
	if len(input) <= size {
		return [][]string{input}
	}

	r := make([][]string, 0, len(input)-size+1)

	for i, j := 0, size; j <= len(input); i, j = i+1, j+1 {

		r = append(r, input[i:j])
	}

	return r
}
