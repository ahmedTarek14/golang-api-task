package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Secret Key for JWT
var jwtKey = []byte("your_secret_key")

// User - Struct for user data
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

// RegisterUser - Create a new user
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "❌ Invalid request body", http.StatusBadRequest)
		return
	}

	// Encrypt the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "❌ Error encrypting password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Save the user in the database
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	err = db.QueryRow(context.Background(), query, user.Name, user.Email, string(hashedPassword)).Scan(&user.ID)
	if err != nil {
		http.Error(w, "❌ Failed to register user", http.StatusInternalServerError)
		return
	}

	// 🛡️ إنشاء التوكين
	expirationTime := time.Now().Add(24 * time.Hour) // التوكين صالح لمدة 24 ساعة
	claims := &jwt.MapClaims{
		"user_id": user.ID,
		"exp":     expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "❌ Could not generate token", http.StatusInternalServerError)
		return
	}

	// // إزالة كلمة المرور من الاستجابة
	// user.Password = "Encrypted"

	// ✅ إرسال الاستجابة مع التوكين
	response := map[string]interface{}{
		"token": tokenString,
		"user":  user,
	}

	// Send the response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// LoginUser - تسجيل الدخول
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "❌ Invalid request body", http.StatusBadRequest)
		return
	}

	var dbUser User
	query := `SELECT id, name, email, password FROM users WHERE email = $1`
	err = db.QueryRow(context.Background(), query, user.Email).Scan(&dbUser.ID, &dbUser.Name, &dbUser.Email, &dbUser.Password)
	if err != nil {
		http.Error(w, "❌ Invalid email or password", http.StatusUnauthorized)
		return
	}

	// التحقق من كلمة المرور
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "❌ Invalid email or password", http.StatusUnauthorized)
		return
	}

	// 🛡️ إنشاء التوكين
	expirationTime := time.Now().Add(24 * time.Hour) // التوكين صالح لمدة 24 ساعة
	claims := &jwt.MapClaims{
		"user_id": dbUser.ID,
		"exp":     expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "❌ Could not generate token", http.StatusInternalServerError)
		return
	}

	// إزالة كلمة المرور من الاستجابة
	dbUser.Password = ""

	// ✅ إرسال الاستجابة مع التوكين
	response := map[string]interface{}{
		"token": tokenString,
		"user":  dbUser,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}