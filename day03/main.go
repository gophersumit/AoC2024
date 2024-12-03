package main

import (
	"AoC2024/lib"
	"fmt"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)

func calculateTotal(data string, part1 bool) uint64 {
	var res uint64
	do := true

	matches := re.FindAllStringSubmatch(data, -1)
	for _, match := range matches {
		action := match[1]
		j := match[2]
		k := match[3]

		if action == "don't()" {
			do = false
		} else if action == "do()" {
			do = true
		} else if do || part1 {
			num1, _ := strconv.Atoi(j)
			num2, _ := strconv.Atoi(k)
			res += uint64(num1 * num2)

		}
	}

	return res
}

func main() {
	part1()
	part2()
}

func part1() {
	lines := lib.GetLines("input.txt")
	var res uint64
	for _, line := range lines {
		res += calculateTotal(line, true)
	}

	fmt.Println(res)

}

func part2() {
	lines := lib.GetLines("input.txt")
	var res uint64
	for _, line := range lines {
		res += calculateTotal(line, false)
	}

	fmt.Println(res)
}
