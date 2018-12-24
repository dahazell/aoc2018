package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

const abc = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

//dabAcCaCBAcCcaDA
func reactPolymers(input string) string {
	//iterate until reaction or end of string
	for i, c := range input {
		if i == len(input)-1 {
			//don't go off the end of string we only need to compare penultimate to last char
			continue
		}

		nextChar := rune(input[i+1])
		//skip if identical
		if c == nextChar {
			continue
		}
		// now make both lower case and if they match react
		c = unicode.ToLower(c)
		nextChar = unicode.ToLower(nextChar)

		if c == nextChar {
			//fmt.Printf("XXX matched %c %c\n", c, nextChar)
			reactedString := ""
			//react
			if i == 0 {
				//edge case for first characters reacting
				reactedString = input[2:]
			}
			reactedString = input[:i] + input[i+2:]

			return reactPolymers(reactedString)
		}

	}
	return input
}

//part 2
func shortestPolymerRemovedUnits(input string) (int, string) {
	//This is very slow but didn't want to rework reactPolymers just for part 2 a future todo perhaps
	shortest := 0
	unitRemoved := ""
	for _, c := range abc {
		reduced := strings.Replace(input, string(c), "", -1)
		reduced = strings.Replace(reduced, string(unicode.ToLower(c)), "", -1)
		s := reactPolymers(reduced)
		if shortest == 0 || len(s) < shortest {
			shortest = len(s)
			unitRemoved = string(c) + string(unicode.ToLower(c))
		}
	}
	return shortest, unitRemoved
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("Error reading: %v\n", err)
		os.Exit(1)
	}
	sanitised := strings.TrimSpace(string(input))
	reacted := reactPolymers(sanitised)
	fmt.Printf("Reacted string: %s\n", reacted)
	fmt.Printf("Len of reacted string: %v\n", len(reacted))
	length, removed := shortestPolymerRemovedUnits(sanitised)
	fmt.Printf("length: %v  removed: %v\n", length, removed)
}
