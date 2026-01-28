package main

import (
	"fmt"
	"sort"
)

type Item struct {
	name  string
	price float64
}

func formatPriceTags(itemNames []string, baseCosts []float64, markupRates []float64) []string {
	// Create items with calculated final prices
	items := make([]Item, len(itemNames))
	for i := 0; i < len(itemNames); i++ {
		finalPrice := baseCosts[i] * (1 + markupRates[i])
		items[i] = Item{name: itemNames[i], price: finalPrice}
	}

	// Sort items by price in ascending order
	sort.Slice(items, func(i, j int) bool {
		return items[i].price < items[j].price
	})

	// Format price tags
	result := make([]string, len(items))
	for i, item := range items {
		result[i] = fmt.Sprintf("%-20s$%.2f", item.name, item.price)
	}

	return result
}
func main() {

	itemNames := []string{"Tire Rotation", "Brake Check"}
	baseCosts := []float64{25, 60}
	markupRates := []float64{0.2, 0.1}

	formatPriceTags(itemNames, baseCosts, markupRates)

}
