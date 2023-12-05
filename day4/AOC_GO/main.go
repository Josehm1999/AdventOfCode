package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
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
	scanner := bufio.NewScanner(file)
	test, _ := lineCounter(file)
	for scanner.Scan() {
		counter := 1
		line := scanner.Text()
		extractSymbols(line)
		counter++
        fmt.Println(counter)
	}
	fmt.Println("%d lines, %d symbols\n", counter, test)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func extractIds(s string) int {
	position_of_separator := strings.Index(s, ":")
	current_id, err := strconv.Atoi(s[5:position_of_separator])
	if err != nil {
		fmt.Println("Cannot extract Id")
	}
	return current_id
}

type cube_data struct {
	index  int
	color  string
	number int
}

func extractsNumOfCubesByColor(s string) []cube_data {
	position_of_separator := strings.Index(s, ":")
	right_side_with_cubes := string(s[position_of_separator+1:])
	var cubes_by_line []cube_data
	splitted_by_dotted_comma := strings.Split(right_side_with_cubes, ";")
	for index, value_dotted_coma := range splitted_by_dotted_comma {
		for _, value_comma := range strings.Split(value_dotted_coma, ",") {

			trimmed_value := strings.Trim(value_comma, " ")
			number_value_cube, _ := strconv.Atoi(strings.Split(trimmed_value, " ")[0])
			cube_info := cube_data{
				index:  index,
				color:  strings.Split(trimmed_value, " ")[1],
				number: number_value_cube,
			}

			cubes_by_line = append(cubes_by_line, cube_info)
		}
	}

	return cubes_by_line
}

func extractNumbers(s string) {
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			number, _ := strconv.Atoi(string(runes[i]))
			fmt.Println(number)
		}
	}
}

func extractSymbols(s string) {
	runes := []rune(s)
	character_to_avoid := "."
	for i := 0; i < len(runes); i++ {
		if !unicode.IsLetter(runes[i]) && !unicode.IsDigit(runes[i]) && character_to_avoid != string(runes[i]) {
			symbol := string(runes[i])
			fmt.Println(symbol)
		}
	}
}

func getLastLineWithSeek(filepath string) string {
	fileHandle, err := os.Open(filepath)

	if err != nil {
		panic("Cannot open file")
	}
	defer fileHandle.Close()

	line := ""
	var cursor int64 = 0
	stat, _ := fileHandle.Stat()
	filesize := stat.Size()
	for {
		cursor -= 1
		fileHandle.Seek(cursor, io.SeekEnd)

		char := make([]byte, 1)
		fileHandle.Read(char)

		if cursor != -1 && (char[0] == 10 || char[0] == 13) { // stop if we find a line
			break
		}

		line = fmt.Sprintf("%s%s", string(char), line) // there is more efficient way

		if cursor == -filesize { // stop if we are at the begining
			break
		}
	}

	return line
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
