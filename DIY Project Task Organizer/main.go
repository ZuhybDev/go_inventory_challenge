package main

import (
	"fmt"
	"sort"
	"strings"
)

type TaskInfo struct {
	Name     string
	Cost     int
	Duration float64
}

func organizeProjectTasks(tasks []string, durations []float64, costs []int, budget int) []string {

	// Write code here
	taskWithCost := make(map[string]TaskInfo)

	for i, t := range tasks {
		taskWithCost[t] = TaskInfo{Name: t, Cost: costs[i], Duration: durations[i]}

	}

	//remove it later
	// fmt.Println(taskWithCost)
	result := make([]TaskInfo, 0)
	for _, task := range tasks {
		info := taskWithCost[task]

		if info.Cost > budget {
			continue
		}
		result = append(result, TaskInfo{Name: task, Cost: info.Cost, Duration: info.Duration})

	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Duration < result[j].Duration
	})

	formatted := []string{}

	for _, t := range result {
		formatted = append(formatted, fmt.Sprintf("%s: %.1gh, $%d", t.Name, t.Duration, t.Cost))
	}

	final := strings.Join(formatted, ",")
	return []string{final}
}

func main() {
	taskNames := []string{"Build shelf", "Install fixture", "Paint wall"}
	taskDurations := []float64{2.5, 4, 1.5}
	taskCosts := []int{30, 45, 60}
	projectBudget := 50

	// 2. Call the function, passing the variables as arguments.
	organizeProjectTasks(taskNames, taskDurations, taskCosts, projectBudget)
}

/*
[Paint] --> task
[1.5] --> duration
[50] --> cost
50 --> budget


[Build shelf, Install fixture, Paint wall]
[2.5, 4, 1.5]
[30, 45, 60] ---> 60 > 50 soo its out
50
*/
