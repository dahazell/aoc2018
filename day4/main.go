package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type (
	guard struct {
		id         int
		timeAsleep int
		minute     int
		minuteFreq int
	}
)

func parseInput(input []string) map[int][]int {
	// guard ID int to array ints representing minutes in an hour and times asleep in that hour
	guardSleepSchedule := make(map[int][]int)
	guardID := 0
	asleep := 0
	awake := 0
	for _, s := range input {
		if strings.Contains(s, "Guard") {
			guardID, _ = strconv.Atoi(strings.Split(s, " ")[3][1:])
			continue
		}

		if strings.Contains(s, "asleep") {
			endofTimeStamp := strings.Split(s, " ")[1]
			hourAndMinute := endofTimeStamp[:len(endofTimeStamp)-1]
			asleep, _ = strconv.Atoi(strings.Split(hourAndMinute, ":")[1])
			continue
		}

		if strings.Contains(s, "wakes") {
			endofTimeStamp := strings.Split(s, " ")[1]
			hourAndMinute := endofTimeStamp[:len(endofTimeStamp)-1]
			awake, _ = strconv.Atoi(strings.Split(hourAndMinute, ":")[1])
			schedule, exists := guardSleepSchedule[guardID]
			if !exists {
				//make our array of times
				schedule = make([]int, 60)
			}
			for i := asleep; i < awake; i++ {
				schedule[i]++
			}
			guardSleepSchedule[guardID] = schedule
		}
	}
	return guardSleepSchedule
}

func findMostFrequentMinAsleep(input []string) guard {
	// yes this is strange something outside of here should parse the input
	// and this should just act on the map returned but this is a quick refactor for part 2
	// might come back to this one day
	guardSleepSchedule := parseInput(input)

	highestFreq := guard{}
	for k, v := range guardSleepSchedule {
		for i, count := range v {
			if count > highestFreq.minuteFreq {
				highestFreq.minuteFreq = count
				highestFreq.minute = i
				highestFreq.id = k
			}
		}
	}
	return highestFreq
}
func findSleepyGuard(input []string) guard {
	// yes this is strange something outside of here should parse the input
	// and this should just act on the map returned but this is a quick refactor for part 2
	// might come back to this one day
	guardSleepSchedule := parseInput(input)

	// guard spent the most time asleep
	mostAsleepGuard := guard{}
	for k, v := range guardSleepSchedule {
		timeAsleep := 0
		for _, i := range v {
			timeAsleep += i
		}
		if timeAsleep > mostAsleepGuard.timeAsleep {
			mostAsleepGuard = guard{
				id:         k,
				timeAsleep: timeAsleep,
			}
		}
	}
	// find favourite minute asleep
	favMin := 0
	freqMin := 0
	schedule := guardSleepSchedule[mostAsleepGuard.id]

	for i, count := range schedule {
		if count > schedule[favMin] {
			favMin = i
			freqMin = count
		}
	}
	mostAsleepGuard.minuteFreq = freqMin
	mostAsleepGuard.minute = favMin
	return mostAsleepGuard
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("Error reading: %v\n", err)
		os.Exit(1)
	}

	lines := strings.Split(string(input), "\n")
	sort.Strings(lines)
	sleepy := findSleepyGuard(lines)
	fmt.Printf("Sleepy guard: %v\n", sleepy)
	fmt.Printf("answer part 1: %v\n", sleepy.id*sleepy.minute)
	mostFreq := findMostFrequentMinAsleep(lines)
	fmt.Printf("most asleep on a minute guard: %v\n", mostFreq)

	fmt.Printf("answer part 2: %v\n", mostFreq.id*mostFreq.minute)
}
