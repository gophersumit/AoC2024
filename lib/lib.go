package lib

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func GetNSlices(lines []string, nums int) [][]int {
	numberSlices := make([][]int, len(lines))
	for _, line := range lines {
		numbers := strings.Fields(line)
		if len(numbers) != nums {
			panic("invalid number of elements in line")
		}
		var n = make([]int, len(numbers))
		for _, num := range numbers {
			v, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			n = append(n, v)
		}
		numberSlices = append(numberSlices, n)
	}

	return numberSlices
}

func GetNSlicesFromFile(filename string, num int) [][]int {
	lines := GetLines(filename)
	numbers := GetNSlices(lines, num)
	return numbers
}
