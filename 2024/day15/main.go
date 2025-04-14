package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	part1()
	// part2(101, 103)
}

type Point struct {
	col int
	row int
}

func walk(maze [][]string, start_position Point, instructions *[]string, seen *[][]bool) bool {

	dirs := map[string][2]int{
		"^": {0, -1},
		"v": {0, 1},
		"<": {-1, 0},
		">": {1, 0},
	}

	// No more instructions to execute we finished
	if len(*instructions) == 0 {
		return true
	}

    if maze[start_position.col][start_position.row] == "#" {

    }
	return false
}

func part1() {

	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	var maze [][]string
	var seen [][]bool
	var path [][]Point
	file := string(data)
	temp := strings.Split(file, "\n\n")

	maze_string_arr := strings.Split(temp[0], "\n")
	current_robot_position := strings.Index(temp[0], "@")

	instructions_arr := strings.Split(temp[1], "")
	instructions_arr = instructions_arr[:len(instructions_arr)-1]

	for i, j := 0, len(instructions_arr)-1; i < j; i, j = i+1, j-1 {
		instructions_arr[i], instructions_arr[j] = instructions_arr[j], instructions_arr[i]
	}

	for _, v := range maze_string_arr {
		maze_char_arr := strings.Split(v, "")
		maze = append(maze, maze_char_arr)
		tmp_seen := make([]bool, len(maze_char_arr))
		tmp_path := make([]Point, len(maze_char_arr))

		seen = append(seen, tmp_seen)
		path = append(path, tmp_path)
	}

	robot_y := current_robot_position / len(maze_string_arr)
	// Se resta el numero de filas hasta llegar al robot @ para tener en consideracion los saltos de linea que descuadra el calculo
	robot_x := (current_robot_position % len(maze_string_arr)) - robot_y

	fmt.Println(robot_y, robot_x, instructions_arr)
	walk(maze, Point{col: robot_y, row: robot_x}, &instructions_arr, &seen)
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
