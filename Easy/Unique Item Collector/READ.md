Create a duplicate removal system that uses Go's idiomatic set pattern to filter unique items from collections of strings. This challenge will test your ability to use the map[string]struct{} idiom to track seen items and build collections containing only unique elements.

You will receive two inputs:

A string containing items separated by commas (e.g., "apple,banana,apple,orange,banana,grape,apple")
A string containing additional items to merge, also separated by commas (e.g., "kiwi,apple,mango,banana,kiwi")
Your task is to:

Create a function called getUniqueItems that takes a slice of strings and returns a new slice containing only the unique strings
Inside this function, use the map[string]struct{} idiom to track which items have been seen
Iterate through the input slice and for each item:
Check if the item exists in your set using the comma ok idiom
If the item hasn't been seen before, add it to both the set and the result slice
Parse the first input by splitting on commas to get the initial item list
Parse the second input by splitting on commas to get the additional items
Display the system header: "=== UNIQUE ITEM COLLECTOR ==="
Display the original items: "Original items: [comma-separated list of all original items]"
Display the additional items: "Additional items: [comma-separated list of all additional items]"
Use your getUniqueItems function to get unique items from the original list
Display the unique original items: "Unique original items: [comma-separated list of unique items from original list]"
Combine both input slices into a single slice containing all items
Use your getUniqueItems function to get unique items from the combined list
Display the final unique items: "Final unique items: [comma-separated list of all unique items]"
Display collection statistics:
"=== COLLECTION STATISTICS ==="
"Total original items: [count of original items]"
"Total additional items: [count of additional items]"
"Total combined items: [count of all items combined]"
"Unique items found: [count of unique items]"
"Duplicates removed: [total combined items minus unique items]"
Display the completion message: "Unique item collection completed successfully"
Use the strings package to split input strings and the fmt package for output. When joining items for display, use strings.Join with a comma separator. This challenge demonstrates how Go's set idiom efficiently solves the common problem of removing duplicates from collections, a pattern you'll use frequently in data processing applications.
