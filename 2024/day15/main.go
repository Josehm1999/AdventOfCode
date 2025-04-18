package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// part1()
	part2()
}

type Point struct {
	col int
	row int
}

func walk(maze [][]string, current Point, instructions *[]string, seen *[][]bool) bool {

	dirs := map[string][2]int{
		"^": {-1, 0},
		"v": {1, 0},
		"<": {0, -1},
		">": {0, 1},
	}

	// No more instructions to execute we finished
	if len(*instructions) == 0 {
		return true
	}
	current_instruction := (*instructions)[len(*instructions)-1]

	if maze[current.row][current.col] == "#" {
		current = back_one_position(current, current_instruction, dirs)
	}

	if maze[current.row][current.col] == "." {
		maze[current.row][current.col] = "@"
		switch current_instruction {
		case "<":
			tmp_col := current.col + dirs[">"][1]
			tmp_row := current.row + dirs[">"][0]
			maze[tmp_row][tmp_col] = "."
		case ">":
			tmp_col := current.col + dirs["<"][1]
			tmp_row := current.row + dirs["<"][0]
			maze[tmp_row][tmp_col] = "."
		case "^":
			tmp_col := current.col + dirs["v"][1]
			tmp_row := current.row + dirs["v"][0]
			maze[tmp_row][tmp_col] = "."
		case "v":

			tmp_col := current.col + dirs["^"][1]
			tmp_row := current.row + dirs["^"][0]
			maze[tmp_row][tmp_col] = "."
		}
	}

	if maze[current.row][current.col] == "]" {
		if current_instruction == "<" || current_instruction == ">" {
			multiplier := 1
			llego_al_punto_o_hashtag := false
			is_dot := false
			// is_hastag := false
			var lastPoint Point
			for !llego_al_punto_o_hashtag {
				if maze[current.row+(dirs[current_instruction][0]*multiplier)][current.col+(dirs[current_instruction][1]*multiplier)] == "#" {
					lastPoint = Point{row: current.row + (dirs[current_instruction][0] * multiplier), col: current.col + (dirs[current_instruction][1] * multiplier)}
					is_dot = false
					llego_al_punto_o_hashtag = true
				}
				if maze[current.row+(dirs[current_instruction][0]*multiplier)][current.col+(dirs[current_instruction][1]*multiplier)] == "." {
					lastPoint = Point{row: current.row + (dirs[current_instruction][0] * multiplier), col: current.col + (dirs[current_instruction][1] * multiplier)}
					is_dot = true
					llego_al_punto_o_hashtag = true
				}
				multiplier = multiplier + 1
			}

			if lastPoint.col-current.col > 0 {
				for i := lastPoint.col; i >= current.col; i-- {
					if is_dot {
						pivot := maze[lastPoint.row][i]
						maze[lastPoint.row][i] = maze[lastPoint.row][i-1]
						maze[lastPoint.row][i-1] = pivot
					} else {
						current = back_one_position(current, current_instruction, dirs)
					}

					if maze[lastPoint.row][i] == "@" {
						current.row = lastPoint.row
						current.col = i
					}
				}
			}

			if lastPoint.col-current.col < 0 {
				for i := lastPoint.col; i <= current.col; i++ {
					if is_dot {
						pivot := maze[lastPoint.row][i]
						maze[lastPoint.row][i] = maze[lastPoint.row][i+1]
						maze[lastPoint.row][i+1] = pivot
					} else {
						current = back_one_position(current, current_instruction, dirs)
					}

					if maze[lastPoint.row][i] == "@" {
						current.row = lastPoint.row
						current.col = i
					}
				}
			}

			if lastPoint.row-current.row > 0 {
				for i := lastPoint.row; i >= current.row; i-- {
					if is_dot {
						pivot := maze[i][lastPoint.col]
						maze[i][lastPoint.col] = maze[i-1][lastPoint.col]
						maze[i-1][lastPoint.col] = pivot
					} else {
						current = back_one_position(current, current_instruction, dirs)
					}

					if maze[i][lastPoint.col] == "@" {
						current.row = i
						current.col = lastPoint.col
					}
				}
			}

			if lastPoint.row-current.row < 0 {

				for i := lastPoint.row; i <= current.row; i++ {
					if is_dot {
						pivot := maze[i][lastPoint.col]
						maze[i][lastPoint.col] = maze[i+1][lastPoint.col]
						maze[i+1][lastPoint.col] = pivot
					} else {
						current = back_one_position(current, current_instruction, dirs)
					}

					if maze[i][lastPoint.col] == "@" {
						current.row = i
						current.col = lastPoint.col
					}
				}
			}
			// fmt.Println(current)
			// fmt.Println(lastPoint)
		}

		if current_instruction == "^" || current_instruction == "v" {
			fmt.Println("Do nothing")
		}
	}

	(*instructions) = (*instructions)[:len(*instructions)-1]
	if len((*instructions)) > 0 {
		current_instruction = (*instructions)[len(*instructions)-1]
	}

	if (walk(maze, Point{col: current.col + dirs[current_instruction][1], row: current.row + dirs[current_instruction][0]}, instructions, seen)) {

		// fmt.Println(current.col, current.row, maze[current.row][current.col])
		return true
	}
	return false
}

func back_one_position(current Point, current_instruction string, dirs map[string][2]int) Point {
	switch current_instruction {
	case "<":
		current.col = current.col + dirs[">"][1]
		current.row = current.row + dirs[">"][0]
	case ">":
		current.col = current.col + dirs["<"][1]
		current.row = current.row + dirs["<"][0]
	case "^":
		current.col = current.col + dirs["v"][1]
		current.row = current.row + dirs["v"][0]
	case "v":
		current.col = current.col + dirs["^"][1]
		current.row = current.row + dirs["^"][0]
	}

	return current
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

	current_robot_position := strings.Index(temp[0], "@")
	dirs := map[string][2]int{
		"^": {-1, 0},
		"v": {1, 0},
		"<": {0, -1},
		">": {0, 1},
	}

	dir_first_ins := dirs[instructions_arr[len(instructions_arr)-1]]
	robot_x := current_robot_position/len(maze_string_arr) + dir_first_ins[0]
	// Se resta el numero de filas hasta llegar al robot @ para tener en consideracion los saltos de linea que descuadra el calculo
	robot_y := ((current_robot_position % len(maze_string_arr)) - robot_x) + dir_first_ins[1]

	// fmt.Println(robot_y, robot_x, instructions_arr, len(maze), len(maze[0]))
	walk(maze, Point{col: robot_y, row: robot_x}, &instructions_arr, &seen)

	final_count := 0
	for i, v := range maze {
		for j, mv := range v {
			if mv == "O" {
				final_count += (100 * i) + j
			}
			fmt.Print(mv)
		}
		fmt.Println()
	}

	fmt.Println(final_count)
}

func part2() {

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

	instructions_arr := strings.Split(temp[1], "")
	instructions_arr = instructions_arr[:len(instructions_arr)-1]

	for i, j := 0, len(instructions_arr)-1; i < j; i, j = i+1, j-1 {
		instructions_arr[i], instructions_arr[j] = instructions_arr[j], instructions_arr[i]
	}

	robot_x := 0
	robot_y := 0
	for i, v := range maze_string_arr {
		maze_char_arr := strings.Split(v, "")

		var new_maze_char_arr []string
		for j, r := range maze_char_arr {
			// fmt.Println(r)
			switch r {
			case "#":
				new_maze_char_arr = append(new_maze_char_arr, []string{"#", "#"}...)
				break
			case "@":
				robot_x = i
				robot_y = j * 2
				new_maze_char_arr = append(new_maze_char_arr, []string{"@", "."}...)
				break
			case "O":
				new_maze_char_arr = append(new_maze_char_arr, []string{"[", "]"}...)
				break
			case ".":
				new_maze_char_arr = append(new_maze_char_arr, []string{".", "."}...)
				break
			}
		}
		maze = append(maze, new_maze_char_arr)
		tmp_seen := make([]bool, len(new_maze_char_arr))
		tmp_path := make([]Point, len(new_maze_char_arr))

		seen = append(seen, tmp_seen)
		path = append(path, tmp_path)
	}

	dirs := map[string][2]int{
		"^": {-1, 0},
		"v": {1, 0},
		"<": {0, -1},
		">": {0, 1},
	}

	dir_first_ins := dirs[instructions_arr[len(instructions_arr)-1]]
	robot_x += dir_first_ins[0]
	// Se resta el numero de filas hasta llegar al robot @ para tener en consideracion los saltos de linea que descuadra el calculo
	robot_y += dir_first_ins[1]
	//
	fmt.Println(robot_y, robot_x, instructions_arr, len(maze), len(maze[0]))
	walk(maze, Point{col: robot_y, row: robot_x}, &instructions_arr, &seen)

	final_count := 0
	for i, v := range maze {
		for j, mv := range v {
			if mv == "O" {
				final_count += (100 * i) + j
			}
			fmt.Print(mv)
		}
		fmt.Println()
	}

	fmt.Println(final_count)
}
