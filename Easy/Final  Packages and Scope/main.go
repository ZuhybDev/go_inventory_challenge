package main

import (
	"fmt"
	"strings"
)

// TODO: Create your stringutil package functions here
// Since we can't use separate files, define the functions directly
// These would normally be in a separate package file

func Reverse(s string) string {
	// Convert to runes, reverse, convert back to string

	r := []rune(s)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)

}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)

}

func Length(s string) int {
	return len(s)
}

type Tests struct {
	FuncList string
	TestList string
}

func main() {
	//use this if you have ready terminal for reading inputs
	// scanner := bufio.NewScanner(os.Stdin)

	// scanner.Scan()
	// functionSpecs := scanner.Text()

	// scanner.Scan()
	// testStrings := scanner.Text()

	functionSpecs := "Reverse:reverses a string,ToUpper:converts to uppercase,ToLower:converts to lowercase,Length:returns string length"
	testStrings := "single|two words|three word phrase|four words in string|five words in this phrase"

	// Parse function specifications
	funcList := strings.Split(functionSpecs, ",")

	// Parse test strings
	testList := strings.Split(testStrings, "|")

	//slice

	fmt.Println("=== STRINGUTIL PACKAGE DEMONSTRATION ===")
	fmt.Println("Available functions in stringutil package:")

	for _, entry := range funcList {
		parts := strings.Split(entry, ":")
		name := parts[0]
		funcList := parts[1]
		fmt.Printf("- %s: %s\n", name, funcList)

	}

	// fmt.Println(testSlice)
	callsMade := 0
	fmt.Println("Testing stringutil functions:")
	for _, test := range testList {
		fmt.Printf("Testing with: \"%s\"\n", test)
		fmt.Printf("  stringutil.Reverse: \"%s\"\n", Reverse(test))
		fmt.Printf("  stringutil.ToUpper: \"%s\"\n", ToUpper(test))
		fmt.Printf("  stringutil.ToLower: \"%s\"\n", ToLower(test))
		fmt.Printf("  stringutil.Length: %d\n", Length(test))
		callsMade++

	}

	totalCalls := 4 * len(testList)

	//"=== STATISTICS ==="

	// shorts and longest

	var (
		longestString = testList[0]

		shortestString = testList[0]
	)

	for _, test := range testList {
		if len(test) > len(longestString) {
			longestString = test
		}

		if len(test) < len(shortestString) {
			shortestString = test
		}
	}

	fmt.Println("=== STATISTICS ===")
	fmt.Printf("Total test strings processed: %d\n", len(testList))
	fmt.Printf("Functions available in package: %d\n", len(funcList))
	fmt.Printf("Total function calls made: %d\n", totalCalls)
	fmt.Printf("Longest string: \"%s\" (length: %d)\n", longestString, len(longestString))
	fmt.Printf("Shortest string: \"%s\" (length: %d)\n", shortestString, len(shortestString))
	fmt.Println("Package demonstration completed successfully")
	//

	//

	// Remember to use your stringutil package functions like:
	// Reverse(testString)
	// ToUpper(testString)
	// ToUpper(testString)
	// ToLower(testString)
	// Length(testString)

}
