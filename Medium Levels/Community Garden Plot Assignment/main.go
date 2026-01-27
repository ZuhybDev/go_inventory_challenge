package main

import "fmt"

func assignGardenPlots(volunteers []string, preferences []int, availablePlots []int) {
	assigned := 0
	waiting := 0

	// Track available plots
	available := make(map[int]bool)
	for _, p := range availablePlots {
		available[p] = true
	}

	// Process each volunteer
	for i := range volunteers {
		plot := preferences[i]

		if available[plot] {
			assigned++
			delete(available, plot) // plot is now taken
		} else {
			waiting++
		}
	}

	fmt.Printf("Assigned: %d plots, Waiting: %d volunteers\n", assigned, waiting)
}

// return string

func main() {
	volunteers := []string{
		"Alice", "Ala", "jaama", "yaas",
	}

	preferences := []int{
		1, 2, 3, 4,
	}
	availablePlots := []int{
		1, 3, 5,
	}

	assignGardenPlots(volunteers, preferences, availablePlots)

}
