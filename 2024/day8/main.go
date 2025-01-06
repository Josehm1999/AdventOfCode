package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	value    int
	operands []int
}

func main() {
	part1()
}

func walk(equation_elements []Equation) int {
	counter := 0
	for i := 0; i < len(equation_elements); i++ {
		is_valid := walk_recursive(equation_elements[i].operands, 0, equation_elements[i].operands[0], equation_elements[i].value)

		if is_valid {
			counter += equation_elements[i].value
			println(counter)
		}

	}
	return counter
}

func walk_recursive(equation_elements []int, curr int, current_result int, result int) bool {

	// println(len(equation_elements), curr, current_result, result, current_result == result)
	if current_result == result {
		return true
	}

	if len(equation_elements)-1 <= curr {
		return false
	}

	// Concatenation is part 2
	concat_part1 := strconv.Itoa(current_result)
	concat_part2 := strconv.Itoa(equation_elements[curr+1])
	concatenation_val, _ := strconv.Atoi(concat_part1 + concat_part2)

	sum := walk_recursive(equation_elements, curr+1, current_result+equation_elements[curr+1], result)
	mul := walk_recursive(equation_elements, curr+1, current_result*equation_elements[curr+1], result)
	concatenation := walk_recursive(equation_elements, curr+1, concatenation_val, result)

	return sum || mul || concatenation
}

func part1() {

	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	// line := 0

	temp := strings.Split(file, "\n")

	sum := 0
	seen := make(map[string]bool)

	for i := 0; i < len(temp)-1; i++ {
		// fmt.Println(temp[i])
		isPossibleDown := false

		if (len(temp)-2)-i > 2 {
			isPossibleDown = true
		}

		letters := strings.Split(temp[i], "")
		// println(isPossibleUp, isPossibleDown, sum)
		for j, v := range letters {
			key := fmt.Sprintf("%d,%d", j, i)
			seen[key] = true

			if v == "." {
				continue
			}

			println("New line", v)
			if isPossibleDown {

				nextLine := strings.Split(temp[i+1], "")
				twoLinesBelow := strings.Split(temp[i+2], "")
				threeLinesBelow := strings.Split(temp[i+3], "")
				//RIGHT
				if (len(letters)-1)-j >= 3 {
					// check for diagonal down - right
					for k := 1; k <= 3; k++ {
						if j+k > (len(letters) - 1) {
							continue
						}

						if v == nextLine[j+k] {
							println(i+1, j+k)
							if j+k+1 <= len(letters)-1 {
								println("Check Top R", i+2, j+k+1)
								sum++
							}
							sum++
						}
						if v == twoLinesBelow[j+k] {
							println(i+2, j+k)
							if j+k+1 <= len(letters)-1 {
								println("Check Top R", i+3, j+k+1)
								sum++
							}
							sum++
						}
						if v == threeLinesBelow[j+k] {
							println(i+3, j+k)
							if j+k+1 <= len(letters)-1 {
								println("Check Top R", i+4, j+k+1)
								sum++
							}
							sum++
						}
					}
				}

				//Left
				if j >= 3 {
					for k := 1; k <= 3; k++ {
						if j-k < 0 {
							continue
						}
						if v == nextLine[j-k] {
							println("1L", i+1, j-k)
							if j-(2*k) >= 0 && i+1 <= len(temp)-1 {
								sum++
							}

							if j+k+1 <= len(letters)-1 && i-1 >= 0 {
								sum++
							}
						}
						if v == twoLinesBelow[j-k] {
							println("2l", i+2, j-k)
							if j-(2*k) >= 0 && i+2 <= len(temp)-1 {
								// println("Check Bot 2L", i+4, j-(2*k), test[j-(2*k)])
								sum++
							}

							if j+k+1 <= len(letters)-1 && i-2 >= 0 {
								sum++
							}
						}
						if v == threeLinesBelow[j-k] {
							println("3L", i+2, j-k)
							if j-(2*k) >= 0 && i+3 <= len(temp)-1 {
								sum++
							}

							if j+k+1 <= len(letters)-1 && i-3 >= 0 {
								sum++
							}
						}
					}
				}
			}

			// if isPossibleUp {
			// 	possibleDiagLine := strings.Split(temp[i-2], "")
			// 	secondPossibleDiagLine := strings.Split(temp[i-1], "")
			// 	if j >= 3 {
			// 		if v == possibleDiagLine[j+1] {
			// 			if (i-4 < len(temp)-2) && (j+2) < len(letters)-1 {
			// 				println("Possible new antinode")
			//                          sum++
			// 			}
			// 		}
			//
			// 		if v == secondPossibleDiagLine[j+2] {
			//
			// 			if (i-3 < len(temp)-2) && (j+3) < len(letters)-1 {
			// 				println("Possible new antinode 2")
			//                          sum++
			// 			}
			// 		}
			// 	}
			//
			// 	// check for diagonal down - left
			// 	if (len(letters)-1)-j >= 3 {
			//
			// 		if v == possibleDiagLine[j-1] {
			// 			if (i-4 < len(temp)-2) && (j-2) < len(letters)-1 {
			// 				println("Possible new antinode")
			//                          sum++
			// 			}
			// 		}
			//
			// 		if v == secondPossibleDiagLine[j-2] {
			//
			// 			if (i-3 < len(temp)-2) && (j-3) < len(letters)-1 {
			// 				println("Possible new antinode 2")
			//                          sum++
			// 			}
			// 		}
			// 	}
			// }
			// println(v, seen[key])
		}

	}
	println(sum)
}
