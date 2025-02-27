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
				line_result := walk(garden_map, current_point, &seen)
				// fmt.Println(line_result)
				garden_sections = append(garden_sections, line_result)
			}
		}
	}

	total := 0

	map_arr := make(map[string][][]Point)

	for t := range garden_sections {
		var arr_arr [][]Point
		curr_letter := garden_map_get(garden_map, garden_sections[t][0])
		// arr = append(arr, garden_sections[t][0])
		for x := range garden_sections[t] {
			var arr []Point
			for y := range garden_sections[t] {
				// fmt.Println(garden_sections[t][x], garden_sections[t][y])

				if garden_sections[t][x].row == garden_sections[t][y].row {
					// fmt.Println(garden_sections[t][x])
					arr = append(arr, garden_sections[t][y])
				}
			}
			arr_arr = append(arr_arr, arr)
		}
		arr_arr = removeDuplicateArrOfArr(arr_arr)
		key := fmt.Sprintf("%v,%v", curr_letter, t)
		map_arr[key] = arr_arr
	}

	// for k := range map_arr {
	// 	fmt.Println(map_arr[k])
	// }

	// fmt.Println(map_arr["B"])

	// No esta evaluando bien los puntos de conexion de la fila con las columnas
	for k := range map_arr {
		if len(map_arr[k]) == 1 {
			// fmt.Println(k, map_arr[k][0])
			total += 4 * len(map_arr[k][0])
		} else {

			// fmt.Println(k, map_arr[k])
			//First list to iterate
			line_count := 0
			items_count := 0
			for i, item := range map_arr[k] {
				initial_sides_count := 4
				// //First list items
				for _, nested_item := range item {
					// fmt.Println(nested_item)

					items_count++
					// Second list to iterate
					for j, item_second_list := range map_arr[k] {
						// Secord list items

						if i != j {
							for _, nested_item_second_list := range item_second_list {
								// 				// fmt.Println(key)
								operation_col := (nested_item.col - nested_item_second_list.col)
								operation_row := (nested_item.row - nested_item_second_list.row)
								operation := math.Abs(float64(operation_col)) + math.Abs(float64(operation_row))

								if math.Abs(float64(operation)) == 1 {
									fmt.Println(k, i, item)
									if nested_item_second_list.col == nested_item.col {
										fmt.Println("aqui", k, i, nested_item_second_list, nested_item)
										initial_sides_count--
									}
								}
							}
						}
					}
				}
				// }
				//
				// // fmt.Println(initial_sides_count)
				line_count += initial_sides_count
			}
			fmt.Println(k, items_count)
			total += line_count * items_count
			// 		// fmt.Println(k, max_perimeter, len(map_arr[k]))
		}
	}
	// for k := range garden_sections {
	// 	current_zone := 0
	// 	curr_area := len(garden_sections[k])
	// 	// garden_sections[k] = removeDuplicate(garden_sections[k])
	// 	// fmt.Println(garden_sections[k], is_one_line(garden_sections[k]))
	//
	// 	for p := range garden_sections[k] {
	// 		current_counter := 0
	// 		// visited_row := 999
	// 		// visited_col := 999
	// 		// current_col := garden_sections[k][p].col
	// 		for q := range garden_sections[k] {
	// 			operation_col := (garden_sections[k][p].col - garden_sections[k][q].col)
	// 			operation_row := (garden_sections[k][p].row - garden_sections[k][q].row)
	// 			operation := math.Abs(float64(operation_col)) + math.Abs(float64(operation_row))
	// 			if math.Abs(float64(operation)) == 1 {
	// 				// fmt.Println(visited_row)
	// 				current_counter++
	//
	// 				// if garden_sections[k][p].row == garden_sections[k][q].row {
	// 				// 	current_counter++
	// 				// }
	// 			}
	// 		}
	// 		current_zone += 4 - current_counter
	// 		// fmt.Println(current_zone)
	// 	}
	// 	total += curr_area * current_zone
	// 	fmt.Println(garden_map_get(garden_map, garden_sections[k][0]), curr_area, current_zone)
	// }
	// }

	fmt.Println(total)
}

func is_one_line(line []Point) bool {

	curr_row := 0
	for i, val := range line {
		if i == 0 {
			curr_row = val.row
		} else {
			if curr_row != val.row {
				return false
			}
		}
	}
	return true
}
