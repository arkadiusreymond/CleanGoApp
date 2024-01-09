// delivery/bird_handler.go
package delivery

import (
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
	// Implement the logic to handle the creation of a new bird via HTTP
}

func (h *BirdHandler) GetBirdByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to handle retrieving a bird by ID via HTTP
}

func (h *BirdHandler) UpdateBirdHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to handle updating a bird via HTTP
}

func (h *BirdHandler) DeleteBirdHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to handle deleting a bird by ID via HTTP
}

func (h *BirdHandler) GetAllBirdsHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to handle retrieving all birds via HTTP
}
