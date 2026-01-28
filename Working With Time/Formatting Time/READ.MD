Create a multi-format date and time display system that formats time.Time objects using Go's unique reference date layout system. This challenge will test your ability to use the Format method with different layout patterns to create various date and time representations.

You will receive two inputs:

A string containing date information in the format "year:month:day:hour:minute" (e.g., "2023:12:25:14:30")
A string containing format requests separated by commas in the format "format1,format2,format3" (e.g., "iso,american,european,time12,time24")
Your task is to:

Parse the first input by splitting on colons to extract year, month, day, hour, and minute values
Convert these string values to integers using the strconv package
Create a time.Time object using time.Date with the parsed values (use time.UTC for location, 0 for seconds and nanoseconds)
Display the application header: "=== DATE/TIME FORMATTER ==="
Display the source time: "Source time: [complete_time_object]"
Parse the second input by splitting on commas to get format requests
Display the formatting header: "Formatted outputs:"
For each format request, use the appropriate layout string with the Format method:
For "iso": Use layout "2006-01-02" and display "ISO format: [formatted_result]"
For "american": Use layout "01/02/2006" and display "American format: [formatted_result]"
For "european": Use layout "02/01/2006" and display "European format: [formatted_result]"
For "long": Use layout "January 2, 2006" and display "Long format: [formatted_result]"
For "time12": Use layout "3:04 PM" and display "12-hour time: [formatted_result]"
For "time24": Use layout "15:04" and display "24-hour time: [formatted_result]"
For "full": Use layout "Monday, January 2, 2006 at 3:04 PM" and display "Full format: [formatted_result]"
For "compact": Use layout "20060102" and display "Compact format: [formatted_result]"
Calculate and display statistics:
"=== FORMATTING SUMMARY ==="
"Total formats applied: [number_of_format_requests]"
"Date components: [year]-[month]-[day]"
"Time components: [hour]:[minute]"
Display reference information:
"=== REFERENCE DATE INFO ==="
"Go uses reference date: Mon Jan 2 15:04:05 MST 2006"
"This represents: January 2nd, 2006 at 3:04:05 PM"
"Layout patterns arrange these components as needed"
Display the completion message: "Date formatting completed using Go's layout system"
Use the time package for time operations, the strings package to split input strings, the strconv package to convert strings to integers, and the fmt package for formatted output. Remember that Go's time formatting uses the reference date Mon Jan 2 15:04:05 MST 2006 as a template, and you arrange these components in the pattern you want. This challenge demonstrates how Go's unique layout-based formatting system works by having you apply multiple different format patterns to the same time object.
