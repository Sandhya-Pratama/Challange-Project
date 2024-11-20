package main

import (
	"fmt"
	"net/http"

	"github.com/Sandhya-Pratama/TEST-BACK-END_Sandhya/data"
	"github.com/Sandhya-Pratama/TEST-BACK-END_Sandhya/routes"
	"github.com/Sandhya-Pratama/TEST-BACK-END_Sandhya/services"
)

func main() {
	// Load data dari file
	data, err := data.LoadData("./data/data.txt")
	if err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

	// Inisialisasi data ke memory
	services.InitData(data)

	// Register routes
	routes.RegisterRoutes()

	// Jalankan server
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
