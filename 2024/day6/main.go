package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	part2()
}

type Point struct {
	x int
	y int
}

type Horientation struct {
	horientation string
	check        bool
}

func part1() {

	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	current_guard_position := strings.Index(file, "^")
	temp := strings.Split(file, "\n")

	// println(file)
	var maze [][]string
	var seen [][]bool
	var path [][]Point

	for _, v := range temp {
		stringArr := strings.Split(v, "")
		maze = append(maze, stringArr)
		tmp_seen := make([]bool, len(stringArr))
		tmp_path := make([]Point, len(stringArr))
		seen = append(seen, tmp_seen)
		path = append(path, tmp_path)
	}

	guard_y := current_guard_position / len(temp)
	guard_x := current_guard_position % (len(temp))

	// println(maze[guard_y][guard_x])
	// current_guard_line := len(temp)
	// println("Empieza en ", guard_x, guard_y)
	counter := 0
	horientation := "up"

	walk(maze, Point{x: guard_x, y: guard_y}, &horientation, &counter, &seen)

	println(counter)
}

func changeHorientation(actualHorientation string) string {
	switch actualHorientation {
	case "up":
		return "right"
	case "right":
		return "down"
	case "down":
		return "left"
	case "left":
		return "up"
	default:
		return "up"
	}
}

func walk(maze [][]string, current Point, horientation *string, counter *int, seen *[][]bool) bool {

	// println("Cuantas veces entra", *counter)
	dirs := map[string][2]int{
		"up":    {-1, 0},
		"down":  {1, 0},
		"left":  {0, -1},
		"right": {0, 1},
	}

	// Off the map we finished
	// println("Primera vali x", current.x < 0 || current.x >= len(maze[0]), current.x, len(maze[0]))
	// println("Segunda vali y", current.y >= len(maze)-1, current.y, len(maze))
	if current.x < 0 || current.x >= len(maze[0]) || current.y < 0 || current.y >= len(maze)-1 {
		// println("Llego al final en posicion", *counter)
		return true
	}

	// println(maze[current.y][current.x])
	// println(current.y, current.x)

	// On a wall
	if maze[current.y][current.x] == "#" {
		switch *horientation {
		case "up":
			current.x = current.x + dirs["down"][1]
			current.y = current.y + dirs["down"][0]
		case "right":
			current.x = current.x + dirs["left"][1]
			current.y = current.y + dirs["left"][0]
		case "down":
			current.x = current.x + dirs["up"][1]
			current.y = current.y + dirs["up"][0]
		case "left":
			current.x = current.x + dirs["right"][1]
			current.y = current.y + dirs["right"][0]
		}

		// *counter--
		*horientation = changeHorientation(*horientation)
	}
	if !(*seen)[current.y][current.x] {
		*counter++
	}

	(*seen)[current.y][current.x] = true
	if (walk(maze, Point{x: current.x + dirs[*horientation][1],
		y: current.y + dirs[*horientation][0]}, horientation, counter, seen)) {
		return true
	}

	return false
}

func walk2(maze [][]string, current Point, horientation *string, counter *int, seen *[][]Horientation) bool {

	// println("Cuantas veces entra", *counter)
	dirs := map[string][2]int{
		"up":    {-1, 0},
		"down":  {1, 0},
		"left":  {0, -1},
		"right": {0, 1},
	}

	// Se salio del laberinto
	if current.x < 0 || current.x >= len(maze[0]) || current.y < 0 || current.y >= len(maze)-1 {
		return false
	}

	// En una pared
	if maze[current.y][current.x] == "#" {
		switch *horientation {
		case "up":
			current.x = current.x + dirs["down"][1]
			current.y = current.y + dirs["down"][0]
		case "right":
			current.x = current.x + dirs["left"][1]
			current.y = current.y + dirs["left"][0]
		case "down":
			current.x = current.x + dirs["up"][1]
			current.y = current.y + dirs["up"][0]
		case "left":
			current.x = current.x + dirs["right"][1]
			current.y = current.y + dirs["right"][0]
		}

		*horientation = changeHorientation(*horientation)
		// println("Regresa una posicion")
		(*seen)[current.y][current.x].check = false
	}

	if (*seen)[current.y][current.x].check && *horientation == (*seen)[current.y][current.x].horientation {
		*counter++
		return true
	}

	// println(*horientation, (*seen)[current.y][current.x].horientation)
	// if (*seen)[current.y][current.x].check && (*seen)[current.y][current.x].horientation != *horientation {
	// 	// *counter++
	// 	// Dependiendo de la horientacion validar a los lados para determinar si ha pasado por ahi
	//
	// 	println(current.y, current.x, *horientation)
	// 	if maze[current.y][current.x] != "^" {
	//
	// 		switch *horientation {
	// 		case "up":
	// 			tmp_x := current.x + dirs["right"][1]
	// 			tmp_y := current.y + dirs["right"][0]
	// 			if (*seen)[tmp_y][tmp_x].check {
	// 				*counter++
	// 			}
	// 		case "right":
	// 			tmp_x := current.x + dirs["down"][1]
	// 			tmp_y := current.y + dirs["down"][0]
	//
	// 			if (*seen)[tmp_y][tmp_x].check {
	// 				*counter++
	// 			}
	// 		case "down":
	// 			tmp_x := current.x + dirs["down"][1]
	// 			tmp_y := current.y + dirs["down"][0]
	//
	// 			if (*seen)[tmp_y][tmp_x].check {
	// 				*counter++
	// 			}
	// 		case "left":
	// 			tmp_x := current.x + dirs["down"][1]
	// 			tmp_y := current.y + dirs["down"][0]
	//
	// 			if (*seen)[tmp_y][tmp_x].check {
	// 				*counter++
	// 			}
	// 		}
	// 	}
	// }

	// Actualizar posicion a visto
	(*seen)[current.y][current.x].check = true
	(*seen)[current.y][current.x].horientation = *horientation

	// Como aun no ha salido entrar en bucle
	if (walk2(maze, Point{x: current.x + dirs[*horientation][1],
		y: current.y + dirs[*horientation][0]}, horientation, counter, seen)) {
		return true
	}

	return false
}

func part2() {

	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	current_guard_position := strings.Index(file, "^")
	temp := strings.Split(file, "\n")

	// println(file)

	var maze [][]string

	var seen [][]Horientation
	for _, v := range temp {
		stringArr := strings.Split(v, "")
		maze = append(maze, stringArr)
		tmp_seen := make([]Horientation, len(stringArr))
		// tmp_path := make([]Point, len(stringArr))
		seen = append(seen, tmp_seen)
		// path = append(path, tmp_path)
	}

	guard_y := current_guard_position / len(temp)
	guard_x := current_guard_position % (len(temp))
	seen = [len(maze)][len(maze)[0]]Horientation{}
	// println(maze[guard_y][guard_x])
	// current_guard_line := len(temp)
	// println("Empieza en ", guard_x, guard_y)
	counter := 0
	horientation := "up"

	for i, mazeRow := range maze {
		for j, mazeCol := range mazeRow {
			if mazeCol != "#" && mazeCol != "^" {
				seen := [len(maze)][len(mazeRow)]Horientation{}
				maze[i][j] = "#"
				walk2(maze, Point{x: guard_x, y: guard_y}, &horientation, &counter, &seen)
				maze[i][j] = "."
				// resetArr(&seen)
				horientation = "up"
			}
		}
	}

	println(counter)
}

// func resetArr(test *[][]Horientation) {
// 	for i, v := range *test {
// 		for j, _ := range v {
// 			(*test)[i][j] = Horientation{}
// 		}
// 	}
// }
