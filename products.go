package main

import (
	"context"
	"encoding/json"
	"net/http"
)

// Product - Struct for product data
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// ListProducts - API to get all available products
func ListProducts(w http.ResponseWriter, r *http.Request) {
	var products []Product

	query := `SELECT id, name, price, quantity FROM products`
	rows, err := db.Query(context.Background(), query)
	if err != nil {
		http.Error(w, "❌ Failed to fetch products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
		if err != nil {
			http.Error(w, "❌ Error scanning product", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}