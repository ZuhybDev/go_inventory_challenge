Create a workout session timer that uses time.Duration to track different exercise intervals and rest periods. This challenge will test your ability to create duration values using Go's predefined constants and perform calculations with them.

You will receive two inputs:

A string containing exercise definitions in the format "exercise1:duration_value:duration_unit,exercise2:duration_value:duration_unit" (e.g., "Push-ups:30:Second,Plank:2:Minute,Squats:45:Second")
A string containing rest period information in the format "rest_duration_value:rest_duration_unit" (e.g., "15:Second")
Your task is to:

Parse the first input by splitting on commas to get individual exercise entries
For each exercise entry, split on colons to extract the exercise name, duration value (as string), and duration unit
Parse the second input by splitting on colons to get rest duration value and unit
Convert duration values from strings to integers and create time.Duration objects using the appropriate constants:
For "Second": multiply by time.Second
For "Minute": multiply by time.Minute
For "Hour": multiply by time.Hour
Display the application header: "=== WORKOUT SESSION TIMER ==="
Display the session setup: "Setting up workout session..."
For each exercise, display: "Exercise: [exercise_name] - Duration: [duration_object]"
Display the rest period: "Rest period between exercises: [rest_duration_object]"
Calculate and display session breakdown:
"=== SESSION BREAKDOWN ==="
For each exercise: "[exercise_name]: [duration_in_seconds] seconds"
"Rest periods: [rest_duration_in_seconds] seconds each"
Calculate total session time (sum of all exercise durations plus rest periods between exercises):
"=== TIMING CALCULATIONS ==="
"Total exercise time: [total_exercise_duration_object]"
"Total rest time: [total_rest_duration_object]" (rest periods between exercises, so one less than number of exercises)
"Complete session duration: [total_session_duration_object]"
Display session statistics:
"=== SESSION STATISTICS ==="
"Number of exercises: [number_of_exercises]"
"Longest exercise: [longest_exercise_name] ([longest_duration_object])"
"Shortest exercise: [shortest_exercise_name] ([shortest_duration_object])"
"Average exercise duration: [average_duration_in_seconds] seconds"
Display time conversions:
"=== TIME CONVERSIONS ==="
"Session duration in minutes: [total_minutes]" (convert total session duration to minutes)
"Session duration in hours: [total_hours]" (convert total session duration to hours, show as decimal)
Display the completion message: "Workout session timer configured successfully"
Use the time package to work with durations, the strings package to split input strings, the strconv package to convert string numbers to integers, and the fmt package for formatted output. Remember that time.Duration values can be added together and converted to different units using methods like Seconds(), Minutes(), and Hours(). This challenge demonstrates how to create and manipulate duration values for timing applications.
