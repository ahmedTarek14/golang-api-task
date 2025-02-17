package main

import (
	"context"
	"encoding/json"
	"net/http"
)

// BuyRequest - Struct for the buy request
type BuyRequest struct {
	ProductIDs []int `json:"product_ids"`
}

// BuyProducts - API to allow users to buy multiple products
func BuyProducts(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from the middleware
	userID := r.Context().Value("user_id").(int)

	var req BuyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || len(req.ProductIDs) == 0 {
		http.Error(w, "❌ Invalid request body", http.StatusBadRequest)
		return
	}

	// Insert purchase records into the database
	query := `INSERT INTO purchases (user_id, product_id) VALUES ($1, $2)`
	for _, productID := range req.ProductIDs {
		_, err := db.Exec(context.Background(), query, userID, productID)
		if err != nil {
			http.Error(w, "❌ Failed to complete purchase", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("✅ Purchase completed successfully"))
}
