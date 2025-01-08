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

	temp := strings.Split(file, "\n")

	var antenna []Antenna
	col_test := strings.Split(temp[0], "")
	cols_len := len(col_test) - 1
	rows_len := len(temp) - 2

	// println(cols_len, rows_len)
	for i := range temp {
		cols := strings.Split(temp[i], "")
		for j, v := range cols {

			if v == "." {
				continue
			}
			antenna = append(antenna, Antenna{row: i, col: j, freq: v})
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
			// println(col_delta, row_delta)
			// findPossibleAntennas(antennas[i], antennas[k], col_delta, row_delta, rows_len, cols_len, &result)
			//
			canTraverse := true
			counter := 0
			for canTraverse {
				counter++

				possibleAntenna := Antenna{row: antennas[k].row - row_delta*counter, col: antennas[k].col - col_delta*counter}
				possibleAntenna2 := Antenna{row: antennas[i].row + row_delta*counter, col: antennas[i].col + col_delta*counter}

				canTraverse = false
				if possibleAntenna.row >= 0 && possibleAntenna.col >= 0 && possibleAntenna.row <= rows_len && possibleAntenna.col <= cols_len {
					canTraverse = true
					result = append(result, possibleAntenna)
				}
				// //
				if possibleAntenna2.row >= 0 && possibleAntenna2.col >= 0 && possibleAntenna2.row <= rows_len && possibleAntenna2.col <= cols_len {
					canTraverse = true
					result = append(result, possibleAntenna2)
				}

				if canTraverse {
					result = append(result, antennas[k])
					result = append(result, antennas[i])
				}

				// fmt.Println(canTraverse, possibleAntenna, possibleAntenna2)
			}
			//
			// // result = append(result, possibleAntenna2)
			// // result = append(result, possibleAntenna)
		}
	}

	return result
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func removeDuplicate(sliceList []Antenna) []Antenna {
	allKeys := make(map[string]bool)

	list := []Antenna{}
	for _, item := range sliceList {
		key := fmt.Sprintf("%d,%d", item.col, item.row)
		if _, value := allKeys[key]; !value {
			allKeys[key] = true
			list = append(list, item)
		}
	}

	return list
}

func findPossibleAntennas(firstAntenna Antenna, secondAntenna Antenna, col_delta int, row_delta int, row_limit int, col_limit int, result *[]Antenna) bool {

	// Pre

	if row_delta == 0 && col_delta == 0 {
		return true
	}
	//    fmt.Println(firstAntenna)
	// fmt.Println(secondAntenna)
	possibleAntenna := Antenna{row: secondAntenna.row - row_delta, col: secondAntenna.col - col_delta, freq: secondAntenna.freq}
	possibleAntenna2 := Antenna{row: firstAntenna.row + row_delta, col: firstAntenna.col + col_delta, freq: firstAntenna.freq}

	// fmt.Println(possibleAntenna)
	// fmt.Println(possibleAntenna2)

	canTraverse := false
	if possibleAntenna.row >= 0 && possibleAntenna.col >= 0 && possibleAntenna.row <= row_limit && possibleAntenna.col <= col_limit {
		canTraverse = true
		*result = append(*result, possibleAntenna)
	}

	if possibleAntenna2.row >= 0 && possibleAntenna2.col >= 0 && possibleAntenna2.row <= row_limit && possibleAntenna2.col <= col_limit {
		canTraverse = true
		*result = append(*result, possibleAntenna2)
	}
	// Recurse

	// test_col := 0
	// test_row := 0
	// if col_delta == 1 {
	// 	test_col = col_delta + 1
	// } else if col_delta == -1 {
	// 	test_col = col_delta - 1
	// } else {
	// 	test_col = col_delta * 2
	// }
	//
	// if row_delta == 1 {
	// 	test_row = row_delta + 1
	// } else if test_row == -1 {
	// 	test_row = row_delta - 1
	// } else {
	// 	test_row = row_delta * 2
	// }

	fmt.Println(possibleAntenna)
	println(col_delta*+1, row_delta*1, col_delta, row_delta)
	if canTraverse {
		findPossibleAntennas(firstAntenna, secondAntenna, col_delta+1, row_delta+1, row_limit, col_limit, result)
	}

	return false
}
