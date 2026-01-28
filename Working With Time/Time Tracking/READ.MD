Create a time tracking application that captures the current time and displays it in multiple formats. This challenge will test your ability to use the time.Now() function and work with time.Time objects to extract different components of the current moment.

You will receive one input:

A string containing display format requests separated by commas in the format "format1,format2,format3" (e.g., "full,date,time,year,month,day")
Your task is to:

Use time.Now() to capture the current time and store it in a variable
Display the application header: "=== TIME TRACKER APPLICATION ==="
Display the timestamp capture confirmation: "Current time captured successfully"
Parse the input by splitting on commas to get individual format requests
Display the formats header: "Displaying time in requested formats:"
For each format request, display the time information according to the specified format:
For "full": Display "Full timestamp: [complete_time_object]"
For "date": Display "Date only: [year]-[month]-[day]" (use the Date() method components)
For "time": Display "Time only: [hour]:[minute]:[second]" (use Hour(), Minute(), Second() methods)
For "year": Display "Year: [year_value]" (use Year() method)
For "month": Display "Month: [month_name]" (use Month() method)
For "day": Display "Day: [day_value]" (use Day() method)
For "weekday": Display "Weekday: [weekday_name]" (use Weekday() method)
For "unix": Display "Unix timestamp: [unix_seconds]" (use Unix() method)
Count and display statistics:
"=== TIME INFORMATION SUMMARY ==="
"Total formats displayed: [number_of_formats]"
"Timestamp components extracted: [number_of_unique_components]"
Display additional time details:
"Additional time details:"
"- Nanoseconds: [nanosecond_value]" (use Nanosecond() method)
"- Location: [timezone_location]" (use Location() method)
"- Is zero time: [true_or_false]" (use IsZero() method)
Display the completion message: "Time tracking completed - all requested formats processed"
Use the time package to work with time objects, the strings package to split the input string, and the fmt package for formatted output. Remember that time.Now() returns a time.Time object that contains all the information about the current moment, and you can extract specific components using the various methods available on this object. This challenge demonstrates how to capture the current time and extract different pieces of information from a single time.Time object.
