package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

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

	values_arr := strings.Split(temp[0], "")
	fmt.Println(slidingWindow(5, values_arr))

}

func slidingWindow(size int, input []string) [][]string {
	if len(input) <= size {
		return [][]string{input}
	}

	r := make([][]string, 0, len(input)-size+1)

	for i, j := 0, size; j <= len(input); i, j = i+1, j+1 {
		r = append(r, input[i:j])
	}

	return r
}
