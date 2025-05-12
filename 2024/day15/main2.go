package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readKey(input chan rune) {
	char, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	input <- char
}

func main2() {
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

	if maze[current.row][current.col] == "O" {
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

func move_for_instructions_left_right(maze *[][]string, current *Point, current_instruction string, dirs map[string][2]int) {
	multiplier := 1
	llego_al_punto_o_hashtag := false
	is_dot := false
	var lastPoint Point
	for !llego_al_punto_o_hashtag {
		if (*maze)[current.row+(dirs[current_instruction][0]*multiplier)][current.col+(dirs[current_instruction][1]*multiplier)] == "#" {
			lastPoint = Point{row: current.row + (dirs[current_instruction][0] * multiplier), col: current.col + (dirs[current_instruction][1] * multiplier)}
			is_dot = false
			llego_al_punto_o_hashtag = true
		}
		if (*maze)[current.row+(dirs[current_instruction][0]*multiplier)][current.col+(dirs[current_instruction][1]*multiplier)] == "." {
			lastPoint = Point{row: current.row + (dirs[current_instruction][0] * multiplier), col: current.col + (dirs[current_instruction][1] * multiplier)}
			is_dot = true
			llego_al_punto_o_hashtag = true
		}
		multiplier = multiplier + 1
	}

	if lastPoint.col-current.col > 0 {
		for i := lastPoint.col; i >= current.col; i-- {
			if is_dot {
				pivot := (*maze)[lastPoint.row][i]
				(*maze)[lastPoint.row][i] = (*maze)[lastPoint.row][i-1]
				(*maze)[lastPoint.row][i-1] = pivot
			} else {
				*current = back_one_position(*current, current_instruction, dirs)
			}

			if (*maze)[lastPoint.row][i] == "@" {
				current.row = lastPoint.row
				current.col = i
			}
		}
	}

	if lastPoint.col-current.col < 0 {
		for i := lastPoint.col; i <= current.col; i++ {
			if is_dot {
				pivot := (*maze)[lastPoint.row][i]
				(*maze)[lastPoint.row][i] = (*maze)[lastPoint.row][i+1]
				(*maze)[lastPoint.row][i+1] = pivot
			} else {
				*current = back_one_position(*current, current_instruction, dirs)
			}

			if (*maze)[lastPoint.row][i] == "@" {
				fmt.Println("Aqui se caga")
				current.row = lastPoint.row
				current.col = i
			}
		}
	}

}

// Part 2 helper functions
// func validate_to_right(maze *[][]string, current *Point, current_instruction string, dirs map[string][2]int, multiplier int) []Point {
// 	r_multiplier := 1
// 	f_multiplier := 0
// 	ended := false
// 	found_right_symbol := false
// 	var points_to_validate []Point
// 	r_first_symbol := (*maze)[(*current).row+dirs[current_instruction][0]*multiplier][(*current).col+f_multiplier]
// 	r_same_line_next_symbol := (*maze)[(*current).row+dirs[current_instruction][0]*(multiplier-1)][(*current).col+f_multiplier+1]
// 	extra_check := (*maze)[(*current).row+dirs[current_instruction][0]*(multiplier-1)][current.col]
//
// 	if extra_check != "]" && r_first_symbol == "." && r_same_line_next_symbol != "." {
// 		for !found_right_symbol {
// 			r_first_symbol = (*maze)[current.row+dirs[current_instruction][0]*multiplier][(*current).col+f_multiplier]
//
// 			if r_first_symbol == "[" || r_first_symbol == "]" {
// 				fmt.Println(f_multiplier)
// 				if f_multiplier <= 2 {
// 					(*current).col = current.col + f_multiplier
// 				}
// 				found_right_symbol = true
// 			}
//
// 			if r_first_symbol == "#" {
// 				found_right_symbol = true
// 			}
// 			f_multiplier++
// 		}
// 	}
// 	for !ended {
// 		rsymbol_to_check := (*maze)[current.row+dirs[current_instruction][0]*multiplier][current.col+r_multiplier]
//
// 		current_symbol := (*maze)[current.row+dirs[current_instruction][0]*(multiplier-1)][current.col]
//
// 		if rsymbol_to_check == "[" || rsymbol_to_check == "]" {
// 			if rsymbol_to_check == "[" {
// 				if r_multiplier == 1 && current_symbol != "]" {
// 					points_to_validate = append(points_to_validate, Point{row: current.row + dirs[current_instruction][0]*multiplier, col: current.col + r_multiplier})
// 				}
// 				if r_multiplier == 0 {
// 					points_to_validate = append(points_to_validate, Point{row: current.row + dirs[current_instruction][0]*multiplier, col: current.col + r_multiplier})
// 				}
// 			}
// 			if rsymbol_to_check == "]" {
//
// 				if r_multiplier-1 == 0 {
// 					points_to_validate = append(points_to_validate, Point{row: current.row + dirs[current_instruction][0]*multiplier, col: current.col + r_multiplier})
// 				}
// 				if r_multiplier-2 == 0 && current_symbol != "]" {
// 					points_to_validate = append(points_to_validate, Point{row: current.row + dirs[current_instruction][0]*multiplier, col: current.col + r_multiplier})
// 				}
// 			}
// 		}
// 		if rsymbol_to_check == "." || rsymbol_to_check == "#" {
// 			ended = true
// 		}
// 		r_multiplier++
// 	}
//
// 	return points_to_validate
// }

func validate_to_right(maze *[][]string, current *Point, current_instruction string, dirs map[string][2]int, multiplier int) []Point {
	var points_to_validate []Point
	visited_points := make(map[Point]bool) // Use a map to track visited points

	targetRow := current.row + dirs[current_instruction][0]*multiplier
	// Start scanning right from the column to the right of the robot's current position
	// or the column the robot would move into if it were not for obstacles
	// (depending on how 'current' is used in move_for_instructions_up_down)
	// Let's assume we are scanning from the column the robot is in relative to the target row
	scanCol := current.col

	r_multiplier := 0 // Start relative multiplier from 0

	for { // Loop indefinitely, break when a stop condition is met
		col_to_check := scanCol + r_multiplier

		if targetRow < 0 || targetRow >= len(*maze) || col_to_check < 0 || col_to_check >= len((*maze)[0]) {
			// Out of bounds
			break
		}

		symbol := (*maze)[targetRow][col_to_check]

		if symbol == "#" {
			// Hit a wall, stop scanning in this direction
			break
		}

		if symbol == "[" {
			// Found the left part of a box
			boxLeft := Point{row: targetRow, col: col_to_check}
			boxRight := Point{row: targetRow, col: col_to_check + 1}

			// Add both parts of the box if the right part is within bounds
			if col_to_check+1 < len((*maze)[0]) {
				// It's likely a valid box, add both parts if not already visited
				if _, exists := visited_points[boxLeft]; !exists {
					points_to_validate = append(points_to_validate, boxLeft)
					visited_points[boxLeft] = true
				}
				if _, exists := visited_points[boxRight]; !exists {
					points_to_validate = append(points_to_validate, boxRight)
					visited_points[boxRight] = true
				}
			} else {
				// Edge case: "[" at the very right edge without a matching "]"
				if _, exists := visited_points[boxLeft]; !exists {
					points_to_validate = append(points_to_validate, boxLeft)
					visited_points[boxLeft] = true
				}
			}
		} else if symbol == "]" {
			// Found the right part of a box
			boxLeft := Point{row: targetRow, col: col_to_check - 1}
			boxRight := Point{row: targetRow, col: col_to_check}

			// Add both parts of the box if the left part is within bounds
			if col_to_check-1 >= 0 {
				// It's likely a valid box, add both parts if not already visited
				if _, exists := visited_points[boxLeft]; !exists {
					points_to_validate = append(points_to_validate, boxLeft)
					visited_points[boxLeft] = true
				}
				if _, exists := visited_points[boxRight]; !exists {
					points_to_validate = append(points_to_validate, boxRight)
					visited_points[boxRight] = true
				}
			} else {
				// Edge case: "]" at the very left edge without a matching "["
				if _, exists := visited_points[boxRight]; !exists {
					points_to_validate = append(points_to_validate, boxRight)
					visited_points[boxRight] = true
				}
			}
		}

		// If we encounter a "." or an empty space, continue scanning.
		// We only stop on a "#".

		r_multiplier++ // Move to the next column to the right
	}

	return points_to_validate
}

func validate_to_left(maze *[][]string, current *Point, current_instruction string, dirs map[string][2]int, multiplier int) []Point {
	var points_to_validate []Point
	visited_points := make(map[Point]bool) // Use a map to track visited points and avoid duplicates

	targetRow := current.row + dirs[current_instruction][0]*multiplier
	// Start scanning left from the robot's current column
	scanCol := current.col

	for scanCol >= 0 {
		if targetRow < 0 || targetRow >= len(*maze) || scanCol >= len((*maze)[0]) {
			// Out of bounds
			break
		}

		symbol := (*maze)[targetRow][scanCol]

		if symbol == "#" {
			// Hit a wall, stop scanning in this direction
			break
		}

		if symbol == "[" {
			// Found the left part of a box
			boxLeft := Point{row: targetRow, col: scanCol}
			boxRight := Point{row: targetRow, col: scanCol + 1}

			// Add both parts of the box if they are within bounds and not already added
			if scanCol+1 < len((*maze)[0]) && ((*maze)[targetRow][scanCol+1] == "]" || (*maze)[targetRow][scanCol+1] == ".") { // Check if the right part is valid
				if _, exists := visited_points[boxLeft]; !exists {
					points_to_validate = append(points_to_validate, boxLeft)
					visited_points[boxLeft] = true
				}
				if _, exists := visited_points[boxRight]; !exists {
					points_to_validate = append(points_to_validate, boxRight)
					visited_points[boxRight] = true
				}
			} else {
				// If the "[" is not followed by a "]", it might be an invalid state or an edge case not fully understood from the prompt.
				// For now, we'll just add the "[" position.
				if _, exists := visited_points[boxLeft]; !exists {
					points_to_validate = append(points_to_validate, boxLeft)
					visited_points[boxLeft] = true
				}
			}
		} else if symbol == "]" {
			// Found the right part of a box
			boxLeft := Point{row: targetRow, col: scanCol - 1}
			boxRight := Point{row: targetRow, col: scanCol}

			// Add both parts of the box if they are within bounds and not already added
			if scanCol-1 >= 0 && ((*maze)[targetRow][scanCol-1] == "[" || (*maze)[targetRow][scanCol-1] == ".") { // Check if the left part is valid
				if _, exists := visited_points[boxLeft]; !exists {
					points_to_validate = append(points_to_validate, boxLeft)
					visited_points[boxLeft] = true
				}
				if _, exists := visited_points[boxRight]; !exists {
					points_to_validate = append(points_to_validate, boxRight)
					visited_points[boxRight] = true
				}
			} else {
				// If the "]" is not preceded by a "[", similar to the "[" case.
				if _, exists := visited_points[boxRight]; !exists {
					points_to_validate = append(points_to_validate, boxRight)
					visited_points[boxRight] = true
				}
			}
		}

		// Continue scanning left
		scanCol--
	}

	return points_to_validate
}

func removeDuplicates(strPoint []Point) []Point {
	allKeys := make(map[Point]bool)
	var list []Point

	for _, item := range strPoint {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}

	return list
}

func move_for_instructions_up_down(maze *[][]string, current *Point, current_instruction string, dirs map[string][2]int) {
	multiplier := 0
	ud_dot_hashtag := false

	var points_to_move []Point
	robot_position := Point{row: current.row + dirs[current_instruction][0]*-1, col: current.col}
	points_to_move = append(points_to_move, robot_position)

	for !ud_dot_hashtag {
		var points_to_validate []Point

		if multiplier == 0 {
			if (*maze)[current.row][current.col] == "]" {
				new_point := Point{row: current.row, col: current.col - 1}
				points_to_validate = append(points_to_validate, *current)
				points_to_validate = append(points_to_validate, new_point)
			}

			if (*maze)[current.row][current.col] == "[" {
				new_point := Point{row: current.row, col: current.col + 1}
				points_to_validate = append(points_to_validate, *current)
				points_to_validate = append(points_to_validate, new_point)
			}
		} else {

			possible_points_r := validate_to_right(maze, current, current_instruction, dirs, multiplier)
			possible_points_l := validate_to_left(maze, current, current_instruction, dirs, multiplier)
			points_to_validate = append(points_to_validate, possible_points_r...)
			points_to_validate = append(points_to_validate, possible_points_l...)

		}

		points_to_validate = removeDuplicates(points_to_validate)
		// Si uno de los puntos es # no hacer nada
		// Si entre los puntos hay [] agregar al array externo de puntos por mover
		// Si todos las siguientes posiciones son . se deja de iterar y se mueve
		keep_going := true
		all_dots := true
		dot_count := 0

		for i, v := range points_to_validate {
			if (*maze)[v.row+dirs[current_instruction][0]][v.col] == "#" {
				keep_going = false
				all_dots = false
			}

			if (*maze)[v.row+dirs[current_instruction][0]][v.col] == "." {
				dot_count++
			}

			if (*maze)[v.row+dirs[current_instruction][0]][v.col] == "[" || (*maze)[v.row+dirs[current_instruction][0]][v.col] == "]" {
				all_dots = false
			}
			if i == len(points_to_validate)-1 {
				if dot_count-1 == len(points_to_validate)-1 {
					keep_going = false
				}

			}
		}

		if all_dots {
			points_to_move = append(points_to_move, points_to_validate...)
			ud_dot_hashtag = true
		} else {
			if keep_going {
				fmt.Println(points_to_validate)
				points_to_move = append(points_to_move, points_to_validate...)
			} else {
				(*current).row = (*current).row + dirs[current_instruction][0]*-1
				points_to_move = []Point{}
				ud_dot_hashtag = true
			}
		}

		multiplier++
	}

	if len(points_to_move) > 1 {
		for i := len(points_to_move) - 1; i >= 0; i-- {
			pivot := (*maze)[points_to_move[i].row][points_to_move[i].col]
			(*maze)[points_to_move[i].row][points_to_move[i].col] = (*maze)[points_to_move[i].row+dirs[current_instruction][0]][points_to_move[i].col]
			(*maze)[points_to_move[i].row+dirs[current_instruction][0]][points_to_move[i].col] = pivot
		}
	}
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

func walk2(maze [][]string, current Point, instructions *[]string, seen *[][]bool) bool {

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
	fmt.Println(current_instruction)
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

	// var points_to_move []Point
	if maze[current.row][current.col] == "]" || maze[current.row][current.col] == "[" {
		if current_instruction == "<" || current_instruction == ">" {
			move_for_instructions_left_right(&maze, &current, current_instruction, dirs)
		} else {
			tmp := current.col

			move_for_instructions_up_down(&maze, &current, current_instruction, dirs)
			current.col = tmp
		}
	}

	(*instructions) = (*instructions)[:len(*instructions)-1]
	if len((*instructions)) > 0 {
		current_instruction = (*instructions)[len(*instructions)-1]
	}

	if (walk2(maze, Point{col: current.col + dirs[current_instruction][1], row: current.row + dirs[current_instruction][0]}, instructions, seen)) {
		return true
	}
	return false
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
	// for _, v := range maze {
	// 	for _, mv := range v {
	// 		fmt.Print(mv)
	// 	}
	// 	fmt.Println()
	// }

	walk2(maze, Point{col: robot_y, row: robot_x}, &instructions_arr, &seen)

	final_count := 0
	for i, v := range maze {
		for j, mv := range v {
			if mv == "[" {
				final_count += (100 * i) + j
			}
			fmt.Print(mv)
		}
		fmt.Println()
	}

	fmt.Println(final_count)
}
