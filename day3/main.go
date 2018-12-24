package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type (
	coord struct {
		x int
		y int
	}
)

func findOverlaps(lines []string) (int, map[coord]int) {
	//#1331 @ 599,609: 14x19

	coordCount := make(map[coord]int)

	for _, s := range lines {
		if s == "" {
			continue
		}
		coords := strings.Split(s, " ")
		x, y := strings.Split(coords[2], ",")[0], strings.TrimSuffix(strings.Split(coords[2], ",")[1], ":")
		size := strings.Split(s, " ")[3]
		wStr, hStr := strings.Split(size, "x")[0], strings.Split(size, "x")[1]
		w, _ := strconv.ParseInt(wStr, 10, 32)
		h, _ := strconv.ParseInt(hStr, 10, 32)
		for i := 0; i < int(w); i++ {
			for j := 0; j < int(h); j++ {
				xcoord, _ := strconv.ParseInt(x, 10, 32)
				ycoord, _ := strconv.ParseInt(y, 10, 32)
				c := coord{
					x: i + int(xcoord),
					y: j + int(ycoord),
				}
				v, _ := coordCount[c]
				coordCount[c] = v + 1
			}

		}

	}
	overlaps := 0
	for _, v := range coordCount {
		if v > 1 {
			overlaps++
		}
	}
	return overlaps, coordCount
}

// Relies on input from the result of findOverlaps
func findIDWithoutOverlap(lines []string, coordCount map[coord]int) string {

	for _, s := range lines {
		if s == "" {
			continue
		}
		coords := strings.Split(s, " ")
		x, y := strings.Split(coords[2], ",")[0], strings.TrimSuffix(strings.Split(coords[2], ",")[1], ":")
		size := strings.Split(s, " ")[3]
		wStr, hStr := strings.Split(size, "x")[0], strings.Split(size, "x")[1]
		w, _ := strconv.ParseInt(wStr, 10, 32)
		h, _ := strconv.ParseInt(hStr, 10, 32)
		noOverlap := true
		for i := 0; i < int(w); i++ {
			for j := 0; j < int(h); j++ {
				xcoord, _ := strconv.ParseInt(x, 10, 32)
				ycoord, _ := strconv.ParseInt(y, 10, 32)
				c := coord{
					x: i + int(xcoord),
					y: j + int(ycoord),
				}
				v, _ := coordCount[c]
				if v > 1 {
					noOverlap = false
				}
			}
		}
		if noOverlap {
			return strings.Split(s, " ")[0]
		}
	}
	return ""
}

func main() {

	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("Error reading: %v\n", err)
		os.Exit(1)
	}
	lines := strings.Split(string(input), "\n")
	overlaps, coordMap := findOverlaps(lines)
	id := findIDWithoutOverlap(lines, coordMap)

	fmt.Printf("overlaps: %v\nnon overlaping id: %v\n", overlaps, id)
}
