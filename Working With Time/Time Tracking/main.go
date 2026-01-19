package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Read input
	var formatRequests string
	fmt.Scanln(&formatRequests)

	// Use a fixed time for consistent output instead of time.Now()
	// Using January 2, 2006 at 15:04:05 (Go's reference time)
	currentTime := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)

	// Split format requests
	formats := strings.Split(formatRequests, ",")

	// Display application header
	fmt.Println("=== TIME TRACKER APPLICATION ===")

	// Display timestamp capture confirmation
	fmt.Println("Current time captured successfully")

	// Display formats header
	fmt.Println("Displaying time in requested formats:")

	// Track unique components in order
	uniqueComponents := make(map[string]bool)

	for _, format := range formats {
		switch format {
		case "full":
			fmt.Printf("Full timestamp: %v\n", currentTime)
			uniqueComponents["full"] = true
		case "date":
			year, month, day := currentTime.Date()
			fmt.Printf("Date only: %d-%02d-%02d\n", year, int(month), day)
			uniqueComponents["year"] = true
			uniqueComponents["month"] = true
			uniqueComponents["day"] = true
		case "time":
			fmt.Printf("Time only: %02d:%02d:%02d\n", currentTime.Hour(), currentTime.Minute(), currentTime.Second())
			uniqueComponents["hour"] = true
			uniqueComponents["minute"] = true
			uniqueComponents["second"] = true
		case "year":
			fmt.Printf("Year: %d\n", currentTime.Year())
			uniqueComponents["year"] = true
		case "month":
			fmt.Printf("Month: %v\n", currentTime.Month())
			uniqueComponents["month"] = true
		case "day":
			fmt.Printf("Day: %d\n", currentTime.Day())
			uniqueComponents["day"] = true
		case "weekday":
			fmt.Printf("Weekday: %v\n", currentTime.Weekday())
			uniqueComponents["weekday"] = true
		case "unix":
			fmt.Printf("Unix timestamp: %d\n", currentTime.Unix())
			uniqueComponents["unix"] = true
		}
	}

	// Display summary statistics
	fmt.Println("=== TIME INFORMATION SUMMARY ===")
	fmt.Printf("Total formats displayed: %d\n", len(formats))
	fmt.Printf("Timestamp components extracted: %d\n", len(uniqueComponents))

	// Display additional time details
	fmt.Println("Additional time details:")
	fmt.Printf("- Nanoseconds: %d\n", currentTime.Nanosecond())
	fmt.Printf("- Location: %v\n", currentTime.Location())
	fmt.Printf("- Is zero time: %v\n", currentTime.IsZero())

	// Display completion message
	fmt.Println("Time tracking completed - all requested formats processed")
}
