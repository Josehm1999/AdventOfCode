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

type FreeSpace struct {
	start int
	end   int
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
	var free_space_record []FreeSpace
	free_space_counter := 0
	for i, val := range values_arr {
		parsed_val, _ := strconv.Atoi(val)
		if i%2 == 0 {
			// fmt.Println("This is multiple of two", i, parsed_val, i/2)
			temp_file := make([]string, parsed_val)

			for tf := range temp_file {
				str_val := strconv.Itoa(i / 2)
				temp_file[tf] = str_val
			}
			disk_map_uncompressed = append(disk_map_uncompressed, temp_file...)
		} else {
			if parsed_val == 0 {
				continue
			}
			temp_free_space := make([]string, parsed_val)
			for tpf := range temp_free_space {
				temp_free_space[tpf] = "."
			}
			// fmt.Println("This is the amount of free space", i, parsed_val)
			free_space_counter += parsed_val
			free_space_record = append(free_space_record, FreeSpace{start: len(disk_map_uncompressed), end: len(disk_map_uncompressed) + parsed_val})
			disk_map_uncompressed = append(disk_map_uncompressed, temp_free_space...)
		}
	}

	// fmt.Println((disk_map_uncompressed))
	endpoint := len(disk_map_uncompressed)
	for _, val := range free_space_record {
		amount_to_move := val.end - val.start
		// possible_files_to_move := disk_map_uncompressed[endpoint-amount_to_move : endpoint]

		// fmt.Println("New line")
		for i := endpoint - 1; i >= endpoint-amount_to_move; i-- {
			if disk_map_uncompressed[i] == "." {

				i = i - 1
				// fmt.Println(disk_map_uncompressed[i])
				endpoint = endpoint - 1
				// i--
			}
			// fmt.Println(disk_map_uncompressed[i])
			disk_map_uncompressed[val.start] = disk_map_uncompressed[i]
			disk_map_uncompressed[i] = "."
			val.start = val.start + 1
		}

		// Get to the first point possible "." then iterate from there until the end
		// Capture next possible position to move a file if its not possible and has reached the end means its done

		// for _, val := range possible_files_to_move {
		// 	if val == "." {
		// 		possible_files_to_move = disk_map_uncompressed[endpoint-amount_to_move-1 : endpoint-1]
		// 	}
		// }
		// for i, value := range possible_files_to_move {
		// 	disk_map_uncompressed[val.start+i] = value
		// 	disk_map_uncompressed[endpoint-amount_to_move+i] = "."
		// }

		// fmt.Println(amount_to_move, possible_files_to_move)
		endpoint -= amount_to_move

	}

	first_pos_dot := 0
	for i, ffdot := range disk_map_uncompressed {
		if ffdot == "." {
			first_pos_dot = i
			break
		}
	}

	next_available_dot := first_pos_dot

	for j := first_pos_dot; j < len(disk_map_uncompressed); j++ {
		if disk_map_uncompressed[j] != "." {
			disk_map_uncompressed[next_available_dot] = disk_map_uncompressed[j]
			disk_map_uncompressed[j] = "."

			next_available_dot++
			// fmt.Println(disk_map_uncompressed[j])
		}
		// fmt.Println(disk_map_uncompressed[j], next_available_dot)
	}

	checksum := 0
	for k, value := range disk_map_uncompressed {
		if value == "." {
			break
		}

		parsed_value, _ := strconv.Atoi(value)

		checksum += k * parsed_value
	}

	fmt.Println(disk_map_uncompressed,checksum)
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
