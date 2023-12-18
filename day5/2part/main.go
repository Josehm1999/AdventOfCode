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
	fmt.Println(seeds)

	for _, value := range seeds {
		for i := 1; i < len(contents); i++ {
			fmt.Println(value)
			parseMap(contents[i])
		}
	}
	// contents := strings.Split(file, "\n\n")
	// sum := 0
	// for scanner.Scan() {
	// 	line := scanner.Text()
	//
	// }

}

func createRange(line string) SeedValues {
	items := strings.Split(line, " ")
	dest, _ := strconv.Atoi(items[0])
	src, _ := strconv.Atoi(items[0])
	length, _ := strconv.Atoi(items[0])

	seedValue := SeedValues{src, dest, length}
	return seedValue
}

func parseMap(data string) {
	// fmt.Println(data)
    // var seed_values_arr =
	lines := strings.Split(data, "\n")
	src_to_dest := strings.Split(strings.Split(lines[0], " ")[0], "-")

	src := src_to_dest[2]
	dest := src_to_dest[0]
	for i := 1; i < len(lines); i++ {
	}
	ranges := createRange(lines[])
	fmt.Println(lines[0])
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
