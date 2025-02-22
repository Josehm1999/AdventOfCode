package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part1(25)
	part2(75)
}

func set_stones(stone_map *map[int]int, stone int) {

	counter := 1
	exists, ok := (*stone_map)[stone]

	if ok {
		counter = exists + 1
	}

	(*stone_map)[stone] = counter

}

func exists_stone(stone_map *map[int]int, stone int) int {
	exists, ok := (*stone_map)[stone]

	if ok {
		return exists
	}
	return 0
}

func get_digits(num int) int {
	counter := 0
	if num == 0 {
		return 1
	}
	for num > 0 {
		counter += 1
		num /= 10
	}
	return counter
}

func part1(num_iter int) {

	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	temp := strings.Split(file, "\n")

	num_arr := strings.Split(temp[0], " ")

	for i := 0; i < num_iter; i++ {
		j := 0
		for j < len(num_arr) {

			if num_arr[j] == "0" {
				num_arr[j] = "1"
				j++
			} else if len(num_arr[j])%2 == 0 {
				halfway := len(num_arr[j]) / 2
				rStone, _ := strconv.Atoi(num_arr[j][:halfway])
				lStore, _ := strconv.Atoi(num_arr[j][halfway:])
				num_arr[j] = strconv.Itoa(lStore)
				// fmt.Println("Tiene un numero par de caracteres", num_arr[j], halfway, left_stone, right_stone)
				// test := append([]string{right_stone}, num_arr[j:]...)
				// num_arr = append(num_arr[:j], test...)

				num_arr = slices.Insert(num_arr, j+1, strconv.Itoa(rStone))
				j += 2
			} else {
				num_int, _ := strconv.Atoi(num_arr[j])
				num_arr[j] = strconv.Itoa(num_int * 2024)
				j++
			}
			// fmt.Println(i, j)
		}
	}

	fmt.Println(len(num_arr))
}

func evalate_stones(stones_map map[int]int) map[int]int {

	result := make(map[int]int)
	for key, val := range stones_map {
		// fmt.Println(key, val)
		if key == 0 {
			new_stone := 1
			result[new_stone] = exists_stone(&result, new_stone) + val
		} else {
			num_digits := get_digits(key)
			if num_digits%2 == 0 {

				half_len := num_digits / 2
				padding := int(math.Pow(10, float64(half_len)))

				lower_part := key / padding
				high_part := key % padding

				// fmt.Println("test", key, half_len, padding, high_part, lower_part)
				lower_count := exists_stone(&result, lower_part) + val
				result[lower_part] = lower_count

				higher_count := exists_stone(&result, high_part) + val
				result[high_part] = higher_count

			} else {
				new_stone := key * 2024
				result[new_stone] = exists_stone(&result, new_stone) + val
			}
		}
	}
	return result
}

func part2(num_iter int) {
	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	temp := strings.Split(file, "\n")

	string_arr := strings.Split(temp[0], " ")
	var num_arr []int
	for _, val := range string_arr {
		num, _ := strconv.Atoi(val)
		num_arr = append(num_arr, num)
	}

	result := make(map[int]int)
	for _, val := range num_arr {
		set_stones(&result, val)
	}

	for i := 0; i < num_iter; i++ {
		result = evalate_stones(result)
	}
	total := 0
	for _, val := range result {
		total += val
	}

	fmt.Println(total)
}
