package main

import (
	"context"
	"encoding/json"
	"net/http"
)

// CreditCard - Struct for credit card data
type CreditCard struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	CardNumber string `json:"card_number"`
	ExpiryDate string `json:"expiry_date"`
	CVV        string `json:"cvv"`
}

// AddCreditCard - API to add a credit card
func AddCreditCard(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from context
	userID, ok := r.Context().Value("user_id").(int)
	if !ok {
		http.Error(w, "❌ Unauthorized access", http.StatusUnauthorized)
		return
	}

	var card CreditCard
	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		http.Error(w, "❌ Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if card.CardNumber == "" || len(card.CVV) != 3 {
		http.Error(w, "❌ Invalid card details", http.StatusBadRequest)
		return
	}

	// Save the credit card in the database
	query := `INSERT INTO credit_cards (user_id, card_number, expiry_date, cvv) VALUES ($1, $2, $3, $4) RETURNING id`
	err = db.QueryRow(context.Background(), query, userID, card.CardNumber, card.ExpiryDate, card.CVV).Scan(&card.ID)
	if err != nil {
		http.Error(w, "❌ Failed to add credit card", http.StatusInternalServerError)
		return
	}

	// Remove CVV from response for security reasons
	card.CVV = ""

	// Set Content-Type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(card)
}


func DeleteCreditCard(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from context
	userID := r.Context().Value("user_id").(int)

	// Decode request body properly
	var requestData struct {
		CardID int `json:"card_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "❌ Invalid request body", http.StatusBadRequest)
		return
	}

	// Delete the card from database
	query := `DELETE FROM credit_cards WHERE id = $1 AND user_id = $2`
	result, err := db.Exec(context.Background(), query, requestData.CardID, userID)
	if err != nil {
		http.Error(w, "❌ Failed to delete credit card", http.StatusInternalServerError)
		return
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "❌ No credit card found to delete", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("✅ Credit card deleted successfully"))
}