package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	part1()

	part2()

}

func part2() {
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	file1, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()
	sampleSimilarity := calculateSimilarity(file)
	fmt.Println("sample similarity:", sampleSimilarity)

	similarity := calculateSimilarity(file1)
	fmt.Println("similarity:", similarity)
}
func part1() {
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	file1, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()
	sampleDist := calculateDistance(file)
	fmt.Println("sample file result:", sampleDist)

	dist := calculateDistance(file1)
	fmt.Println("input file result:", dist)

}

func calculateDistance(file *os.File) int {
	scanner := bufio.NewScanner(file)
	first, second := []int{}, []int{}
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			log.Fatal(errors.New("invalid input"))
		}

		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)
		}

		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)
		}
		first = append(first, num1)
		second = append(second, num2)
	}

	slices.Sort(first)
	slices.Sort(second)

	distance := make([]int, len(first))

	for i, _ := range first {
		d := int(math.Abs(float64(second[i] - first[i])))
		distance[i] = d
	}

	res := 0
	for _, n := range distance {
		res += n
	}

	return res
}

func calculateSimilarity(file *os.File) int {
	scanner := bufio.NewScanner(file)
	first := []int{}
	second := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			log.Fatal(errors.New("invalid input"))
		}

		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)
		}

		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)
		}

		first = append(first, num1)
		if _, ok := second[num2]; ok {
			second[num2]++
		} else {
			second[num2] = 1
		}

	}
	similarity := 0
	for _, n := range first {
		similarity += n * second[n]
	}

	return similarity
}
