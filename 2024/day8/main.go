package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	// "strconv"
	"slices"
	"strings"
)

type Antenna struct {
	row  int
	col  int
	freq string
}

func main() {
	part1()
}

func part1() {

	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	// line := 0

	temp := strings.Split(file, "\n")

	var antenna []Antenna
	cols_start := strings.Split(temp[0], "")
	cols_len := len(cols_start)
	rows_len := len(temp) - 1
	for i := range temp {
		// for i := 0; i < len(temp)-1; i++ {

		cols := strings.Split(temp[i], "")
		for j, v := range cols {

			if v == "." {
				continue
			}
			antenna = append(antenna, Antenna{row: i, col: j, freq: v})
			// nextLine := strings.Split(temp[i+1], "")
			// twoLinesBelow := strings.Split(temp[i+2], "")
			// threeLinesBelow := strings.Split(temp[i+3], "")
			//RIGHT
			// if (len(letters)-1)-j >= 3 {
			// 	// check for diagonal down - right
			// 	for k := 1; k <= 3; k++ {
			// 		if j+k > (len(letters) - 1) {
			// 			continue
			// 		}
			//
			// 		if v == nextLine[j+k] {
			// 			println(i+1, j+k)
			// 			if j+k+1 <= len(letters)-1 {
			// 				println("Check Top R", i+2, j+k+1)
			// 				sum++
			// 			}
			// 			sum++
			// 		}
			// 		if v == twoLinesBelow[j+k] {
			// 			println(i+2, j+k)
			// 			if j+k+1 <= len(letters)-1 {
			// 				println("Check Top R", i+3, j+k+1)
			// 				sum++
			// 			}
			// 			sum++
			// 		}
			// 		if v == threeLinesBelow[j+k] {
			// 			println(i+3, j+k)
			// 			if j+k+1 <= len(letters)-1 {
			// 				println("Check Top R", i+4, j+k+1)
			// 				sum++
			// 			}
			// 			sum++
			// 		}
			// 	}
			// }

			//Left
			// if j >= 3 {
			// 	for k := 1; k <= 3; k++ {
			// 		if j-k < 0 {
			// 			continue
			// 		}
			// 		if v == nextLine[j-k] {
			// 			println("1L", i+1, j-k)
			// 			if j-(2*k) >= 0 && i+1 <= len(temp)-1 {
			// 				sum++
			// 			}
			//
			// 			if j+k+1 <= len(letters)-1 && i-1 >= 0 {
			// 				sum++
			// 			}
			// 		}
			// 		if v == twoLinesBelow[j-k] {
			// 			println("2l", i+2, j-k)
			// 			if j-(2*k) >= 0 && i+2 <= len(temp)-1 {
			// 				// println("Check Bot 2L", i+4, j-(2*k), test[j-(2*k)])
			// 				sum++
			// 			}
			//
			// 			if j+k+1 <= len(letters)-1 && i-2 >= 0 {
			// 				sum++
			// 			}
			// 		}
			// 		if v == threeLinesBelow[j-k] {
			// 			println("3L", i+2, j-k)
			// 			if j-(2*k) >= 0 && i+3 <= len(temp)-1 {
			// 				sum++
			// 			}
			//
			// 			if j+k+1 <= len(letters)-1 && i-3 >= 0 {
			// 				sum++
			// 			}
			// 		}
			// 	}
			// }
		}
	}

	freqCmp := func(a, b Antenna) int {
		return cmp.Compare(a.freq, b.freq)
	}
	slices.SortFunc(antenna, freqCmp)

	var groupedAntennas [][]Antenna
	current_group := []Antenna{antenna[0]}

	for i := 1; i < len(antenna); i++ {
		if antenna[i].freq == antenna[i-1].freq {
			current_group = append(current_group, antenna[i])
		} else {
			groupedAntennas = append(groupedAntennas, current_group)
			current_group = []Antenna{antenna[i]}
		}
	}

	groupedAntennas = append(groupedAntennas, current_group)

	var result []Antenna
	for _, group := range groupedAntennas {
		// fmt.Println(group, cols_len, rows_len)
		result = append(result, product(group, cols_len, rows_len)...)
	}

	fmt.Println(len(removeDuplicate(result)))
}

func product(antennas []Antenna, cols_len int, rows_len int) []Antenna {
	var result []Antenna

	for i := 0; i < len(antennas); i++ {
		for k := i + 1; k < len(antennas); k++ {
			if antennas[k].freq != antennas[i].freq {
				continue
			}

			col_delta := antennas[i].col - antennas[k].col
			row_delta := antennas[i].row - antennas[k].row

			possibleAntenna := Antenna{row: antennas[k].row - row_delta, col: antennas[k].col - col_delta}
			possibleAntenna2 := Antenna{row: antennas[i].row + row_delta, col: antennas[i].col + col_delta}

			// result = append(result, possibleAntenna2)
			// result = append(result, possibleAntenna)
			if possibleAntenna.row >= 0 && possibleAntenna.col >= 0 && possibleAntenna.row <= rows_len && possibleAntenna.col <= cols_len {
				result = append(result, possibleAntenna)
			}
			//
			if possibleAntenna2.row >= 0 && possibleAntenna2.col >= 0 && possibleAntenna2.row <= rows_len && possibleAntenna2.col <= cols_len {
				result = append(result, possibleAntenna2)
			}
			// if col_delta == 3 && row_delta == 1 {
			// 	if antennas[i].row-1 >= 0 && antennas[i].col+3 <= cols_len {
			// 		result = append(result, Antenna{row: antennas[i].row - 1, col: antennas[i].col + 3})
			// 	}
			//
			// 	if antennas[k].row+1 <= rows_len && antennas[k].col-3 >= 0 {
			// 		result = append(result, Antenna{row: antennas[k].row + 1, col: antennas[k].col - 3})
			// 	}
			// }
			//
			// if col_delta == 1 && row_delta == 2 {
			// 	if antennas[i].row-2 >= 0 && antennas[i].col+1 <= cols_len {
			// 		result = append(result, Antenna{row: antennas[i].row - 2, col: antennas[i].col + 1})
			// 	}
			//
			// 	if antennas[k].row+2 <= rows_len && antennas[k].col-1 >= 0 {
			// 		result = append(result, Antenna{row: antennas[k].row + 2, col: antennas[k].col - 1})
			// 	}
			// }
			//
			// if col_delta == 2 && row_delta == 1 {
			// 	if antennas[i].row-1 >= 0 && antennas[i].col+2 <= cols_len {
			// 		result = append(result, Antenna{row: antennas[i].row - 1, col: antennas[i].col + 2})
			// 	}
			//
			// 	if antennas[k].row+1 <= rows_len && antennas[k].col-2 >= 0 {
			// 		result = append(result, Antenna{row: antennas[k].row + 1, col: antennas[k].col - 2})
			// 	}
			// }
			//
			// if col_delta == 2 && row_delta == 3 {
			// 	if antennas[i].row-3 >= 0 && antennas[i].col+2 <= cols_len {
			// 		result = append(result, Antenna{row: antennas[i].row - 3, col: antennas[i].col + 2})
			// 	}
			//
			// 	if antennas[k].row+3 <= rows_len && antennas[k].col-2 >= 0 {
			// 		result = append(result, Antenna{row: antennas[k].row + 3, col: antennas[k].col - 2})
			// 	}
			// }
			//
			// if col_delta == 1 && row_delta == 1 {
			// 	if antennas[i].row-1 >= 0 && antennas[i].col+1 <= cols_len {
			// 		result = append(result, Antenna{row: antennas[i].row - 1, col: antennas[i].col + 1})
			// 	}
			//
			// 	if antennas[k].row+1 <= rows_len && antennas[k].col-1 >= 0 {
			// 		result = append(result, Antenna{row: antennas[k].row + 1, col: antennas[k].col - 1})
			// 	}
			// }
			//
			// if col_delta == 4 && row_delta == 3 {
			// 	if antennas[i].row-3 >= 0 && antennas[i].col+4 <= cols_len {
			// 		result = append(result, Antenna{row: antennas[i].row - 3, col: antennas[i].col + 4})
			// 	}
			//
			// 	if antennas[k].row+3 <= rows_len && antennas[k].col-4 >= 0 {
			// 		result = append(result, Antenna{row: antennas[k].row + 3, col: antennas[k].col - 4})
			// 	}
			// }
		}
	}

	// fmt.Println(result)
	return result
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func removeDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)

	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}

	return list
}
