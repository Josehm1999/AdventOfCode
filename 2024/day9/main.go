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

	fmt.Println((disk_map_uncompressed))
	// fmt.Println((disk_map_uncompressed[len(disk_map_uncompressed)-1]))
	// fmt.Println(free_space_record[len])

	// 1.- Contar cuantas posiciones desde el final se deben de tomar
	// 2.- Mover caracteres desde la posicion final a una variable temporal mientras no sean .

	// 1.- Find the most right value that isn't a .
	// 2.- Find the most left value that is a .
	// 3.- Keep switching until the first left position that is a . and the furthest right position summed up give me the total number of free spaces

	endpoint := len(disk_map_uncompressed)
	for _, val := range free_space_record {
		amount_to_move := val.end - val.start
		possible_files_to_move := disk_map_uncompressed[endpoint-amount_to_move : endpoint]

		for _, val := range possible_files_to_move {
			if val == "." {
				possible_files_to_move = disk_map_uncompressed[endpoint-amount_to_move-1 : endpoint-1]
			}
		}
		for i, value := range possible_files_to_move {
			disk_map_uncompressed[val.start+i] = value
			disk_map_uncompressed[endpoint-amount_to_move+i] = "."
            fmt.Println()
		}

		// fmt.Println(amount_to_move, possible_files_to_move)
		endpoint -= amount_to_move

		// fmt.Println(disk_map_uncompressed[len(disk_map_uncompressed)-amount_to_move : len(disk_map_uncompressed)])
	}

	fmt.Println((disk_map_uncompressed))
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
