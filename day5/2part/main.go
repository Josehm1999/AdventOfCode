package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type SeedValues struct {
	src    int
	dest   int
	length int
}

func main() {
	file, err := os.ReadFile("../input_day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	contents := strings.Split(string(file[:]), "\n\n")
	seeds := strings.Split(strings.Split(contents[0], ": ")[1], " ")
	// fmt.Println(seeds)
	_, new_content := contents[0], contents[1:]
	// fmt.Println(new_content[6])
	for _, curr_seed := range seeds {
		// curr_seed_int, _ := strconv.Atoi(curr_seed)
		fmt.Println(curr_seed)
		for i := 0; i < len(new_content); i++ {
			// fmt.Println(new_content[i])
			for _, val := range parseMap(new_content[i]) {
				fmt.Println(val)
				// 	// fmt.Println(curr_seed, val)
				// 	if curr_seed_int >= val.src && curr_seed_int <= val.src+val.length-1 {
				// 		curr_seed_int = curr_seed_int + val.dest - val.src
				// 		fmt.Println(curr_seed, curr_seed_int)
				// 	} else {
				//
				// 		fmt.Println(curr_seed, curr_seed_int)
				// 	}
			}
		}
	}
}

func createRange(line string) SeedValues {
	items := strings.Split(line, " ")
	dest, _ := strconv.Atoi(items[0])
	src, _ := strconv.Atoi(items[1])
	length, _ := strconv.Atoi(items[2])
	seedValue := SeedValues{src, dest, length}
	return seedValue
}

func parseMap(data string) []SeedValues {
    fmt.Println(data)
	var seed_values_arr []SeedValues
	lines := strings.Split(data, "\n")
	src_to_dest := strings.Split(strings.Split(lines[0], " ")[0], "-")
	_, new_line := lines[0], lines[1:]
	src := src_to_dest[0]
	dest := src_to_dest[2]

	for _, val := range new_line {
		seed_values_arr = append(seed_values_arr, createRange(val))
	}
	fmt.Println(src, dest)
	return seed_values_arr
}

// func loopArraryOfArrays(wn_arr_of_arr [][]int, my_arr_arr [][]int, sum int) int {
// 	curr_matches := 0
// 	// fmt.Println("Nuevo bucle1")
// 	for i, wm_v := range wn_arr_of_arr {
// 		for j, my_v := range my_arr_arr {
// 			if i == j {
// 				curr_matches = findMatches(wm_v, my_v)
// 			}
// 		}
//
// 		if curr_matches > 0 {
// 			new_wn_arr_of_arr := wn_arr_of_arr[i+1 : i+curr_matches+1]
// 			new_my_arr_of_arr := my_arr_arr[i+1 : i+curr_matches+1]
// 			sum += loopArraryOfArrays(new_wn_arr_of_arr, new_my_arr_of_arr, curr_matches)
// 		}
// 	}
// 	return sum
// }
//
// func findMatches(wn_arr []int, my_arr []int) int {
// 	num_ocurr := 0
// 	for _, v := range wn_arr {
// 		for _, my_v := range my_arr {
// 			if v == my_v {
// 				num_ocurr++
// 			}
// 		}
// 	}
// 	return num_ocurr
// }
//
// func extractNumbers(s string) ([]int, []int) {
// 	// runes := []rune(s)
// 	var wn []int
// 	var myn []int
//
// 	split_at_colon := strings.Split(s, ":")
// 	// fmt.Println(split_at_colon[1])
//
// 	split_at_vertical_bar := strings.Split(split_at_colon[1], "|")
// 	split_at_space_wn := strings.Split(strings.Trim(split_at_vertical_bar[0], " "), " ")
// 	split_at_space_myn := strings.Split(strings.Trim(split_at_vertical_bar[1], " "), " ")
//
// 	for _, v := range split_at_space_wn {
// 		curr_number, err := strconv.Atoi(v)
// 		if err == nil {
// 			wn = append(wn, curr_number)
// 		}
// 	}
//
// 	for _, v := range split_at_space_myn {
// 		curr_number, err := strconv.Atoi(v)
// 		if err == nil {
// 			myn = append(myn, curr_number)
// 		}
// 	}
// 	return wn, myn
// }
