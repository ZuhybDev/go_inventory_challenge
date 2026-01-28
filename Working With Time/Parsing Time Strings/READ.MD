Create a log file parser that converts timestamp strings from various formats into time.Time objects for analysis. This challenge will test your ability to use time.Parse with different layout patterns to handle multiple date and time formats commonly found in system logs.

Important Note: Since log messages and layout patterns may contain spaces, you'll need to use bufio.Scanner to read entire lines of input. The fmt.Scanln function stops reading at the first space, which would break log entries and layouts with spaces. Use bufio.NewScanner(os.Stdin) and call scanner.Scan() followed by scanner.Text() to read complete lines.

You will receive two inputs:

A string containing log entries in the format "timestamp1:message1,timestamp2:message2,timestamp3:message3" (e.g., "2023-10-26:System startup,2023/12/25:Holiday shutdown,Jan 15 2024:Maintenance window")
A string containing layout patterns separated by commas in the format "layout1,layout2,layout3" (e.g., "2006-01-02,2006/01/02,Jan 2 2006")
Your task is to:

Use bufio.Scanner to read the input lines properly to handle spaces in log messages and layout patterns
Parse the first input by splitting on commas to get individual log entries
For each log entry, split on colons to extract the timestamp string and message
Parse the second input by splitting on commas to get layout patterns
Display the application header: "=== LOG TIMESTAMP PARSER ==="
Display the parsing status: "Processing log entries with multiple timestamp formats..."
For each log entry, attempt to parse the timestamp using each layout pattern until one succeeds:
Try each layout pattern with time.Parse until parsing succeeds (no error)
When successful, display: "Parsed: [timestamp_string] -> [parsed_time_object] (using layout: [successful_layout])"
Also display: " Message: [log_message]"
If no layout works for a timestamp, display: "Failed to parse: [timestamp_string] - no matching layout found"
After processing all entries, display parsing statistics:
"=== PARSING RESULTS ==="
"Total log entries processed: [number_of_entries]"
"Successfully parsed timestamps: [number_of_successful_parses]"
"Failed parsing attempts: [number_of_failed_parses]"
Display layout usage statistics:
"Layout pattern usage:"
For each layout pattern: " [layout_pattern]: [number_of_times_used] times"
If any timestamps were successfully parsed, display chronological analysis:
"=== CHRONOLOGICAL ANALYSIS ==="
"Earliest timestamp: [earliest_parsed_time]"
"Latest timestamp: [latest_parsed_time]"
"Time span covered: [time_difference_description]"
Display the completion message: "Log parsing completed - timestamps converted to time objects"
Use the time package for parsing operations, the strings package to split input strings, bufio and os packages for proper input handling, and the fmt package for formatted output. Remember that time.Parse returns both a time.Time object and an error - you must check the error to determine if parsing was successful. When parsing fails with one layout, try the next layout pattern until you find one that works or exhaust all options. This challenge demonstrates how to handle multiple timestamp formats in real-world log processing scenarios.
