Create a project deadline calculator that uses time arithmetic to determine project milestones and deadlines. This challenge will test your ability to use the Add and Sub methods on time.Time objects to calculate future dates and time differences between events.

You will receive two inputs:

A string containing project start information in the format "year:month:day:hour:minute" (e.g., "2024:1:15:9:0")
A string containing task durations in the format "task1:duration_value:duration_unit,task2:duration_value:duration_unit" (e.g., "Planning:5:Day,Development:3:Week,Testing:10:Day")
Your task is to:

Parse the first input by splitting on colons to extract year, month, day, hour, and minute values
Convert these string values to integers and create a project start time using time.Date (use time.UTC for location, 0 for seconds and nanoseconds)
Parse the second input by splitting on commas to get individual task entries
For each task entry, split on colons to extract task name, duration value, and duration unit
Convert duration values to integers and create time.Duration objects using appropriate constants:
For "Day": multiply by 24 _ time.Hour
For "Week": multiply by 7 _ 24 \* time.Hour
For "Hour": multiply by time.Hour
Display the application header: "=== PROJECT DEADLINE CALCULATOR ==="
Display the project start: "Project start time: [start_time_object]"
Calculate task deadlines by adding each task duration to the previous task's end time (first task starts at project start time):
For each task, display:
"Task: [task_name]"
" Duration: [duration_object]"
" Start: [task_start_time]"
" End: [task_end_time]" (calculated using Add method)
Calculate and display project timeline:
"=== PROJECT TIMELINE ==="
"Project start: [start_time_object]"
"Project end: [final_task_end_time]"
"Total project duration: [total_duration_object]" (calculated using Sub method between project end and start)
Display task analysis:
"=== TASK ANALYSIS ==="
"Number of tasks: [number_of_tasks]"
"Longest task: [longest_task_name] ([longest_duration_object])"
"Shortest task: [shortest_task_name] ([shortest_duration_object])"
Calculate time between consecutive tasks and display gaps:
"=== TASK TRANSITIONS ==="
For each pair of consecutive tasks: "[first_task_name] to [second_task_name]: [time_difference_object]" (should be 0 since tasks are consecutive)
Display project statistics:
"=== PROJECT STATISTICS ==="
"Total working days: [total_duration_in_hours_divided_by_24]"
"Total working hours: [total_duration_in_hours]"
"Project spans [number_of_calendar_days] calendar days"
Display the completion message: "Project timeline calculation completed successfully"
Use the time package for time operations, the strings package to split input strings, the strconv package to convert strings to integers, and the fmt package for formatted output. Remember that the Add method returns a new time.Time object representing the original time plus the duration, and the Sub method returns a time.Duration representing the difference between two times. This challenge demonstrates how to build a complete project timeline using time arithmetic operations.
