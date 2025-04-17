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

func walk(maze [][]string, current Point, instructions *[]string, seen *[][]bool) bool {

	dirs := map[string][2]int{
		"^": {-1, 0},
		"v": {1, 0},
		"<": {0, -1},
		">": {0, 1},
	}

	// fmt.Println(maze[current.row][current.col])
	// No more instructions to execute we finished
	// fmt.Println(current.col, current.row)
	if len(*instructions) == 0 {
		return true
	}
	current_instruction := (*instructions)[len(*instructions)-1]

	if maze[current.row][current.col] == "#" {
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
		// Si se choca con una pared se regresa a la posicion a la que estaba
	}

    if maze[current.row][current.col] == "." {
        
    }

	(*instructions) = (*instructions)[:len(*instructions)-1]
	if len((*instructions)) > 0 {
		current_instruction = (*instructions)[len(*instructions)-1]
	}

	// curr_symbol := maze[current.row][current.col]
	// next_symbol := maze[current.row+dirs[current_instruction][0]][current.col+dirs[current_instruction][1]]

	// if (curr_symbol == "." || curr_symbol == "@") && next_symbol == "." {
	// 	maze[current.row][current.col] = "."
	// 	maze[current.row+dirs[current_instruction][0]][current.col+dirs[current_instruction][1]] = "@"
	// }

	// if curr_symbol == "@" && next_symbol == "O" {
	//
	// 	after_next_symbol := maze[current.row+(dirs[current_instruction][0]*2)][current.col+(dirs[current_instruction][1]*2)]
	// 	if after_next_symbol == "." {
	// 		maze[current.row][current.col] = "."
	// 		maze[current.row+(dirs[current_instruction][0])][current.col+dirs[current_instruction][1]] = "@"
	// 		maze[current.row+(dirs[current_instruction][0]*2)][current.col+(dirs[current_instruction][1]*2)] = "O"
	// 	}
	//
	// 	if after_next_symbol == "O" {
	// 		// fmt.Println("Empieza recursividad de nuevo")
	// 		llego_al_punto_o_hashtag := false
	// 		is_dot := false
	// 		// is_hashtag := false
	// 		multiplier := 2
	// 		var lastPoint Point
	// 		for !llego_al_punto_o_hashtag {
	// 			// fmt.Println(current.row+(dirs[current_instruction][0]*multiplier), current.col+(dirs[current_instruction][1]*multiplier))
	// 			fmt.Println(current_instruction)
	// 			if maze[current.row+(dirs[current_instruction][0]*multiplier)][current.col+(dirs[current_instruction][1]*multiplier)] == "." {
	// 				is_dot = true
	// 				lastPoint = Point{row: current.row + (dirs[current_instruction][0] * multiplier), col: current.col + (dirs[current_instruction][1] * multiplier)}
	// 				llego_al_punto_o_hashtag = true
	// 			}
	// 			// if maze[current.row+(dirs[current_instruction][0]*multiplier)][current.col+(dirs[current_instruction][1]*multiplier)] == "#" {
	// 			// 	// maze[current.row+(dirs[current_instruction][0]*multiplier)][current.col+(dirs[current_instruction][1]*multiplier)]
	// 			//
	// 			// 	// (*instructions) = (*instructions)[:len(*instructions)-1]
	// 			// 	lastPoint = Point{row: current.row + (dirs[current_instruction][0] * multiplier), col: current.col + (dirs[current_instruction][1] * multiplier)}
	// 			// 	is_hashtag = true
	// 			// 	llego_al_punto_o_hashtag = true
	// 			// 	fmt.Println("Choco", is_hashtag)
	// 			// }
	// 			multiplier = multiplier + 1
	// 		}
	//
	// 		if is_dot {
	// 			// Falta hacer si es validacion vertical
	// 			// fmt.Println(maze[current.row][current.col])
	// 			if lastPoint.col-current.col > 0 {
	// 				for i := lastPoint.col; i > current.col; i-- {
	// 					pivot := maze[lastPoint.row][i]
	// 					maze[lastPoint.row][i] = maze[lastPoint.row][i-1]
	// 					maze[lastPoint.row][i-1] = pivot
	// 					// fmt.Println(maze[lastPoint.row][i])
	// 					if maze[lastPoint.row][i] == "@" {
	// 						current.row = lastPoint.row
	// 						current.col = i
	// 					}
	// 				}
	// 			}
	// 			fmt.Println(maze[current.row][current.col])
	// 			// fmt.Println(lastPoint.row - current.row)
	// 			// if lastPoint.row-current.row > 0 {
	// 			// 	for i := lastPoint.row; i >= current.row; i-- {
	// 			// 		pivot := maze[i][lastPoint.col]
	// 			// 		maze[i][lastPoint.col] = maze[i-1][lastPoint.col]
	// 			// 		maze[i-1][lastPoint.col] = pivot
	// 			// 	}
	// 			// }
	// 		}
	//
	// 		// fmt.Println(current.col, current.row, is_dot, is_hashtag, lastPoint)
	// 	}
	// }

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

	robot_x := current_robot_position / len(maze_string_arr)
	// Se resta el numero de filas hasta llegar al robot @ para tener en consideracion los saltos de linea que descuadra el calculo
	robot_y := (current_robot_position % len(maze_string_arr)) - robot_x

	// fmt.Println(robot_y, robot_x, instructions_arr, len(maze), len(maze[0]))
	walk(maze, Point{col: robot_y, row: robot_x}, &instructions_arr, &seen)

	for _, v := range maze {
		for _, mv := range v {
			fmt.Print(mv)
		}
		fmt.Println()
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
