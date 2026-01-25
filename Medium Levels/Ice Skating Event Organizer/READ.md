Write a function organizeSkatingEvent that takes pastSessions, requestedSessions, and boothNames and returns a formatted event summary.

The function merges skating sessions, calculates statistics, and analyzes booth names for palindromes and anagrams.

Logic:

Merge past and requested session durations into one sorted array (ascending order)
Calculate total skating time and average duration (rounded to 1 decimal place)
Check each booth name: mark as "palindrome" if it reads the same forwards/backwards (case-insensitive), or "anagram" if it's an anagram of any other booth name, otherwise "normal"
Format the results into a multi-line string
Parameters:

pastSessions ([]int): Array of past session durations in minutes
requestedSessions ([]int): Array of requested session durations in minutes
boothNames ([]string): Array of booth names to analyze
Returns: Multi-line string with event summary. Format: Sessions: [1,2,3] Total: 6 minutes Average: 2.0 minutes Booths: name1(palindrome), name2(anagram), name3(normal)
