package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	safeCounter, tolerateCounter := getSafeCounter(f)

	fmt.Printf("For file %s, total safe lists = %d\n", f.Name(), safeCounter)
	fmt.Printf("For file %s, total tolarate lists = %d\n", f.Name(), tolerateCounter)

	f2, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	safeCounter, tolerateCounter = getSafeCounter(f2)
	fmt.Printf("For file %s, total safe lists = %d\n", f2.Name(), safeCounter)
	fmt.Printf("For file %s, total tolerate lists = %d\n", f2.Name(), tolerateCounter)

}

func getSafeCounter(f *os.File) (int, int) {
	scanner := bufio.NewScanner(f)
	safeCounter := 0
	tolerateCounter := 0
	for scanner.Scan() {
		line := scanner.Text()
		readings := strings.Fields(line)
		if safe(readings) {
			//	fmt.Printf("List %s is safe\n", readings)
			safeCounter++
			tolerateCounter++
		} else {
			if canTolerateSingleBadLevel(readings) {
				tolerateCounter++
			}
		}
	}
	return safeCounter, tolerateCounter
}

func safe(list []string) bool {
	increasing, decreasing := false, false
	for i := 0; i < len(list)-1; i++ {
		first, err := strconv.ParseInt(list[i], 10, 64)
		if err != nil {
			panic(err)
		}
		second, err := strconv.ParseInt(list[i+1], 10, 64)
		if err != nil {
			panic(err)
		}

		diff := int(math.Abs(float64(first - second)))

		if !increasing && !decreasing {
			if second > first && diff <= 3 {
				increasing = true
			} else if second < first && diff <= 3 {
				decreasing = true
			} else {
				return false
			}
		}

		if increasing {
			if second <= first || diff > 3 {
				return false
			}
		}

		if decreasing {

			if second >= first || diff > 3 {
				return false
			}
		}

	}

	return true
}

func canTolerateSingleBadLevel(list []string) bool {
	n := len(list)
	for i := 0; i < n; i++ {
		subList := append([]string{}, list[:i]...)
		subList = append(subList, list[i+1:]...)
		if isSafe := safe(subList); isSafe {
			return true
		}
	}
	return false
}
