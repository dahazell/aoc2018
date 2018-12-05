package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func calculateFrequencyTotal(frequencies []string) (total int64, err error) {

	for _, s := range frequencies {
		if s == "" {
			continue
		}
		value, err := strconv.ParseInt(s, 10, 64)
		if err != nil {

			return 0, err
		}
		total += value
	}

	return total, nil
}

func firstFrequencyFoundTwice(frequencies []string) (total int64, err error) {
	foundFrequencies := make(map[int64]int)
	foundFrequencies[0] = 1
	foundTwice := false
	for !foundTwice {
		for _, s := range frequencies {
			if s == "" {
				continue
			}
			value, err := strconv.ParseInt(s, 10, 64)
			if err != nil {

				return 0, err
			}
			total += value
			_, ok := foundFrequencies[total]
			if ok {
				return total, nil
			}
			foundFrequencies[total] = 1
		}
	}
	return total, nil
}

func main() {

	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("Error reading: %v\n", err)
		os.Exit(1)
	}
	frequencies := strings.Split(string(input), "\n")

	total, err := calculateFrequencyTotal(frequencies)
	if err != nil {
		fmt.Printf("Error calculating frequency: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Total Frequency: %v\n", total)
	foundTwice, err := firstFrequencyFoundTwice(frequencies)
	if err != nil {
		fmt.Printf("Error finding first frequency hit twice: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("First Frequency found twice: %v\n", foundTwice)

}
