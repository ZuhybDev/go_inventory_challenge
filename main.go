package main

import (
	"fmt"
	"strings"
)

func main() {
	// Read input
	var teamMembersInput string
	var peopleToCheckInput string
	fmt.Scanln(&teamMembersInput)
	fmt.Scanln(&peopleToCheckInput)

	// Parse input strings
	teamMembers := strings.Split(teamMembersInput, ",")
	peopleToCheck := strings.Split(peopleToCheckInput, ",")

	// TODO: Write your code below
	// Create a set using map[string]struct{}
	// Add team members to the set
	// Check each person using the comma ok idiom
	// Display verification results and summary

}
