package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const file = "inputs.txt"

func main() {
	// Open the file
	file, _ := os.Open(file)
	scanner := bufio.NewScanner(file)
	sumArray := make([]int, 0)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		// Check for a line break
		if line == "" {
			// Add finished sum to array
			sumArray = append(sumArray, sum)

			// Reset sum for the next set
			sum = 0
		}

		// Split line by spaces
		values := strings.Split(line, " ")

		// Sum up parsed integers
		for _, val := range values {
			num, err := strconv.Atoi(val)
			if err == nil {
				sum += num
			}
		}
	}

	sumArray = sortArray(sumArray)
	biggestSum := sumArray[0]
	threeBiggestSum := sumArray[0] + sumArray[1] + sumArray[2]

	fmt.Println("Biggest sum:", biggestSum)
	fmt.Println("Three biggest sum:", threeBiggestSum)
}

func sortArray(arr []int) []int {
	// Sort the array in descending order
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	return arr
}
