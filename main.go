package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/arkadiusreymond/CleanGoApp/delivery"
	"github.com/arkadiusreymond/CleanGoApp/repository"
	"github.com/arkadiusreymond/CleanGoApp/usecase"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	// Connect to the database
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/poultry")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize repository, use case, and handler
	birdRepo := repository.NewBirdRepository(db)
	birdUseCase := usecase.NewBirdUseCase(*birdRepo)
	birdHandler := delivery.NewBirdHandler(*birdUseCase)

	// Create a new router using gorilla/mux
	router := mux.NewRouter()

	// Register routes with the router
	router.HandleFunc("/birds", birdHandler.CreateBirdHandler).Methods("POST")
	router.HandleFunc("/birds/{id:[0-9]+}", birdHandler.GetBirdByIDHandler).Methods("GET")
	router.HandleFunc("/birds/{id:[0-9]+}", birdHandler.UpdateBirdHandler).Methods("PUT")
	router.HandleFunc("/birds/{id:[0-9]+}", birdHandler.DeleteBirdHandler).Methods("DELETE")
	router.HandleFunc("/birds", birdHandler.GetAllBirdsHandler).Methods("GET")

	// Start the server
	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
