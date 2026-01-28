package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Product struct {
	Price    float64
	Quantity int
}

func checkStock(inventory map[string]Product, productName string) (int, error) {
	product, exists := inventory[productName]
	if !exists {
		return 0, errors.New("product not found")
	}
	return product.Quantity, nil
}

func addNewItem(inventory map[string]Product, name string, price float64, quantity int) error {
	if _, exists := inventory[name]; exists {
		return errors.New("product already exists")
	}
	inventory[name] = Product{Price: price, Quantity: quantity}
	return nil
}

func updateStock(inventory map[string]Product, productName string, change int) error {
	product, exists := inventory[productName]
	if !exists {
		return errors.New("product not found")
	}
	newQuantity := product.Quantity + change
	if newQuantity < 0 {
		return errors.New("insufficient stock")
	}
	product.Quantity = newQuantity
	inventory[productName] = product
	return nil
}

func generateReport(inventory map[string]Product, reportType string, threshold int) {
	var products []string
	for name := range inventory {
		products = append(products, name)
	}
	sort.Strings(products)

	fmt.Printf("=== %s REPORT ===\n", strings.ToUpper(reportType))

	if reportType == "low" {
		fmt.Printf("Products with stock below %d:\n", threshold)
		for _, name := range products {
			product := inventory[name]
			if product.Quantity < threshold {
				fmt.Printf("- %s: %d units (Price: $%.2f)\n", name, product.Quantity, product.Price)
			}
		}
	} else if reportType == "full" {
		fmt.Printf("Complete inventory (minimum threshold: %d):\n", threshold)
		for _, name := range products {
			product := inventory[name]
			status := "OK"
			if product.Quantity < threshold {
				status = "LOW STOCK"
			}
			fmt.Printf("- %s: %d units @ $%.2f each [%s]\n", name, product.Quantity, product.Price, status)
		}
	}
}

func main() {
	var initialData string
	var operations string
	var parameters string

	fmt.Scanln(&initialData)
	fmt.Scanln(&operations)
	fmt.Scanln(&parameters)

	inventory := make(map[string]Product)

	// Parse initial inventory data
	items := strings.Split(initialData, ",")
	for _, item := range items {
		parts := strings.Split(item, ":")
		name := parts[0]
		price, _ := strconv.ParseFloat(parts[1], 64)
		quantity, _ := strconv.Atoi(parts[2])
		inventory[name] = Product{Price: price, Quantity: quantity}
	}

	// Display startup message
	fmt.Println("=== INVENTORY MANAGEMENT SYSTEM ===")
	fmt.Printf("System initialized with %d products\n", len(inventory))
	fmt.Println("Starting interactive session...")

	// Parse operations and parameters
	opList := strings.Split(operations, ",")
	paramList := strings.Split(parameters, "|")

	// Process each operation
	for i, operation := range opList {
		switch operation {
		case "check":
			fmt.Println("--- STOCK CHECK ---")
			productName := paramList[i]
			fmt.Printf("Checking stock for: %s\n", productName)
			quantity, err := checkStock(inventory, productName)
			if err != nil {
				fmt.Println("Product not found in inventory")
			} else {
				fmt.Printf("Stock level: %d units\n", quantity)
			}

		case "add":
			fmt.Println("--- ADD ITEM ---")
			parts := strings.Split(paramList[i], ":")
			name := parts[0]
			price, _ := strconv.ParseFloat(parts[1], 64)
			quantity, _ := strconv.Atoi(parts[2])
			fmt.Printf("Adding new product: %s\n", name)
			err := addNewItem(inventory, name, price, quantity)
			if err != nil {
				fmt.Printf("Failed to add product: %s\n", err.Error())
			} else {
				fmt.Println("Product added successfully")
			}

		case "update":
			fmt.Println("--- UPDATE STOCK ---")
			parts := strings.Split(paramList[i], ":")
			productName := parts[0]
			change, _ := strconv.Atoi(parts[1])
			fmt.Printf("Updating stock for: %s\n", productName)
			err := updateStock(inventory, productName, change)
			if err != nil {
				fmt.Printf("Update failed: %s\n", err.Error())
			} else {
				newQuantity := inventory[productName].Quantity
				if change >= 0 {
					fmt.Printf("Added %d units. New stock: %d\n", change, newQuantity)
				} else {
					fmt.Printf("Removed %d units. New stock: %d\n", -change, newQuantity)
				}
			}

		case "report":
			fmt.Println("--- GENERATE REPORT ---")
			parts := strings.Split(paramList[i], ",")
			reportType := parts[0]
			threshold, _ := strconv.Atoi(parts[1])
			fmt.Printf("Generating %s report with threshold %d\n", reportType, threshold)
			generateReport(inventory, reportType, threshold)

		case "exit":
			fmt.Println("--- SYSTEM EXIT ---")
			totalProducts := len(inventory)
			totalItems := 0
			totalValue := 0.0
			for _, product := range inventory {
				totalItems += product.Quantity
				totalValue += float64(product.Quantity) * product.Price
			}
			fmt.Println("Final inventory status:")
			fmt.Printf("Total products: %d\n", totalProducts)
			fmt.Printf("Total items: %d\n", totalItems)
			fmt.Printf("Total value: $%.2f\n", totalValue)
			fmt.Println("Session completed successfully")
			fmt.Println("Thank you for using the Inventory Management System")
			return
		}

		if operation != "exit" {
			fmt.Println("Operation completed. Continuing to next operation...")
		}
	}
}
