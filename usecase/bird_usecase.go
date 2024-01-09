// usecase/bird_usecase.go
package usecase

import "github.com/arkadiusreymond/CleanGoApp/repository"

type BirdUseCase struct {
	BirdRepo repository.BirdRepository
}

func NewBirdUseCase(birdRepo repository.BirdRepository) *BirdUseCase {
	return &BirdUseCase{BirdRepo: birdRepo}
}

func (uc *BirdUseCase) CreateBird(name, color string) error {
	// Implement the business logic for creating a new bird
	newBird := &repository.Bird{
		Name:  name,
		Color: color,
	}
	return uc.BirdRepo.Create(newBird)
}

func (uc *BirdUseCase) GetBirdByID(id int) (*repository.Bird, error) {
	// Implement the business logic for getting a bird by ID
	return uc.BirdRepo.GetByID(id)
}

func (uc *BirdUseCase) UpdateBird(bird *repository.Bird) error {
	// Implement the business logic for updating a bird
	return uc.BirdRepo.Update(bird)
}

func (uc *BirdUseCase) DeleteBird(id int) error {
	// Implement the business logic for deleting a bird by ID
	return uc.BirdRepo.Delete(id)
}

func (uc *BirdUseCase) GetAllBirds() ([]*repository.Bird, error) {
	// Implement the business logic for getting all birds
	return uc.BirdRepo.GetAll()
}
