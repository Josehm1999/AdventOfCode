package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	// part1()
	part2()
}

type Point struct {
	row int
	col int
}

func isOffMap(current Point, width int, height int) bool {

	return current.row < 0 || current.row > height || current.col < 0 || current.col >= width
}

func garden_map_get(garden_map [][]string, point Point) string {
	mapValue := garden_map[point.row][point.col]
	return mapValue
}

func walk(garden_map [][]string, current Point, seen *map[Point]bool) []Point {
	dirs := [4][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	var result []Point

	//La posicion actual se marca como vista
	(*seen)[current] = true
	result = append(result, current)

	//Recurse
	for _, value := range dirs {
		next_point := Point{col: current.col + value[1], row: current.row + value[0]}
		isIt := isOffMap(next_point, len(garden_map[0]), len(garden_map)-1)
		if !isIt {
			if garden_map_get(garden_map, current) == garden_map_get(garden_map, next_point) && !(*seen)[next_point] {
				new_line := walk(garden_map, next_point, seen)
				result = append(result, new_line...)
			}
		}
	}

	return result
}

func walk2(garden_map [][]string, current Point, seen *map[Point]bool) []Point {
	dirs := [4][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	var result []Point
	//La posicion actual se marca como vista
	(*seen)[current] = true
	result = append(result, current)

	//Recurse
	for _, value := range dirs {
		next_point := Point{col: current.col + value[1], row: current.row + value[0]}
		isIt := isOffMap(next_point, len(garden_map[0]), len(garden_map)-1)
		if !isIt {
			if garden_map_get(garden_map, current) == garden_map_get(garden_map, next_point) && !(*seen)[next_point] {
				new_line := walk(garden_map, next_point, seen)
				result = append(result, new_line...)
			}
		}
	}
	return result
}

func removeDuplicate(sliceList []Point) []Point {
	allKeys := make(map[string]bool)
	list := []Point{}

	for _, item := range sliceList {
		key := fmt.Sprintf("%v,%v", item.col, item.row)

		if _, value := allKeys[key]; !value {
			allKeys[key] = true
			list = append(list, item)
		}
	}

	return list
}

func part1() {

	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	temp := strings.Split(file, "\n")
	var garden_map [][]string
	for i := range temp {
		if temp[i] != "" {
			curr_line := strings.Split(temp[i], "")
			garden_map = append(garden_map, curr_line)
		}
	}

	//Start from 0,0
	seen := make(map[Point]bool)
	var garden_sections [][]Point
	for i := range garden_map {
		for r := range garden_map[i] {
			current_point := Point{row: i, col: r}
			_, ok := seen[current_point]
			if !ok {
				line_result := walk(garden_map, current_point, &seen)
				fmt.Println(line_result)
				garden_sections = append(garden_sections, line_result)
			}
		}
	}

	total := 0
	for k := range garden_sections {
		current_zone := 0
		for p := range garden_sections[k] {
			current_counter := 0
			for q := range garden_sections[k] {
				operation_col := (garden_sections[k][p].col - garden_sections[k][q].col)
				operation_row := (garden_sections[k][p].row - garden_sections[k][q].row)
				operation := math.Abs(float64(operation_col)) + math.Abs(float64(operation_row))

				if math.Abs(float64(operation)) == 1 {
					// fmt.Println(garden_map_get(garden_map, garden_sections[k][p]), garden_sections[k][p], garden_sections[k][q])
					current_counter++
				}
			}
			current_zone += 4 - current_counter
		}
		// fmt.Println(garden_map_get(garden_map, garden_sections[k][0]), len(garden_sections[k]), current_zone)
		total += len(garden_sections[k]) * current_zone
	}

	fmt.Println(total)
}

func part2() {
	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	temp := strings.Split(file, "\n")
	var garden_map [][]string
	for i := range temp {
		if temp[i] != "" {
			curr_line := strings.Split(temp[i], "")
			garden_map = append(garden_map, curr_line)
		}
	}

	//Start from 0,0
	seen := make(map[Point]bool)
	var garden_sections [][]Point
	for i := range garden_map {
		for r := range garden_map[i] {
			current_point := Point{row: i, col: r}
			_, ok := seen[current_point]
			if !ok {
				line_result := walk2(garden_map, current_point, &seen)
				// fmt.Println(line_result)
				garden_sections = append(garden_sections, line_result)
			}
		}
	}

	total := 0

	// for t := range garden_sections {
	// 	for x := range garden_sections[t] {
	// 		for y := x + 1; y < len(garden_sections[t]); y++ {
	// 			// fmt.Println(garden_sections[t][x], garden_sections[t][y])
	// 			if garden_sections[t][x].row == garden_sections[t][y].row && garden_sections[t][x].col < garden_sections[t][y].col {
	// 				garden_sections[t][x].col = garden_sections[t][y].col
	// 			}
	// 		}
	// 	}
	// 	// garden_sections[t] = removeDuplicate(garden_sections[t])
	// }
	//
	// fmt.Println(garden_sections)

	for k := range garden_sections {
		current_zone := 0
		curr_area := len(garden_sections[k])
		// garden_sections[k] = removeDuplicate(garden_sections[k])
		// fmt.Println(garden_sections[k])
		for p := range garden_sections[k] {
			current_counter := 0
			for q := range garden_sections[k] {

				operation_col := (garden_sections[k][p].col - garden_sections[k][q].col)
				operation_row := (garden_sections[k][p].row - garden_sections[k][q].row)
				operation := math.Abs(float64(operation_col)) + math.Abs(float64(operation_row))
				if math.Abs(float64(operation)) == 1 {
					current_counter++

					if garden_sections[k][p].row == garden_sections[k][q].row {
						current_counter++
					}
				}
			}
			current_zone += 4 - current_counter
		}

		total += curr_area * current_zone
		// fmt.Println(garden_map_get(garden_map, garden_sections[k][0]), curr_area, current_zone)
	}

	fmt.Println(total)
}
