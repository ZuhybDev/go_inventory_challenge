package main

import (
	"strings"
)

func organizeGardenRevival(plotStatus []string, volunteerNames []string, toolInventory []string) string {
	var overgrownPlots []string
	var assignments []string

	// Extract overgrown plots
	for _, status := range plotStatus {
		if strings.Contains(status, "overgrown") {
			plotID := strings.Split(status, ":")[0]
			overgrownPlots = append(overgrownPlots, plotID)
		}
	}

	// Assign volunteers to overgrown plots using round-robin
	for i, plot := range overgrownPlots {
		volunteer := volunteerNames[i%len(volunteerNames)]
		assignment := plot + "-" + volunteer
		assignments = append(assignments, assignment)
	}

	// Create the formatted report
	result := "Overgrown plots: " + strings.Join(overgrownPlots, ",") + "\n"
	result += "Volunteer assignments: " + strings.Join(assignments, ",") + "\n"
	result += "Restock needed: " + strings.Join(toolInventory, ",")

	return result
}

func main() {
	plotStatus := []string{"Plot1:overgrown, Plot2:ready, Plot3:overgrown"}
	volunteerNames := []string{"Zuhaib", "Ahmed"}
	toolInventory := []string{"shovel, rake, hose"}
	organizeGardenRevival(plotStatus, volunteerNames, toolInventory)
}
