package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	// Read input
	var existingUsernames string = "longusername1,longusername2,longusername3"
	var newRegistrations string = "shortname,longusername1,anotherlongname,longusername2,verylongusernamehere,longusername3"
	fmt.Scanln(&existingUsernames)
	fmt.Scanln(&newRegistrations)

	// Parse input strings
	existing := strings.Split(existingUsernames, ",")
	attempts := strings.Split(newRegistrations, ",")

	// Create set using Go idiom
	userSet := make(map[string]struct{})

	// TODO: Write your code below
	// 1. Add existing usernames to the set
	for _, user := range existing {
		userSet[user] = struct{}{}
	}
	// 2. Process each registration attempt
	registrationAttempts := 0
	successfulRegistration := 0
	duplicated := 0
	for _, user := range attempts {
		if _, ok := userSet[user]; ok {
			fmt.Printf("Username taken: %s\n", user)
			duplicated++
		} else {
			userSet[user] = struct{}{}
			fmt.Printf("Registered: %s\n", user)
			successfulRegistration++
		}
		registrationAttempts++
	}
	// 3. Display registration results for each attempt
	// 4. Calculate and display the summary

	fmt.Println("Registration Summary:")
	fmt.Printf("Initial usernames: %d\n", len(existing))
	fmt.Printf("Registration attempts: %d\n", registrationAttempts)
	fmt.Printf("Successful registrations: %d\n", successfulRegistration)
	fmt.Printf("Rejected (duplicated): %d\n", duplicated)
	fmt.Printf("Total registered usernames: %d\n", len(userSet))

	// 5. List all registered usernames (remember to sort for consistent output)

	sortedUsers := []string{}

	for users := range userSet {
		sortedUsers = append(sortedUsers, users)
	}
	slices.Sort(sortedUsers)
	fmt.Println("All registered usernames:")
	for _, user := range sortedUsers {
		fmt.Printf("@%s\n", user)
	}

}
