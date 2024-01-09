// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/arkadiusreymond/CleanGoApp/delivery"
	"github.com/arkadiusreymond/CleanGoApp/repository"
	"github.com/arkadiusreymond/CleanGoApp/usecase"
)

func main() {
	// Connect to the database
	db, err := sql.Open("mysql", "your_mysql_user:your_mysql_password@tcp(your_mysql_host:your_mysql_port)/your_mysql_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize repository, use case, and handler
	birdRepo := repository.NewBirdRepository(db)
	birdUseCase := usecase.NewBirdUseCase(*birdRepo)
	birdHandler := delivery.NewBirdHandler(*birdUseCase)

	// Set up HTTP routes
	http.HandleFunc("/birds", birdHandler.CreateBirdHandler)
	http.HandleFunc("/birds/{id}", birdHandler.GetBirdByIDHandler)
	http.HandleFunc("/birds/{id}", birdHandler.UpdateBirdHandler)
	http.HandleFunc("/birds/{id}", birdHandler.DeleteBirdHandler)
	http.HandleFunc("/birds", birdHandler.GetAllBirdsHandler)

	// Start the server
	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
