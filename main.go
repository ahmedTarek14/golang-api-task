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


	// 🛠️ Define API Endpoints
	r.HandleFunc("/api/user/register", RegisterUser).Methods("POST")
	r.HandleFunc("/api/user/login", LoginUser).Methods("POST")

	r.HandleFunc("/api/user/add-credit-card", AuthMiddleware(AddCreditCard)).Methods("POST")
	r.HandleFunc("/api/user/delete-credit-card", AuthMiddleware(DeleteCreditCard)).Methods("DELETE")

	r.HandleFunc("/api/user/products", ListProducts).Methods("GET")

	r.HandleFunc("/api/user/products/buy", AuthMiddleware(BuyProducts)).Methods("POST")

	/* Print registered routes for test */
	// r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	// 	path, err := route.GetPathTemplate()
	// 	if err == nil {
	// 		fmt.Println("✅ Registered Route:", path)
	// 	}
	// 	return nil
	// })

	// Start the server
	fmt.Println("🚀 Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}