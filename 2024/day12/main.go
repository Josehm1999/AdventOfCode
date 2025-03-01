package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
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
		if !isIt && garden_map_get(garden_map, current) == garden_map_get(garden_map, next_point) && !(*seen)[next_point] {
			new_line := walk(garden_map, next_point, seen)
			result = append(result, new_line...)
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

func removeDuplicateArrOfArr(sliceList [][]Point) [][]Point {
	allKeys := make(map[string]bool)
	list := [][]Point{}

	for _, item := range sliceList {
		var line_key string
		for _, nested_item := range item {
			key := fmt.Sprintf("%v,%v", nested_item.col, nested_item.row)
			line_key += key
		}

		if _, value := allKeys[line_key]; !value {
			allKeys[line_key] = true
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

func exists_in_garden(sliceList []Point, coords []Point) []int {

	var res_arr []int
	for _, value := range coords {

		exists := false
		for _, list_val := range sliceList {
			if value == list_val {
				exists = true
			}
		}

		if exists {
			res_arr = append(res_arr, 1)
		} else {
			res_arr = append(res_arr, 0)
		}
	}
	return res_arr
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

	seen := make(map[Point]bool)
	var garden_sections [][]Point
	for i := range garden_map {
		for r := range garden_map[i] {
			current_point := Point{row: i, col: r}
			_, ok := seen[current_point]
			if !ok {
				line_result := walk(garden_map, current_point, &seen)
				garden_sections = append(garden_sections, line_result)
			}
		}
	}

	total := 0

	for _, line := range garden_sections {
		mini, maxi, minj, maxj := find_bounding_box(line)
		arr_corners := [][]int{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
			{1, 1, 1, 0},
			{1, 1, 0, 1},
			{1, 0, 1, 1},
			{0, 1, 1, 1},
		}

		arr_double_corners := [][]int{
			{1, 0, 0, 1},
			{0, 1, 1, 0},
		}

		item_corner := 0

		for rows := mini - 1; rows <= maxi; rows++ {
			for cols := minj - 1; cols <= maxj; cols++ {
				res := exists_in_garden(line, []Point{{row: rows, col: cols}, {row: rows + 1, col: cols}, {row: rows, col: cols + 1}, {row: rows + 1, col: cols + 1}})
				for _, poss_cor := range arr_corners {
					if slices.Equal(poss_cor, res) {
						item_corner++
					}
				}
				for _, poss_double_cor := range arr_double_corners {
					if slices.Equal(poss_double_cor, res) {
						item_corner += 2
					}
				}
			}
		}

		total += item_corner * len(line)
	}

	fmt.Println(total)
}

func find_bounding_box(line []Point) (mini, maxi, minj, maxj int) {

	if len(line) > 0 {
		mini, maxi = line[0].row, line[0].row
		minj, maxj = line[0].col, line[0].col
	} else {
		return 0, 0, 0, 0
	}

	for _, point := range line {
		i, j := point.row, point.col

		if i < mini {
			mini = i
		}
		if i > maxi {
			maxi = i
		}
		if j < minj {
			minj = j
		}
		if j > maxj {
			maxj = j
		}
	}

	return mini, maxi, minj, maxj
}
