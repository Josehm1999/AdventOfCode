package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coordinates struct {
	x int
	y int
}

type NumberAndCoordinates struct {
	value       int
	lenght      int
	coordinates Coordinates
	found       bool
}

func main() {
	file, err := os.Open("../input_day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	sum := 0

	var maze []NumberAndCoordinates
	var stored_lines []string
	for scanner.Scan() {
		line := scanner.Text()
		stored_lines = append(stored_lines, line)
	}

	for i, value := range stored_lines {
		arr_number_and_coordinates := parseNumbersAndCoordinates(value, i)
		maze = append(maze, arr_number_and_coordinates...)
	}

	for s_i, s_value := range stored_lines {
		for index, value := range []rune(s_value) {
			if string(value) == "*" {
				var curr_arr []int
				for _, dx := range []int{0, 1, 2} {
					for _, dy := range []int{0, 1, 2} {

						cx := index + dx - 1
						cy := s_i + dy - 1
						for m_i, curr_num := range maze {
							if curr_num.coordinates.y == cy {
								for index2 := curr_num.coordinates.x; index2 < curr_num.coordinates.x+curr_num.lenght; index2++ {
									if index2 == cx && !curr_num.found {
										curr_num.found = true
										maze[m_i] = curr_num
										// sum += curr_num.value
										curr_arr = append(curr_arr, curr_num.value)
									}
								}
							}
						}
					}
				}
				multiplication_result := 1
				if len(curr_arr) >= 2 {
					for _, val := range curr_arr {
						multiplication_result = multiplication_result * val
					}
					sum += multiplication_result
				}
			}
		}
	}

	// fmt.Println(maze)
	fmt.Println(sum)
	// for i := range maze {
	//
	// 	// When its the first line
	// 	if 0 == i {
	// 		fmt.Println("First Line")
	// 		var curr_number []rune
	//
	// 		for j, _ := range maze[i] {
	// 			if 0 == j {
	// 				fmt.Println("First Character")
	// 			}
	//
	// 			if unicode.IsDigit(maze[i][j]) {
	// 				if checkForSymbol(maze[i][j-1]) || checkForSymbol(maze[i][j+1]) || checkForSymbol(maze[i+1][j]) || checkForSymbol(maze[i+1][j-1]) || checkForSymbol(maze[i+1][j+1]) {
	//
	// 					for left := j - 1; left >= 0; left-- {
	// 						if unicode.IsDigit(maze[i][left]) {
	// 							curr_number = append(curr_number, maze[i][left])
	// 						}
	// 					}
	//
	// 					// number_info := NumberAndCoordinates{
	// 					// 	value:     string(maze[i][j]) + "1",
	// 					// 	yposition: j}
	// 					// curr_number = append(curr_number, number_info)
	// 					fmt.Println("Tiene en algun lado un simbolo")
	// 				}
	// 			}
	//
	// 			// if !unicode.IsLetter(maze[i][j]) && !unicode.IsDigit(maze[i][j]) && "." != string(maze[i][j]) {
	// 			// 	fmt.Println(maze[i][j])
	// 			//
	// 			// 	// Upwards
	// 			// 	if unicode.IsDigit(maze[i-1][j]) {
	// 			// 		fmt.Println(maze[i][j], "Tiene arriba a ", maze[i-1][j])
	// 			//
	// 			//
	// 			// 	}
	// 			// }
	// 			if len(maze[i])-1 == j {
	// 				fmt.Println("Last character")
	// 			}
	// 		}
	// 		fmt.Println(curr_number)
	// 	}

	// 	for j, _ := range maze[i] {
	// 		if 0 == j {
	// 			// if unicode.IsDigit(maze[i][j]) {
	// 			// 	fmt.Println("This is a number")
	// 			// }
	// 			fmt.Println("First Character")
	// 		}
	//
	// 		if len(maze[i])-1 == j {
	// 			fmt.Println("Last character")
	// 		}
	// 	}
	// }
	// }
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

//	func makeRange(min int, max int) []int {
//		a := make([]int, max-min+1)
//		for i := range a {
//			a[i] = min + 1
//		}
//		fmt.Println(a)
//		return a
//	}
func checkForSymbol(s rune) bool {

	if !unicode.IsLetter(s) && !unicode.IsDigit(s) && "." != string(s) {
		return true
	}

	return false
}

func parseNumbersAndCoordinates(s string, counter int) []NumberAndCoordinates {
	runes := []rune(s)
	var arr_number_and_coordinates []NumberAndCoordinates
	left_side := 0
	for left_side < len(runes) {
		right_side := len(runes)
		for right_side > left_side {
			sub_line := runes[left_side:right_side]
			if val, err := strconv.Atoi(string(sub_line)); err == nil {
				if val < 0 {
					val = val * -1
				}

				// fmt.Println(val)
				curr_number := NumberAndCoordinates{
					value:  val,
					lenght: len(sub_line),
					coordinates: Coordinates{
						x: left_side,
						y: counter,
					},
				}
				left_side += len(sub_line)
				arr_number_and_coordinates = append(arr_number_and_coordinates, curr_number)
			}
			right_side = right_side - 1
		}
		left_side += 1
	}
	return arr_number_and_coordinates
}
