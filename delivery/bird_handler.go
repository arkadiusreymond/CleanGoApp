// delivery/bird_handler.go
package delivery

import (
	"encoding/json"
	"github.com/arkadiusreymond/CleanGoApp/usecase"
	"net/http"
)

type BirdHandler struct {
	BirdUseCase usecase.BirdUseCase
}

func NewBirdHandler(birdUseCase usecase.BirdUseCase) *BirdHandler {
	return &BirdHandler{BirdUseCase: birdUseCase}
}

func (h *BirdHandler) CreateBirdHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	err := h.BirdUseCase.CreateBird(request.Name, request.Color)
	if err != nil {
		http.Error(w, "Failed to create bird2", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *BirdHandler) GetBirdByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Extract bird ID from URL path parameter
	// For example: /birds/1
	// ID is 1 in this case
	// Handle this using a router library like gorilla/mux or chi
	// For simplicity, I'm using a basic example here

	id := 1 // Replace this with the actual ID extraction logic

	bird, err := h.BirdUseCase.GetBirdByID(id)
	if err != nil {
		http.Error(w, "Failed to get bird by ID", http.StatusInternalServerError)
		return
	}

	if bird == nil {
		http.Error(w, "Bird not found", http.StatusNotFound)
		return
	}

	jsonResponse(w, bird)
}

func (h *BirdHandler) UpdateBirdHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to update a bird via HTTP
	// Similar to CreateBirdHandler, extract data from the request and call BirdUseCase.UpdateBird
}

func (h *BirdHandler) DeleteBirdHandler(w http.ResponseWriter, r *http.Request) {
	// Extract bird ID from URL path parameter
	// Implement the logic to delete a bird by ID via HTTP
	// Similar to GetBirdByIDHandler, extract ID and call BirdUseCase.DeleteBird
}

func (h *BirdHandler) GetAllBirdsHandler(w http.ResponseWriter, r *http.Request) {
	birds, err := h.BirdUseCase.GetAllBirds()
	if err != nil {
		http.Error(w, "Failed to get all birds", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, birds)
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(data); err != nil {
		http.Error(w, "Failed to encode response as JSON", http.StatusInternalServerError)
		return
	}
}
