package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to database
	connectDB()
	// create tables
	CreateTables()

	// Setting up the Router
	r := mux.NewRouter()
	// 🛠️ تعريف الـ API Endpoints
	r.HandleFunc("/api/register", RegisterUser).Methods("POST")
	r.HandleFunc("/api/login", LoginUser).Methods("POST")

	
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "🚀 Server is running!")
	})

	// تشغيل السيرفر
	fmt.Println("🚀 Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}