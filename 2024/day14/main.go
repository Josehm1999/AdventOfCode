package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// part1(101, 103)
	part2(101, 103)
}

type Robot struct {
	row        int
	col        int
	velo_right int
	velo_down  int
}

func part1(grid_col int, grid_row int) {

	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	temp := strings.Split(file, "\n")

	var robots []Robot
	for i := range temp {
		// fmtPrintln(temp[i])
		if temp[i] != "" {

			parts := strings.Split(temp[i], " ")

			// fmt.Println(parts)

			position := strings.Split(parts[0], ",")
			velocities := strings.Split(parts[1], ",")

			equal_sign := strings.Index(position[0], "=")
			parsed_col, _ := strconv.Atoi(position[0][equal_sign+1:])

			parsed_row, _ := strconv.Atoi(position[1])

			equal_sign_velo := strings.Index(velocities[0], "=")
			parsed_vright, _ := strconv.Atoi(velocities[0][equal_sign_velo+1:])

			parsed_vdown, _ := strconv.Atoi(velocities[1])

			robot := Robot{
				col:        parsed_col,
				row:        parsed_row,
				velo_right: parsed_vright,
				velo_down:  parsed_vdown,
			}
			robots = append(robots, robot)
		}
	}

	var affected_row1 = 0
	affected_row1 = grid_row / 2

	var affected_col1 = 0
	affected_col1 = grid_col / 2

	num_seconds := 100

	quadrant1 := 0
	quadrant2 := 0
	quadrant3 := 0
	quadrant4 := 0
	for _, v := range robots {
		new_vright := v.velo_right * num_seconds
		new_vdown := v.velo_down * num_seconds

		new_col := (v.col + new_vright)

		new_col = new_col % grid_col

		if new_col < 0 {
			new_col = grid_col + new_col
		}
		new_row := (v.row + new_vdown)

		new_row = (new_row) % grid_row
		if new_row < 0 {
			new_row = grid_row + new_row
		}

		if new_col < affected_col1 && new_row < affected_row1 {
			quadrant1++
		}

		if new_col > affected_col1 && new_row < affected_row1 {
			quadrant2++
		}

		if new_col < affected_col1 && new_row > affected_row1 {
			quadrant3++
		}

		if new_col > affected_col1 && new_row > affected_row1 {
			quadrant4++
		}

		fmt.Println(new_col, new_row)
	}

	fmt.Println("Quadrant1", quadrant1)
	fmt.Println("Quadrant2", quadrant2)
	fmt.Println("Quadrant3", quadrant3)
	fmt.Println("Quadrant4", quadrant4)
	fmt.Println(quadrant1 * quadrant2 * quadrant3 * quadrant4)
	// fmt.Println(robots)
}

func part2(grid_col int, grid_row int) {

	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	temp := strings.Split(file, "\n")

	var robots []Robot
	for i := range temp {
		if temp[i] != "" {

			parts := strings.Split(temp[i], " ")

			// fmt.Println(parts)

			position := strings.Split(parts[0], ",")
			velocities := strings.Split(parts[1], ",")

			equal_sign := strings.Index(position[0], "=")
			parsed_col, _ := strconv.Atoi(position[0][equal_sign+1:])

			parsed_row, _ := strconv.Atoi(position[1])

			equal_sign_velo := strings.Index(velocities[0], "=")
			parsed_vright, _ := strconv.Atoi(velocities[0][equal_sign_velo+1:])

			parsed_vdown, _ := strconv.Atoi(velocities[1])

			robot := Robot{
				col:        parsed_col,
				row:        parsed_row,
				velo_right: parsed_vright,
				velo_down:  parsed_vdown,
			}
			robots = append(robots, robot)
		}
	}

	var affected_row1 = 0
	affected_row1 = grid_row / 2

	var affected_col1 = 0
	affected_col1 = grid_col / 2

	start := 8179
	num_seconds := 8180

	for start < num_seconds {

		var new_robot_arr []Robot
		for _, v := range robots {
			new_vright := v.velo_right * start
			new_vdown := v.velo_down * start

			new_col := (v.col + new_vright)

			new_col = new_col % grid_col

			if new_col < 0 {
				new_col = grid_col + new_col
			}
			new_row := (v.row + new_vdown)

			new_row = (new_row) % grid_row
			if new_row < 0 {
				new_row = grid_row + new_row
			}

			if affected_row1 != new_row && affected_col1 != new_col {

				newRobot := Robot{
					col: new_col,
					row: new_row,
				}

				new_robot_arr = append(new_robot_arr, newRobot)
			}
			// fmt.Println(new_col, new_row)
		}

		i := 0
		j := 0
		current := ""
		current_counter := 0
		for j < grid_row {
			for i < grid_col {
				printString := "."
				for _, v := range new_robot_arr {
					if v.col == i && v.row == j {
						if current == "#" {
							current_counter++
							// fmt.Println(current_counter)
							if current_counter >= 8 {
								// j = grid_row
								// fmt.Println(start)
								// break
							}
						}
						printString = "#"
					}
				}
				if printString == "." {
					current_counter = 0
				}
				current = printString
				fmt.Print(printString)
				i++
			}
			i = 0
			fmt.Print("\n")
			j++
		}

		// fmt.Println("", start)
		start++
	}
}
