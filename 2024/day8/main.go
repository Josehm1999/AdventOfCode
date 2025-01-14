package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"

	// "strconv"

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

	println(cols_len, rows_len)

	//Find all antennas
	for i := range temp {
		cols := strings.Split(temp[i], "")
		for j, v := range cols {
			if v == "." {
				continue
			}
			antenna = append(antenna, Antenna{row: i, col: j, freq: v})
		}
	}

	//Find the combinantions of the antennas

	freqCmp := func(a, b Antenna) int {
		return cmp.Compare(a.freq, b.freq)
	}
	slices.SortFunc(antenna, freqCmp)

	//Check the antinodes to be found

	output := make(map[string][]Antenna)
	var groupedAntennas [][]Antenna
	current_group := []Antenna{antenna[0]}
	output[antenna[0].freq] = current_group

	// for _, value := range antenna {
	// 	val, ok := output[value.freq]
	// 	if ok {
	// 		val = append(val, value)
	// 	} else {
	// 		groupedAntennas = append(groupedAntennas, output[value.freq])
	// 	}
	// }
	// var groupedAntennas [][]Antenna
	// current_group := []Antenna{antenna[0]}
	// counter := 0
	//
	for i := 1; i < len(antenna); i++ {
		if antenna[i].freq == antenna[i-1].freq {
			current_group = append(current_group, antenna[i])
			if i == len(antenna)-1 {
				groupedAntennas = append(groupedAntennas, current_group)
			}
		} else {
			groupedAntennas = append(groupedAntennas, current_group)
			current_group = []Antenna{antenna[i]}
		}
	}
	// fmt.Println(groupedAntennas)
	// fmt.Println((antenna))
	//Remove duplicates

	//Count antinodes

	//
	// // fmt.Println(groupedAntennas)
	// groupedAntennas = append(groupedAntennas, current_group)
	//
	var result []Antenna
	for _, group := range groupedAntennas {
		// fmt.Println(group, cols_len, rows_len)
		result = append(result, product(group, cols_len, rows_len)...)
	}
	//
	fmt.Println(len(removeDuplicate(result)))

	// new_result := removeDuplicate(result)

}

func product(antennas []Antenna, cols_len int, rows_len int) []Antenna {
	var result []Antenna
	// seen := make(map[Antenna]int)
	for i := 0; i < len(antennas); i++ {
		for k := i + 1; k < len(antennas); k++ {
			if antennas[k].freq != antennas[i].freq {
				continue
			}

			// fmt.Println(antennas[i], antennas[k])
			// col_delta := antennas[i].col - antennas[k].col
			// row_delta := antennas[i].row - antennas[k].row
			// println(col_delta, row_delta)
			findPossibleAntennas2(antennas[i], antennas[k], rows_len, cols_len, &result)

			// canTraverse := true
			// counter := 0
			// for canTraverse {
			// counter++
			// fmt.Println("New Antenna")
			// fmt.Println(antennas[k], antennas[i], row_delta*counter)
			// possibleAntenna := Antenna{row: antennas[k].row - row_delta, col: antennas[k].col - col_delta}
			// possibleAntenna2 := Antenna{row: antennas[i].row + row_delta, col: antennas[i].col + col_delta}
			//
			// // canTraverse = false
			// if possibleAntenna.row >= 0 && possibleAntenna.col >= 0 && possibleAntenna.row <= rows_len && possibleAntenna.col <= cols_len {
			// 	// canTraverse = true
			// 	result = append(result, possibleAntenna)
			// }
			//
			// if possibleAntenna2.row >= 0 && possibleAntenna2.col >= 0 && possibleAntenna2.row <= rows_len && possibleAntenna2.col <= cols_len {
			// 	// canTraverse = true
			// 	result = append(result, possibleAntenna2)
			// }

			// fmt.Println(canTraverse, possibleAntenna, possibleAntenna2)
			// if canTraverse {
			// 	if seen[antennas[k]] >= 2 {
			// 		result = append(result, antennas[k])
			// 	}
			//
			// 	if seen[antennas[i]] >= 2 {
			// 		result = append(result, antennas[i])
			// 	}
			//
			// 	seen[antennas[k]]++
			// 	seen[antennas[i]]++
			// }

			// }
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
		key := fmt.Sprintf("%v,%v", item.col, item.row)
		// fmt.Println(key)
		if _, value := allKeys[key]; !value {
			allKeys[key] = true
			list = append(list, item)
		}
	}

	return list
}

func findPossibleAntennas(firstAntenna Antenna, secondAntenna Antenna, col_delta int, row_delta int, row_limit int, col_limit int, result *[]Antenna) bool {

	possibleAntenna := Antenna{row: secondAntenna.row - row_delta, col: secondAntenna.col - col_delta, freq: secondAntenna.freq}
	possibleAntenna2 := Antenna{row: firstAntenna.row + row_delta, col: firstAntenna.col + col_delta, freq: firstAntenna.freq}

	fmt.Println(firstAntenna, possibleAntenna)
	fmt.Println(secondAntenna)
	fmt.Println(row_delta, col_delta)

	canTraverse := false
	if possibleAntenna.row >= 0 && possibleAntenna.col >= 0 && possibleAntenna.row <= row_limit && possibleAntenna.col <= col_limit {
		canTraverse = true
		*result = append(*result, possibleAntenna)
	}

	if possibleAntenna2.row >= 0 && possibleAntenna2.col >= 0 && possibleAntenna2.row <= row_limit && possibleAntenna2.col <= col_limit {
		canTraverse = true
		*result = append(*result, possibleAntenna2)
	}

	fmt.Println(canTraverse)

	if canTraverse {
		*result = append(*result, firstAntenna)
		*result = append(*result, secondAntenna)
		findPossibleAntennas(possibleAntenna, possibleAntenna2, col_delta, row_delta, row_limit, col_limit, result)
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

	// fmt.Println(possibleAntenna)
	// println(col_delta*+1, row_delta*1, col_delta, row_delta)

	// if findPossibleAntennas(firstAntenna, secondAntenna, col_delta*2, row_delta*2, row_limit, col_limit, result) {
	// 	return true
	// }

	return false
}

func resonance(antenna Antenna, row_delta int, col_delta int, row_limit int, col_limit int) []Antenna {

	var antinodes []Antenna
	// Si la antenna pasa los bordes del maze
	for antenna.col+col_delta >= 0 && antenna.col+col_delta <= col_limit && antenna.row+row_delta >= 0 && antenna.row+row_delta <= row_limit {
		antinodes = append(antinodes, Antenna{col: col_delta + antenna.col, row: row_delta + antenna.row, freq: antenna.freq})
		antenna.col = antenna.col + col_delta
		antenna.row = antenna.row + row_delta
	}
	return antinodes
}

func findPossibleAntennas2(firstAntenna Antenna, secondAntenna Antenna, col_limit int, row_limit int, result *[]Antenna) {

	col_delta := firstAntenna.col - secondAntenna.col
	row_delta := firstAntenna.row - secondAntenna.row

	possibleAntenna := Antenna{col: firstAntenna.col - col_delta, row: firstAntenna.row - row_delta, freq: firstAntenna.freq}
	possibleAntenna2 := Antenna{col: secondAntenna.col + col_delta, row: secondAntenna.row + row_delta, freq: secondAntenna.freq}

	// fmt.Println(resonance(possibleAntenna, row_delta, col_delta, row_limit, col_limit))
	// fmt.Println(resonance(possibleAntenna2, -row_delta, -col_delta, row_limit, col_limit))

	antinodes1 := resonance(possibleAntenna, row_delta, col_delta, row_limit, col_limit)
	antinodes2 := resonance(possibleAntenna2, -row_delta, -col_delta, row_limit, col_limit)

	*result = append(*result, antinodes1...)
	*result = append(*result, antinodes2...)
}
