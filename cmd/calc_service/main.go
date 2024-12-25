package main

import (
	"fmt"
	"net/http"

	"calc_service/pkg/handlers"
)

func main() {
	http.HandleFunc("/api/v1/calculate", handlers.CalculateHandler)
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
