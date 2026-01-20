package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Create scanner for reading input lines
	scanner := bufio.NewScanner(os.Stdin)

	// Read log entries
	scanner.Scan()
	logEntries := scanner.Text()

	// Read layout patterns
	scanner.Scan()
	layoutPatterns := scanner.Text()

	//tests
	// logEntries := "2023-09-01:Back to school,Sep 15 2023:Mid month,2023/09/30:Month end,Oct 1 2023:New month"
	// layoutPatterns := "2006-01-02,Jan 2 2006,2006/01/02,Jan 2 2006"

	// Parse inputs
	entries := strings.Split(logEntries, ",")
	layouts := strings.Split(layoutPatterns, ",")

	// // Initialize counters for statistics
	totalEntries := len(entries)
	successfulParses := 0
	failedParses := 0
	layoutUsage := make(map[string]int)
	var parsedTimes []time.Time

	// TODO: Write your code below
	// 1. Display application header
	fmt.Println("=== LOG TIMESTAMP PARSER ===")
	// 2. Display processing status
	fmt.Println("Processing log entries with multiple timestamp formats...")
	// 3. Process each log entry:
	//    - Split entry on ":" to get timestamp and message
	//    - Try each layout pattern until one works
	//    - Display parsing results
	//    - Update statistics

	for i, log := range entries {
		trimmed := strings.TrimSpace(log)

		parts := strings.Split(trimmed, ":")
		date := parts[0]
		message := parts[1]

		parsedTime, err := time.Parse(layouts[i], date)

		if err != nil {
			fmt.Printf("Failed to parse: %s - no matching layout found\n", date)
			failedParses++
		} else {
			fmt.Printf("Parsed: %s -> %v (using layout: %s)\n", date, parsedTime, layouts[i])
			fmt.Printf("  Message: %s\n", message)
			successfulParses++
		}
		layoutUsage[layouts[i]]++
		parsedTimes = append(parsedTimes, parsedTime)

	}

	// 4. Display parsing statistics
	fmt.Println("=== PARSING RESULTS ===")
	fmt.Printf("Total log entries processed: %d\n", totalEntries)
	fmt.Printf("Successfully parsed timestamps: %d\n", successfulParses)
	fmt.Printf("Failed parsing attempts: %d\n", failedParses)
	// 5. Display layout usage statistics
	fmt.Println("Layout pattern usage:")
	for _, layout := range layouts {
		fmt.Printf("  %s: %d times\n", layout, layoutUsage[layout])
	}
	// 6. If timestamps were parsed, display chronological analysis
	fmt.Println("=== CHRONOLOGICAL ANALYSIS ===")

	if len(parsedTimes) > 0 {
		earliest := parsedTimes[0]
		latest := parsedTimes[0]

		for _, t := range parsedTimes {
			if t.Before(earliest) {
				earliest = t
			}

			if t.After(latest) {
				latest = t
			}
		}
		timeSpan := latest.Sub(earliest)
		fmt.Printf("Earliest timestamp: %v\n", earliest)
		fmt.Printf("Latest timestamp: %v\n", latest)
		fmt.Printf("Time span covered: %v\n", timeSpan)
	} else {
		fmt.Println("No valid timestamps parsed.")
	}
	// 7. Display completion message
	fmt.Println("Log parsing completed - timestamps converted to time objects")

}
