package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware - Middleware to validate JWT tokens
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "🚫 Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Extract token from "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // No "Bearer " prefix found
			http.Error(w, "🚫 Invalid token format", http.StatusUnauthorized)
			return
		}

		// Parse and validate token
		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		// Handle token errors
		if err != nil || !token.Valid {
			http.Error(w, "🚫 Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Store user ID in request context for further use
		userID := (*claims)["user_id"]
		fmt.Println("✅ Authenticated User ID:", userID)

		// Proceed to the next handler
		next(w, r)
	}
}