Create a historical event timeline application that uses time.Date to represent specific dates and times from history. This challenge will test your ability to create time.Time objects for precise historical moments and display them in an organized format.

You will receive two inputs:

A string containing historical events in the format "event1:year:month:day:hour:minute,event2:year:month:day:hour:minute" (e.g., "Moon Landing:1969:7:20:20:17,Berlin Wall Fall:1989:11:9:21:0")
A string containing display options separated by commas in the format "option1,option2,option3" (e.g., "chronological,details,summary")
Your task is to:

Parse the first input by splitting on commas to get individual event entries
For each event entry, split on colons to extract the event name, year, month, day, hour, and minute
Convert the numeric strings to integers and create a time.Time object using time.Date for each event (use time.UTC for the location parameter and 0 for seconds and nanoseconds)
Display the application header: "=== HISTORICAL TIMELINE APPLICATION ==="
Display the events loading confirmation: "Loading historical events..."
For each event, display: "Loaded: [event_name] - [complete_time_object]"
Parse the second input by splitting on commas to get display options
Process each display option:
For "chronological": Display "=== CHRONOLOGICAL ORDER ===" followed by events sorted by date (earliest first), showing "[event_name]: [year]-[month]-[day] at [hour]:[minute] UTC"
For "details": Display "=== DETAILED INFORMATION ===" followed by each event showing:
"Event: [event_name]"
" Date: [year]-[month]-[day]"
" Time: [hour]:[minute] UTC"
" Weekday: [weekday_name]"
" Unix timestamp: [unix_seconds]"
For "summary": Display "=== TIMELINE SUMMARY ===" followed by:
"Total events: [number_of_events]"
"Earliest event: [earliest_event_name] ([year])"
"Latest event: [latest_event_name] ([year])"
"Time span: [years_between_earliest_and_latest] years"
For "years": Display "=== EVENTS BY DECADE ===" followed by events grouped by decade, showing "[decade]s: [comma_separated_event_names]"
Display the completion message: "Historical timeline processing completed"
Use the time package to create specific dates with time.Date, the strings package to split input strings, the strconv package to convert string numbers to integers, and the fmt package for formatted output. Remember that time.Date takes parameters in the exact order: year, month, day, hour, minute, second, nanosecond, location. For sorting chronologically, you can compare time.Time objects using the Before() method. This challenge demonstrates how to create specific moments in time and work with historical data using Go's time functionality.
