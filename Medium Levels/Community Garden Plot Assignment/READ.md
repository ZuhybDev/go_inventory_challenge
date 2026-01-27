Write a function assignGardenPlots that takes volunteers, preferences, availablePlots and returns a formatted assignment report string.

The function processes volunteer plot assignments for the Columbus community garden, matching volunteer preferences with available plots and generating a summary report.

Logic:

Initialize counters for assigned plots and waiting volunteers
For each volunteer, check if their preferred plot is available
If available, assign the plot (remove from available list) and increment assigned counter
If not available, increment waiting counter
Use logical operators to validate assignments and string manipulation to format the final report
Parameters:

volunteers ([]string): List of volunteer names
preferences ([]int): Preferred plot numbers for each volunteer (same order as volunteers)
availablePlots ([]int): List of currently available plot numbers
Returns: Assignment summary report. Format: Assigned: X plots, Waiting: Y volunteers
