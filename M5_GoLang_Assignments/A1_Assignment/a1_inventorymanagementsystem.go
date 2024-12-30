package main

import (
	"errors"
	"fmt"
	"sort"
)

// Product struct to represent a product in the inventory
type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

// Slice to store all products in the inventory
var inventory []Product

// Function to add a new product to the inventory
func addProduct(id int, name string, price interface{}, stock int) error {
	// Type casting for price to float64
	priceFloat, ok := price.(float64)
	if !ok {
		return errors.New("invalid price type; must be a float64")
	}

	// Validate that stock is non-negative
	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	// Add the new product to the inventory
	inventory = append(inventory, Product{ID: id, Name: name, Price: priceFloat, Stock: stock})
	return nil
}

// Function to update the stock of a product
func updateStock(id int, newStock int) error {
	if newStock < 0 {
		return errors.New("stock cannot be negative")
	}

	// Find the product and update stock
	for i, p := range inventory {
		if p.ID == id {
			inventory[i].Stock = newStock
			return nil
		}
	}
	return errors.New("product not found")
}

// Function to search for a product by ID or Name
func searchProduct(query interface{}) (*Product, error) {
	// Search by ID or name
	for _, p := range inventory {
		switch v := query.(type) {
		case int:
			if p.ID == v {
				return &p, nil
			}
		case string:
			if p.Name == v {
				return &p, nil
			}
		}
	}
	return nil, errors.New("product not found")
}

// Function to display the inventory in a formatted table
func displayInventory() {
	fmt.Println("ID\tName\t\tPrice\tStock")
	for _, p := range inventory {
		fmt.Printf("%d\t%s\t\t$%.2f\t%d\n", p.ID, p.Name, p.Price, p.Stock)
	}
}

// Function to sort products by price
func sortByPrice() {
	sort.SliceStable(inventory, func(i, j int) bool {
		return inventory[i].Price < inventory[j].Price
	})
}

// Function to sort products by stock
func sortByStock() {
	sort.SliceStable(inventory, func(i, j int) bool {
		return inventory[i].Stock < inventory[j].Stock
	})
}

func main() {
	// Adding some sample products to the inventory
	err := addProduct(1, "Laptop", 799.99, 10)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = addProduct(2, "Smartphone", 499.49, 20)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = addProduct(3, "Headphones", 150.00, 50)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Display the initial inventory
	fmt.Println("Initial Inventory:")
	displayInventory()

	// Update stock of a product
	err = updateStock(2, 18)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("\nUpdated Inventory (Stock of product 2 updated):")
	displayInventory()

	// Search for a product
	product, err := searchProduct(2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("\nProduct found: %+v\n", *product)
	}

	// Search for a product by name
	product, err = searchProduct("Headphones")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("\nProduct found: %+v\n", *product)
	}

	// Sort and display inventory by price
	sortByPrice()
	fmt.Println("\nInventory Sorted by Price:")
	displayInventory()

	// Sort and display inventory by stock
	sortByStock()
	fmt.Println("\nInventory Sorted by Stock:")
	displayInventory()
}
