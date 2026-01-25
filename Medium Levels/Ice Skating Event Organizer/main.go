package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func organizeSkatingEvent(pastSessions []int, requestedSessions []int, boothNames []string) string {

	combinedData := append(pastSessions, requestedSessions...)

	sort.Ints(combinedData)

	temStrings := make([]string, len(combinedData))

	for i, v := range combinedData {
		temStrings[i] = strconv.Itoa(v)
	}

	numbers := strings.Join(temStrings, ",")

	fmt.Printf("Sessions: %s\n", numbers)

	var total int
	for _, n := range combinedData {
		total += n
	}

	result := float64(total) / float64(len(combinedData))

	fmt.Printf("Total: %d minutes\n", total)
	fmt.Printf("Average: %.1f minutes\n", result)

	isPalindrome := func(s1 string) bool {
		n := len(s1)

		for i := 0; i < n/2; i++ {
			if s1[i] != s1[n-1-i] {
				return false
			}

		}

		return true
	}

	isAnagram := func(s1, s2 string) bool {

		if len(s1) != len(s2) {
			return false
		}

		//helper func
		counts := make(map[rune]int)

		for _, char := range s1 {
			counts[char]++
		}

		for _, char := range s2 {
			counts[char]--
			if counts[char] < 0 {
				return false
			}
		}

		return true

	}

	var boothResults []string
	for i, booth := range boothNames {
		status := "normal"

		if isPalindrome(booth) {
			status = "palindrome"
		} else {
			// Check current booth against ALL other booths
			for j, other := range boothNames {
				if i == j {
					continue
				}
				if isAnagram(booth, other) {
					status = "anagram"
					break
				}
			}
		}
		boothResults = append(boothResults, fmt.Sprintf("%s(%s)", booth, status))
	}

	return fmt.Sprintf("Booths: %s\n", strings.Join(boothResults, ", "))

}

func main() {
	pastSessions := []int{
		10, 5,
	}

	requestedSessions := []int{
		7,
	}

	boothNames := []string{
		"kayak", "hello", "world",
	}

	organizeSkatingEvent(pastSessions, requestedSessions, boothNames)
}
