package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Read input
	var exerciseInput string
	var restInput string
	fmt.Scanln(&exerciseInput)
	fmt.Scanln(&restInput)

	// Parse exercises from input
	exerciseEntries := strings.Split(exerciseInput, ",")

	// Parse rest period from input
	restParts := strings.Split(restInput, ":")

	// Parse exercises
	var exercises []struct {
		name     string
		duration time.Duration
	}

	for _, entry := range exerciseEntries {
		parts := strings.Split(entry, ":")
		name := parts[0]
		durationValue, _ := strconv.Atoi(parts[1])
		unit := parts[2]

		var duration time.Duration
		switch unit {
		case "Second":
			duration = time.Duration(durationValue) * time.Second
		case "Minute":
			duration = time.Duration(durationValue) * time.Minute
		case "Hour":
			duration = time.Duration(durationValue) * time.Hour
		}

		exercises = append(exercises, struct {
			name     string
			duration time.Duration
		}{name, duration})
	}

	// Parse rest duration
	restValue, _ := strconv.Atoi(restParts[0])
	restUnit := restParts[1]
	var restDuration time.Duration
	switch restUnit {
	case "Second":
		restDuration = time.Duration(restValue) * time.Second
	case "Minute":
		restDuration = time.Duration(restValue) * time.Minute
	case "Hour":
		restDuration = time.Duration(restValue) * time.Hour
	}

	// Display header
	fmt.Println("=== WORKOUT SESSION TIMER ===")
	fmt.Println("Setting up workout session...")

	// Display exercises
	for _, exercise := range exercises {
		fmt.Printf("Exercise: %s - Duration: %v\n", exercise.name, exercise.duration)
	}
	fmt.Printf("Rest period between exercises: %v\n", restDuration)

	// Session breakdown
	fmt.Println("=== SESSION BREAKDOWN ===")
	for _, exercise := range exercises {
		fmt.Printf("%s: %.0f seconds\n", exercise.name, exercise.duration.Seconds())
	}
	fmt.Printf("Rest periods: %.0f seconds each\n", restDuration.Seconds())

	// Calculate totals
	var totalExerciseTime time.Duration
	for _, exercise := range exercises {
		totalExerciseTime += exercise.duration
	}

	totalRestTime := time.Duration(len(exercises)-1) * restDuration
	totalSessionTime := totalExerciseTime + totalRestTime

	// Timing calculations
	fmt.Println("=== TIMING CALCULATIONS ===")
	fmt.Printf("Total exercise time: %v\n", totalExerciseTime)
	fmt.Printf("Total rest time: %v\n", totalRestTime)
	fmt.Printf("Complete session duration: %v\n", totalSessionTime)

	// Session statistics
	fmt.Println("=== SESSION STATISTICS ===")
	fmt.Printf("Number of exercises: %d\n", len(exercises))

	// Find longest and shortest exercises
	longest := exercises[0]
	shortest := exercises[0]
	for _, exercise := range exercises {
		if exercise.duration > longest.duration {
			longest = exercise
		}
		if exercise.duration < shortest.duration {
			shortest = exercise
		}
	}

	fmt.Printf("Longest exercise: %s (%v)\n", longest.name, longest.duration)
	fmt.Printf("Shortest exercise: %s (%v)\n", shortest.name, shortest.duration)

	averageDuration := totalExerciseTime.Seconds() / float64(len(exercises))
	fmt.Printf("Average exercise duration: %.0f seconds\n", averageDuration)

	// Time conversions
	fmt.Println("=== TIME CONVERSIONS ===")
	fmt.Printf("Session duration in minutes: %.2f\n", totalSessionTime.Minutes())
	fmt.Printf("Session duration in hours: %.2f\n", totalSessionTime.Hours())

	fmt.Println("Workout session timer configured successfully")
}
