package main

import (
	"bufio"
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

func main() {
	file, err := os.Open("../input_day4.txt")
	if err != nil {
		log.Fatal(err)
	}

	// sum := 0
	scanner := bufio.NewScanner(file)
	var winning_numbers_arr [][]int
	var my_numbers_arr [][]int
	for scanner.Scan() {
		line := scanner.Text()
		winning_numbers, my_numbers := extractNumbers(line)
		winning_numbers_arr = append(winning_numbers_arr, winning_numbers)
		my_numbers_arr = append(my_numbers_arr, my_numbers)

	}

	sum := 0
	sum += loopArraryOfArrays(winning_numbers_arr, my_numbers_arr, sum)
	fmt.Println(sum + len(winning_numbers_arr))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func loopArraryOfArrays(wn_arr_of_arr [][]int, my_arr_arr [][]int, sum int) int {
	curr_matches := 0
	// fmt.Println("Nuevo bucle1")
	for i, wm_v := range wn_arr_of_arr {
		for j, my_v := range my_arr_arr {
			if i == j {
				curr_matches = findMatches(wm_v, my_v)
			}
		}

		if curr_matches > 0 {
			new_wn_arr_of_arr := wn_arr_of_arr[i+1 : i+curr_matches+1]
			new_my_arr_of_arr := my_arr_arr[i+1 : i+curr_matches+1]
			sum += loopArraryOfArrays(new_wn_arr_of_arr, new_my_arr_of_arr, curr_matches)
		}
	}
	return sum
}

func findMatches(wn_arr []int, my_arr []int) int {
	num_ocurr := 0
	for _, v := range wn_arr {
		for _, my_v := range my_arr {
			if v == my_v {
				num_ocurr++
			}
		}
	}
	return num_ocurr
}

func extractNumbers(s string) ([]int, []int) {
	// runes := []rune(s)
	var wn []int
	var myn []int

	split_at_colon := strings.Split(s, ":")
	// fmt.Println(split_at_colon[1])

	split_at_vertical_bar := strings.Split(split_at_colon[1], "|")
	split_at_space_wn := strings.Split(strings.Trim(split_at_vertical_bar[0], " "), " ")
	split_at_space_myn := strings.Split(strings.Trim(split_at_vertical_bar[1], " "), " ")

	for _, v := range split_at_space_wn {
		curr_number, err := strconv.Atoi(v)
		if err == nil {
			wn = append(wn, curr_number)
		}
	}

	for _, v := range split_at_space_myn {
		curr_number, err := strconv.Atoi(v)
		if err == nil {
			myn = append(myn, curr_number)
		}
	}
	return wn, myn
}
