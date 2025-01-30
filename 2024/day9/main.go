package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part1()
}

type DiskSpace struct {
	value string
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
	var free_space_record []DiskSpace
	var file_space_record []DiskSpace
	var disk_map_uncompressed []string
	for i, val := range values_arr {
		parsed_val, _ := strconv.Atoi(val)
		if i%2 == 0 {
			temp_file := make([]string, parsed_val)

			for tf := range temp_file {
				str_val := strconv.Itoa(i / 2)
				temp_file[tf] = str_val
			}

			file_space_record = append(file_space_record, DiskSpace{value: temp_file[0], start: len(disk_map_uncompressed), end: len(disk_map_uncompressed) + parsed_val})
			disk_map_uncompressed = append(disk_map_uncompressed, temp_file...)
		} else {
			if parsed_val == 0 {
				continue
			}
			temp_free_space := make([]string, parsed_val)
			for tpf := range temp_free_space {
				temp_free_space[tpf] = "."
			}

			free_space_record = append(free_space_record, DiskSpace{value: "", start: len(disk_map_uncompressed), end: len(disk_map_uncompressed) + parsed_val})
			disk_map_uncompressed = append(disk_map_uncompressed, temp_free_space...)
		}
	}

	fmt.Println((disk_map_uncompressed))
	// for i, val := range disk_map_uncompressed {

	freqCmp := func(a, b DiskSpace) int {
		return cmp.Compare(b.start, a.start)
	}
	slices.SortFunc(file_space_record, freqCmp)

	for i := range file_space_record {
		for j := range free_space_record {
			fmt.Println(file_space_record[i], free_space_record[j])
		}
	}
	// for i := range free_space_record {
	// 	for _, file_space_range := range file_space_record {
	//
	// 		// Avoid looking for file spaces before free spaces
	// 		if free_space_record[i].start < file_space_range.start {
	// 			// Check if file space fits the free space
	// 			if free_space_record[i].end-free_space_record[i].start >= file_space_range.end-file_space_range.start {
	// 				fmt.Println("Matches", free_space_record[i], file_space_range)
	// 				file_space := disk_map_uncompressed[file_space_range.start:file_space_range.end]
	//
	// 				//Copy file spaces into the free spaces
	// 				copy(disk_map_uncompressed[free_space_record[i].start:free_space_record[i].start+file_space_range.end-file_space_range.start], file_space)
	//
	// 				//Replace file spaces with "."
	// 				for k := file_space_range.start; k < file_space_range.end; k++ {
	// 					disk_map_uncompressed[k] = "."
	// 				}
	//
	// 				if free_space_record[i].end-free_space_record[i].start > file_space_range.end-file_space_range.start {
	// 					free_space_record[i].start += file_space_range.end - file_space_range.start
	// 					fmt.Println(free_space_record[i])
	// 				}
	//
	// 				if free_space_record[i].end-free_space_record[i].start == file_space_range.end-file_space_range.start {
	// 					break
	// 				}
	// 				file_space_record = file_space_record[1:]
	// 				fmt.Println((disk_map_uncompressed))
	//
	// 				// break
	// 			}
	// 		}
	// 	}
	// 	// fmt.Println(i, val)
	// }
	fmt.Println((disk_map_uncompressed))
	// Part1
	// for i := 0; i < len(disk_map_uncompressed)-1; i++ {
	// 	if disk_map_uncompressed[i] == "." {
	// possible_num := disk_map_uncompressed[len(disk_map_uncompressed)-1]
	// disk_map_uncompressed = disk_map_uncompressed[:len(disk_map_uncompressed)-1]
	//
	// for possible_num == "." && len(disk_map_uncompressed) > i {
	// 	possible_num = disk_map_uncompressed[len(disk_map_uncompressed)-1]
	// 	disk_map_uncompressed = disk_map_uncompressed[:len(disk_map_uncompressed)-1]
	// }
	//
	// if len(disk_map_uncompressed) > i {
	// 	disk_map_uncompressed[i] = possible_num
	// }
	// 	}
	// }

	checksum := 0
	for k, value := range disk_map_uncompressed {
		parsed_value, error := strconv.Atoi(value)
		if error == nil {
			checksum += k * parsed_value
		}
	}

	fmt.Println(checksum)
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
