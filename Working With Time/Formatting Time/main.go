package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Read input
	var dateInput string = "2027:09:12:21:10"
	var formatInput string = "iso,american,european,long,time12,time24,full,compact"
	// fmt.Scanln(&dateInput)
	// fmt.Scanln(&formatInput)

	// Parse date components by splitting on colons
	dateComponents := strings.Split(dateInput, ":")

	// Convert string components to integers
	year, _ := strconv.Atoi(dateComponents[0])
	month, _ := strconv.Atoi(dateComponents[1])
	day, _ := strconv.Atoi(dateComponents[2])
	hour, _ := strconv.Atoi(dateComponents[3])
	minute, _ := strconv.Atoi(dateComponents[4])

	// Create time.Time object
	sourceTime := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)

	// Parse format requests
	formatRequests := strings.Split(formatInput, ",")

	// TODO: Write your code below
	// Display the application header
	fmt.Println("=== DATE/TIME FORMATTER ===")
	// Display the source time
	fmt.Printf("Source time: %v\n", sourceTime)
	// Process each format request and display formatted
	fmt.Println("Formatted outputs:")
	for _, fr := range formatRequests {

		switch fr {
		case "iso":
			formatted := sourceTime.Format("2006-01-02")
			fmt.Printf("ISO format: %v\n", formatted)
		case "american":
			formatted := sourceTime.Format("01/02/2006")
			fmt.Printf("American format: %v\n", formatted)
		case "european":
			formatted := sourceTime.Format("02/01/2006")
			fmt.Printf("European format: %v\n", formatted)
		case "long":
			formatted := sourceTime.Format("January 2, 2006")
			fmt.Printf("Long format: %v\n", formatted)
		case "time12":
			formatted := sourceTime.Format("3:04 PM")
			fmt.Printf("12-hour time: %v\n", formatted)
		case "time24":
			formatted := sourceTime.Format("15:04")
			fmt.Printf("24-hour time: %v\n", formatted)
		case "full":
			formatted := sourceTime.Format("Monday, January 2, 2006 at 3:04 PM")
			fmt.Printf("Full format: %v\n", formatted)
		case "compact":
			formatted := sourceTime.Format("20060102")
			fmt.Printf("Compact format: %v\n", formatted)
		}

	}
	// Calculate and display statistics

	fmt.Println("=== FORMATTING SUMMARY ===")
	fmt.Printf("Total formats applied: %d\n", len(formatRequests))
	fmt.Printf("Date components: %d-%d-%d\n", year, month, day)
	fmt.Printf("Time components: %d:%d\n", hour, minute)
	// Display reference information
	fmt.Println("=== REFERENCE DATE INFO ===")
	fmt.Println("Go uses reference date: Mon Jan 2 15:04:05 MST 2006")
	fmt.Println("This represents: January 2nd, 2006 at 3:04:05 PM")
	fmt.Println("Layout patterns arrange these components as needed")
	// Display completion message
	fmt.Println("Date formatting completed using Go's layout system")

}
