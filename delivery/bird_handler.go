// delivery/bird_handler.go
package delivery

import (
	"encoding/json"
	"github.com/arkadiusreymond/CleanGoApp/repository"
	"github.com/arkadiusreymond/CleanGoApp/usecase"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	// Extract bird ID from URL path parameter using gorilla/mux
	params := mux.Vars(r)
	idStr, ok := params["id"]
	if !ok {
		http.Error(w, "Invalid request. Missing bird ID.", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid bird ID. Must be an integer.", http.StatusBadRequest)
		return
	}

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
	// Extract bird ID from URL path parameter using gorilla/mux
	params := mux.Vars(r)
	idStr, ok := params["id"]
	if !ok {
		http.Error(w, "Invalid request. Missing bird ID.", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid bird ID. Must be an integer.", http.StatusBadRequest)
		return
	}

	// Parse request body for updated bird data
	var updatedBird repository.Bird
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedBird); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Set the ID of the bird to be updated
	updatedBird.ID = id

	// Call BirdUseCase.UpdateBird with the updated bird data
	if err := h.BirdUseCase.UpdateBird(&updatedBird); err != nil {
		http.Error(w, "Failed to update bird", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BirdHandler) DeleteBirdHandler(w http.ResponseWriter, r *http.Request) {
	// Extract bird ID from URL path parameter
	params := mux.Vars(r)
	idStr, ok := params["id"]
	if !ok {
		http.Error(w, "Invalid request. Missing bird ID.", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid bird ID. Must be an integer.", http.StatusBadRequest)
		return
	}

	// Call BirdUseCase.DeleteBird with the ID of the bird to be deleted
	if err := h.BirdUseCase.DeleteBird(id); err != nil {
		http.Error(w, "Failed to delete bird", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
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
