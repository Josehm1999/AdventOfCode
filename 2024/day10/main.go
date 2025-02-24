package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2()
}

type Point struct {
	row int
	col int
}

func part1() {

	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	temp := strings.Split(file, "\n")

	var start_points []Point
	var topo_map [][]string

	for i, v := range temp {
		stringArr := strings.Split(v, "")
		for j, val := range stringArr {
			if val == "0" {
				start_points = append(start_points, Point{row: i, col: j})
			}
		}
		topo_map = append(topo_map, stringArr)
	}

	counter := 0
	// Start from every 0 position

	for _, start := range start_points {
		result := walk(topo_map, start)
		counter += len(removeDuplicate(result))
	}

	fmt.Println(counter)
}

func part2() {
	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	temp := strings.Split(file, "\n")

	var start_points []Point
	var topo_map [][]string

	for i, v := range temp {
		stringArr := strings.Split(v, "")
		for j, val := range stringArr {
			if val == "0" {
				start_points = append(start_points, Point{row: i, col: j})
			}
		}
		topo_map = append(topo_map, stringArr)
	}

	counter := 0
	// Start from every 0 position

	for _, start := range start_points {
		result := walk(topo_map, start)
		counter += len((result))
	}

	fmt.Println(counter)
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
func isOffMap(current Point, width int, height int) bool {

	return current.row < 0 || current.row >= height || current.col < 0 || current.col >= width
}
func topoMapGet(topoMap [][]string, point Point) int {
	mapValue := topoMap[point.row][point.col]

	parsedValue, err := strconv.Atoi(mapValue)

	if err != nil {
		return 0
	} else {
		return parsedValue
	}
}
func walk(topoMap [][]string, current Point) []Point {

	dirs := [4][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	var result []Point
	if topoMapGet(topoMap, current) == 9 {
		return append(result, current)
	} else {

		//Recurse
		for _, value := range dirs {
			nextPoint := Point{col: current.col + value[1], row: current.row + value[0]}

			//Check if the new point is out of bounds
			isIt := isOffMap(nextPoint, len(topoMap)-1, len(topoMap[0]))
			//Check if the current number follows the secuence
			if !isIt && topoMapGet(topoMap, current)+1 == topoMapGet(topoMap, nextPoint) {
				// fmt.Println(current, nextPoint, topoMapGet(topoMap, current), topoMapGet(topoMap, nextPoint))
				newList := walk(topoMap, nextPoint)
				result = append(result, newList...)
			}
		}

		return result
	}
}
