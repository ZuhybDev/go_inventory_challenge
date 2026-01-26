package main

import (
	"fmt"
	"slices"
	"strings"
)

func getUniqueItems(originalItems []string) []string {

	originalItemsMap := make(map[string]struct{})
	finalData := make([]string, 0)

	for _, value := range originalItems {
		if _, ok := originalItemsMap[value]; ok {
			continue
		} else {
			originalItemsMap[value] = struct{}{}
			finalData = append(finalData, value)
		}
	}

	return finalData

}

func main() {
	// Read input
	var input1 string = "cat,dog,cat,bird"
	var input2 string = "fish,cat,hamster"
	// fmt.Scanln(&input1)
	// fmt.Scanln(&input2)

	// Parse input strings into slices
	originalItems := strings.Split(input1, ",")
	additionalItems := strings.Split(input2, ",")

	allItems := make([]string, 0)

	allItems = append(originalItems, additionalItems...)

	fmt.Println(allItems)

	// TODO: Write your code here
	// 1. Create the getUniqueItems function that uses map[string]struct{} idiom
	// 2. Process the original items to get unique items

	finalData := getUniqueItems(allItems)

	// 3. Combine both slices and get unique items from combined list
	// 4. Calculate

	duplicates := len(allItems) - len(finalData)

	// Display system header
	fmt.Println("=== UNIQUE ITEM COLLECTOR ===")

	// Display original and additional items
	fmt.Printf("Original items: %s\n", strings.Join(originalItems, ","))
	fmt.Printf("Additional items: %s\n", strings.Join(additionalItems, ","))

	uniqueItems := getUniqueItems(originalItems)

	slices.Sort(uniqueItems)

	fmt.Printf("Unique original items: %s\n", strings.Join(uniqueItems, ","))

	// TODO: Display unique original items, final unique items, and statistics
	fmt.Printf("Final unique items: %s\n", strings.Join(finalData, ","))
	// Use fmt.Printf and strings.Join for output formatting

	fmt.Println("=== COLLECTION STATISTICS ===")
	fmt.Printf("Total original items: %d\n", len(originalItems))
	fmt.Printf("Total additional items: %d\n", len(additionalItems))
	fmt.Printf("Total combined items: %d\n", len(allItems))
	fmt.Printf("Unique items found: %d\n", len(finalData))
	fmt.Printf("Duplicates removed: %d\n", duplicates)
	fmt.Println("Unique item collection completed successfully")

}
