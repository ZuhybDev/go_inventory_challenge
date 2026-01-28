Write a function organizeProjectTasks that takes tasks, durations, costs, and budget and returns a string array of formatted task summaries.

The function filters out tasks that exceed the budget, sorts remaining tasks by duration (shortest first), and formats each as a project summary.

Logic:

Filter tasks where cost is within budget
Sort filtered tasks by duration in ascending order
Format each task as "TaskName: Durationh, $Cost"
Parameters:

tasks ([]string): Task names for the home project
durations ([]float64): Time required for each task in hours
costs ([]int): Cost of each task in dollars
budget (int): Maximum budget available
Returns: String array of formatted task summaries sorted by duration. Format: ["Build shelf: 2.5h, $30", "Install fixture: 4h, $45"]