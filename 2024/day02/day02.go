package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var bad int
	var total int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var numbers []int
		for _, strNum := range strings.Fields(line) {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				panic(err) // Handle error appropriately
			}
			numbers = append(numbers, num)

		}
		total++
		fmt.Println(numbers)
		start_pattern := getPattern(numbers[0], numbers[1])
		for i := 0; i < len(numbers); i++ {
			if i == 0 {
				continue
			}
			same := getSame(numbers[i-1], numbers[i])
			if same {
				fmt.Printf("SAME NUMBER - BAD: %d %d\n", numbers[i-1], numbers[i])

				bad++
				break
			}

			cur_pattern := getPattern(numbers[i-1], numbers[i])
			if cur_pattern != start_pattern {
				fmt.Printf("BROKE PATTERN - BAD: %d %d\n", numbers[i-1], numbers[i])
				bad++
				break
			}

			change := absDiffInt(numbers[i-1], numbers[i])
			if change < 1 || change > 3 {
				fmt.Printf("INVALID SEQ - BAD: %d %d\n", numbers[i-1], numbers[i])
				bad++
				break
			}
		}
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total Safe: %d\n", total-bad)
}

func getPattern(a int, b int) bool {
	if a > b {
		return true
	} else {
		return false
	}
}

func getSame(a int, b int) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

func absDiffInt(a int, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
