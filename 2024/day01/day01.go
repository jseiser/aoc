package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var left []int
	var right []int

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		left_id, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}

		right_id, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}

		left = append(left, left_id)
		right = append(right, right_id)
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	slices.Sort(left)
	slices.Sort(right)
	var counter int

	for i := 0; i < len(left); i++ {
		if left[i] >= right[i] {
			d := left[i] - right[i]
			counter = counter + d
		} else {
			d := right[i] - left[i]
			counter = counter + d
		}
	}

	fmt.Printf("Total Distance: %d\n", counter)

	fm := getFM(right)

	var score int64 = 0
	for i := 0; i < len(left); i++ {
		score += int64(left[i]) * int64(fm[left[i]])
	}

	fmt.Printf("Sim Score: %d\n", score)

}

func getFM(right []int) map[int]int {
	m := make(map[int]int)
	for _, v := range right {
		m[v]++
	}
	return m
}
