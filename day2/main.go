package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func sortString(s string) string {
	result := strings.Split(s, "")
	sort.Strings(result)
	return strings.Join(result, "")
}

//finds repeats but only to one occurence e.g. aaaa returns 1,0
func findRepeats(s string) (doubles, triples int) {
	s = sortString(s)
	skipCount := 0
	for i, c := range s {
		if skipCount > 0 {
			skipCount--
			continue
		}
		if i+1 < len(s) && c == rune(s[i+1]) {
			// found at least a double
			if i+2 < len(s) && c == rune(s[i+2]) {
				//could be 4

				if i+3 >= len(s) || c != rune(s[i+3]) {
					triples = 1
				}
				skipCount = 2
			} else {
				doubles = 1
			}
		}
		if triples == 1 && doubles == 1 {
			break
		}
	}
	return doubles, triples
}

func calculateCheckSum(input []string) int {
	doubles := 0
	triples := 0
	for _, s := range input {
		a, b := findRepeats(s)
		doubles += a
		triples += b
	}
	return doubles * triples
}

func matchedIDs(inputs []string) {
	count := 0
	for i := range inputs {
		//for j := i + 1; j < len(inputs); j++ {

		for j := range inputs {
			if differByOne(inputs[i], inputs[j]) {
				fmt.Printf("first part: %v second part: %v\n", inputs[i], inputs[j])
				break
			} else {
				count++
			}
		}
	}
}

func differByOne(a, b string) bool {
	diffs := 0
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diffs++
		}
		if diffs > 1 {
			return false
		}
	}
	return diffs == 1
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("Error reading: %v\n", err)
		os.Exit(1)
	}
	inputs := strings.Split(string(input), "\n")
	fmt.Printf("checksum: %v\n", calculateCheckSum(inputs))
	matchedIDs(inputs)

}
