package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
}

type Point struct {
	x int
	y int
}

type Horientation struct {
	horientation string
	check        bool
	counter      int
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

	result := walk(maze, Point{x: guard_x, y: guard_y}, &horientation, &counter, &seen)

	println(result)
	println(counter)
	//Base case
	//1.- Its an obstacle
	//2.- Off the map end
	//3.- Have we been here
	//4.- Each new tile needs to sum to a counter

	//Recurse
	//1.- Only go forward
	//2.- Keep track of horientation
	//3.- If hits a wall go right depending horientation

	//Pre
	//Recurse

	// if (walk)
	// walk()
	//Post
	// println(current_guard_position, temp)
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

func part2() {

	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	// line := 0

	temp := strings.Split(file, "\n\n")

	rulesArr := strings.Split(temp[0], "\n")
	pagesToUpdate := strings.Split(temp[1], "\n")

	sum := 0
	// rules := make(map[int]int)

	// orderedPagesArr := make([]map[string]int, len(pagesToUpdate)-1)

	for i := 0; i < len(pagesToUpdate)-1; i++ {
		numArr := strings.Split(pagesToUpdate[i], ",")
		orderedPages := make(map[string]int)

		// currentMiddleNum := ""
		for j, numToUpdate := range numArr {
			orderedPages[numToUpdate] = j
			// if (len(numArr)-1)/2 == j {
			// 	currentMiddleNum = numToUpdate
			// }
		}

		isValid := true
		for _, v := range rulesArr {
			rulesLevels := strings.Split(v, "|")
			lowerLevel := rulesLevels[0]
			upperLevel := rulesLevels[1]

			// println(lowerLevel, upperLevel)
			lowerBoundIndex, existsLower := orderedPages[lowerLevel]
			upperBoundIndex, existsUpper := orderedPages[upperLevel]

			if !existsLower || !existsUpper {
				continue
			}

			if lowerBoundIndex > upperBoundIndex {
				isValid = false
			}

		}

		// if isValid {
		// 	value, _ := strconv.Atoi(currentMiddleNum)
		// 	sum += value
		// }

		if !isValid {

			// println("New Line")
			for _, v := range rulesArr {
				rulesLevels := strings.Split(v, "|")
				lowerLevel := rulesLevels[0]
				upperLevel := rulesLevels[1]

				lowerBoundIndex, existsLower := orderedPages[lowerLevel]
				upperBoundIndex, existsUpper := orderedPages[upperLevel]

				if !existsLower || !existsUpper {
					continue
				}

				if lowerBoundIndex > upperBoundIndex {
					orderedPages[lowerLevel] = upperBoundIndex
					orderedPages[upperLevel] = lowerBoundIndex
					numArr[upperBoundIndex] = lowerLevel
					numArr[lowerBoundIndex] = upperLevel
				}

				// println(lowerBoundIndex, lowerLevel)
				// println(upperBoundIndex, upperLevel)
				//
				// println("------")
			}

			test := numArr[(len(numArr)-1)/2]

			// for _, v := range numArr {
			// 	println(v)
			// }
			// println(test)
			value, _ := strconv.Atoi(test)
			// println(test)
			sum += value
		}
	}

	println(sum)
}
