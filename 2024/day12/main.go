package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	part1()
	// part2()
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
				garden_sections = append(garden_sections, line_result)
			}
		}
	}

	for i, val := range garden_sections {
		fmt.Println(i, val)
	}
}

func part2() {
	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	temp := strings.Split(file, "\n")

	fmt.Println(temp)
}
