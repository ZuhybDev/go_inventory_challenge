package main

import (
	"fmt"
	"sort"
)

type PlantInfo struct {
	Name   string
	Height int
	method string
}

func organizeGardenTasks(plants []string, heights []int, methods []string) {
	// []string  return this
	// Write code here

	plantSlice := make([]PlantInfo, 0)

	for idx, plant := range plants {
		plantSlice = append(plantSlice, PlantInfo{Name: plant, Height: heights[idx], method: methods[idx]})
	}
	//sort by method
	sort.SliceStable(plantSlice, func(i, j int) bool {
		return plantSlice[i].method > plantSlice[j].method
	})

	//sort by height
	sort.SliceStable(plantSlice, func(i, j int) bool {
		return plantSlice[i].Height < plantSlice[j].Height
	})

	fmt.Println(plantSlice)

	final := make([]string, 0)

	for _, plant := range plantSlice {
		if plant.Height >= 30 {
			continue
		}
		if plant.Height <= 15 {
			final = append(
				final,
				fmt.Sprintf("%s: replace by %s,", plant.Name, plant.method),
			)
		} else {
			final = append(
				final,
				fmt.Sprintf("%s: support by %s,", plant.Name, plant.method),
			)
		}

	}

	// return final

}

func main() {
	plants := []string{"orchid", "lily", "ivy"}
	heights := []int{16, 35, 15}
	methods := []string{"prune", "trim", "prune"}
	organizeGardenTasks(plants, heights, methods)
}
