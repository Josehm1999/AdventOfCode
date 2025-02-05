package main

import (
	"fmt"
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
	point        Point
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

func walk2(maze [][]string, current Point, horientation *string, counter *int, seen *map[string]bool) bool {

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

		key := fmt.Sprintf("%d,%d,%s", current.x, current.y, *horientation)
		(*seen)[key] = false

		*horientation = changeHorientation(*horientation)

		//To do instead of storing all the posible postions and changing the bool value i need to convert the y,x and dir values into a
		//string key to use a map and speed up the search
		// (*seen) = false

	}

	key := fmt.Sprintf("%d,%d,%s", current.x, current.y, *horientation)
	if (*seen)[key] == true {
		*counter++
		return true
	}

	// Actualizar posicion a visto
	(*seen)[key] = true
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
	// var seen = [][]Horientation{}

	// println(len(temp))
	// maze := make([][]Horientation, len(temp))
	for _, v := range temp {
		stringArr := strings.Split(v, "")
		maze = append(maze, stringArr)
		// tmp_seen := make([]Horientation, len(stringArr))
		// seen = append(seen, tmp_seen)
	}

	guard_y := current_guard_position / len(temp)
	guard_x := current_guard_position % (len(temp))
	counter := 0
	horientation := "up"

	// println(len(seen)-1, len(seen[0]))
	for i, mazeRow := range maze {
		for j, mazeCol := range mazeRow {
			if mazeCol == "#" || mazeCol == "^" {
				continue
			}
			seen := make(map[string]bool)
			maze[i][j] = "#"
			walk2(maze, Point{x: guard_x, y: guard_y}, &horientation, &counter, &seen)
			maze[i][j] = "."
			horientation = "up"
		}
	}

	println(counter)
}
