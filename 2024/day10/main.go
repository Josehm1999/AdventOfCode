package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	part1()
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
	var seen [][]bool
	var path [][]Point

	for i, v := range temp {
		stringArr := strings.Split(v, "")
		for j, val := range stringArr {
			if val == "0" {
				start_points = append(start_points, Point{row: i, col: j})
			}
		}
		topo_map = append(topo_map, stringArr)
		tmp_seen := make([]bool, len(stringArr))
		tmp_path := make([]Point, len(stringArr))
		seen = append(seen, tmp_seen)
		path = append(path, tmp_path)
	}

	counter := 0
	horientation := "down"
	// fmt.Println(topo_map, seen, path, counter)

	// Start from every 0 position
	for _, start := range start_points {
		walk(topo_map, start, &horientation, &counter, &seen)
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

	fmt.Println(temp)
}

func walk(maze [][]string, current Point, horientation *string, counter *int, seen *[][]bool) bool {
	// fmt.Println(maze, current, *counter, *seen)

	// dirs := map[string][2]int{
	// 	"up":    {-1, 0},
	// 	"down":  {1, 0},
	// 	"left":  {0, -1},
	// 	"right": {0, 1},
	// }

	// Off the map - bad
	if current.row < 0 || current.row >= len(maze[0]) || current.col < 0 || current.col >= len(maze)-1 {
		return true
	}

	//End of trail - good
	if maze[current.col][current.row] == "9" {
		*counter++
		return true
	}

    // If all sorounding nums don't follow the secuence and

	//Check if the current number follows the secuence

	//Recurse

	return false
}
