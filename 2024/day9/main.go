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

	//Part 2
	for i := range file_space_record {

		// fmt.Println(file_space_record[i])
		for j := range free_space_record {
			//Avoid looking for file spaces before free_spaces
			if free_space_record[j].start > file_space_record[i].start {
				break
			}

			// Check if file space fits the free space
			if free_space_record[j].end-free_space_record[j].start < file_space_record[i].end-file_space_record[i].start {
				continue
			}

			// Fits the free space so it should replace the dots
			file_space := disk_map_uncompressed[file_space_record[i].start:file_space_record[i].end]
			copy(disk_map_uncompressed[free_space_record[j].start:free_space_record[j].start+file_space_record[i].end-file_space_record[i].start], file_space)

			// Replace the file places with dots
			for k := file_space_record[i].start; k < file_space_record[i].end; k++ {
				disk_map_uncompressed[k] = "."
			}

			// If the file fits buts doesn't take all the free space update the free space if some other file could use the space
			free_space_record[j].start += file_space_record[i].end - file_space_record[i].start
			break
		}
	}
	fmt.Println((disk_map_uncompressed))

	checksum := 0
	for k, value := range disk_map_uncompressed {
		parsed_value, error := strconv.Atoi(value)
		if error == nil {
			checksum += k * parsed_value
		}
	}

	fmt.Println(checksum)
}

func part2() {
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

	//Part 2
	for i := range file_space_record {
		// fmt.Println(file_space_record[i])
		for j := range free_space_record {
			//Avoid looking for file spaces before free_spaces
			if free_space_record[j].start > file_space_record[i].start {
				break
			}

			// Check if file space fits the free space
			if free_space_record[j].end-free_space_record[j].start < file_space_record[i].end-file_space_record[i].start {
				continue
			}

			// Fits the free space so it should replace the dots
			file_space := disk_map_uncompressed[file_space_record[i].start:file_space_record[i].end]
			copy(disk_map_uncompressed[free_space_record[j].start:free_space_record[j].start+file_space_record[i].end-file_space_record[i].start], file_space)

			// Replace the file places with dots
			for k := file_space_record[i].start; k < file_space_record[i].end; k++ {
				disk_map_uncompressed[k] = "."
			}

			// If the file fits buts doesn't take all the free space update the free space if some other file could use the space
			free_space_record[j].start += file_space_record[i].end - file_space_record[i].start
			break
		}
	}
	fmt.Println((disk_map_uncompressed))

	checksum := 0
	for k, value := range disk_map_uncompressed {
		parsed_value, error := strconv.Atoi(value)
		if error == nil {
			checksum += k * parsed_value
		}
	}

	fmt.Println(checksum)
}
