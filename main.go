package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Read input
	var exerciseInput string = "Warm-up:10:Minute,Stre ngth:30:Minute,Cardio:20:Minute,Cool-down:10:Minute,Stretching:15:Minute"
	var restInput string = "3:Minute"
	// fmt.Scanln(&exerciseInput)
	// fmt.Scanln(&restInput)

	// Parse exercises from input
	exerciseEntries := strings.Split(exerciseInput, ",")

	// Parse rest period from input
	restParts := strings.Split(restInput, ":")

	var exerciseDuration time.Duration
	exerciseMap := make(map[string]int)
	var completeSession time.Duration
	var totalTime time.Duration
	var totalRest time.Duration

	fmt.Println("=== WORKOUT SESSION TIMER ===")
	fmt.Println("Setting up workout session...")
	// TODO: Write your code below
	// 1. Parse each exercise entry to extract name, duration value, and unit
	// 2. Convert duration values to integers and create time.Duration objects

	for _, entry := range exerciseEntries {
		parts := strings.Split(strings.TrimSpace(entry), ":")
		name := parts[0]
		valInt, _ := strconv.Atoi(parts[1])
		unit := parts[2]

		switch unit {
		case "Second":
			exerciseDuration = time.Duration(valInt) * time.Second
		case "Minute":
			exerciseDuration = time.Duration(valInt) * time.Minute
		case "Hour":
			exerciseDuration = time.Duration(valInt) * time.Hour
		}

		fmt.Printf("Exercise: %s - Duration: %v\n", name, exerciseDuration)
		exerciseMap[name] = valInt
		completeSession += exerciseDuration
		totalTime += exerciseDuration
	}
	// 3. Parse rest duration and create time.Duration object
	restVal, _ := strconv.Atoi(restParts[0])
	restUnit := restParts[1]
	var restDuration time.Duration

	switch restUnit {
	case "Second":
		restDuration = time.Duration(restVal) * time.Second
	case "Minute":
		restDuration = time.Duration(restVal) * time.Minute
	case "Hour":
		restDuration = time.Duration(restVal) * time.Hour
	}

	numRests := len(exerciseEntries) - 1
	if numRests > 0 {
		totalRestTime := restDuration * time.Duration(numRests)
		completeSession += totalRestTime
		totalRest += totalRestTime
	}

	fmt.Printf("Rest period between exercises: %v\n", restDuration)

	// 4. Display workout session information and calculations
	// 5. Calculate session statistics and time conversions

	fmt.Println("=== SESSION BREAKDOWN ===")

	var (
		longestName     string
		shortestName    string
		averageDuration time.Duration
	)

	for key, val := range exerciseMap {
		valInSec := val * 60

		if len(key) == 0 {
			shortestName = key
			longestName = key
		} else {
			if len(key) < len(shortestName) {
				shortestName = key
			}
			if len(key) > len(longestName) {
				longestName = key
			}
		}

		total := time.Duration(val) * time.Minute
		averageDuration += total

		fmt.Printf("%s: %v seconds\n", key, valInSec)
	}

	fmt.Printf("Rest periods: %v seconds each\n", restDuration.Seconds())

	// Remember to output all the required information as specified in the challenge
	fmt.Println("=== TIMING CALCULATIONS ===")
	fmt.Printf("Total exercise time: %v\n", totalTime)
	fmt.Printf("Total rest time: %v\n", totalRest)
	fmt.Printf("Complete session duration: %v\n", completeSession)

	fmt.Println("=== SESSION STATISTICS ===")
	fmt.Printf("Number of exercises: %d\n", len(exerciseMap))
	fmt.Printf("Longest exercise: %s\n", longestName)
	fmt.Printf("Shortest exercise: %s\n", shortestName)
	fmt.Printf("Average exercise duration: %d seconds\n", averageDuration)

}
