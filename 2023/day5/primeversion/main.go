package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SeedRanges struct {
	start int
	end   int
}

type SeedValues struct {
	src    int
	dest   int
	length int
}

type Ranges struct {
	from   string
	to     string
	maping []SeedValues
}

func main() {
	file, err := os.ReadFile("../input_day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	contents := strings.Split(strings.ReplaceAll(string(file[:]), "\r\n", "\n"), "\n\n")
	seeds := strings.Split(strings.Split(contents[0], ": ")[1], " ")
	var seed_ranges_arr []SeedRanges
	for i := 0; i < len(seeds); i = i + 2 {
		start, _ := strconv.Atoi(seeds[i])
		len, _ := strconv.Atoi(seeds[i+1])
		end := start + len
		seed_ranges_arr = append(seed_ranges_arr, SeedRanges{start, end})
	}

	_, new_content := contents[0], contents[1:]
	for _, value := range new_content {
		curr_range_line := parseMap(value)
		curr_range_line.maping = createNegativeRanges(curr_range_line.maping)
		fmt.Println(curr_range_line)
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

func createNegativeRanges(ranges []SeedValues) []SeedValues {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].src < ranges[j].src
	})

	start := 0

	for i := 0; i < len(ranges); i++ {
		if ranges[i].src > start {
			ranges = insertValueAtIndex(ranges, i, SeedValues{
				src:    start,
				dest:   start,
				length: ranges[i].src - start,
			})
			i++
		}
		start = ranges[i].src + ranges[i].length
	}
	return ranges
}

func parseMap(data string) Ranges {
	var seed_values_arr []SeedValues
	var ranges Ranges

	lines := strings.Split(data, "\n")
	src_to_dest := strings.Split(strings.Split(lines[0], " ")[0], "-")
	from := src_to_dest[0]
	to := src_to_dest[2]
	_, new_line := lines[0], lines[1:]
	for _, val := range new_line {
		if val != "" {
			seed_values_arr = append(seed_values_arr, createRange(val))
		}
	}
	ranges = Ranges{from, to, seed_values_arr}
	return ranges
}

func insertValueAtIndex(arr []SeedValues, index int, value SeedValues) []SeedValues {
	if index < 0 || index > len(arr) {
		fmt.Println("Intex out of bounds")
		return arr
	}
	newArr := make([]SeedValues, len(arr)+1)

	copy(newArr[:index], arr[:index])

	newArr[index] = value

	copy(newArr[index+1:], arr[index:])

	return newArr
}
