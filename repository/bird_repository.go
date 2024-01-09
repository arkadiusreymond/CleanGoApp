// repository/bird_repository.go
package repository

import "database/sql"

type BirdRepository struct {
	DB *sql.DB
}

func NewBirdRepository(db *sql.DB) *BirdRepository {
	return &BirdRepository{DB: db}
}

func (r *BirdRepository) Create(bird *Bird) error {
	_, err := r.DB.Exec("INSERT INTO birds (name, color) VALUES (?, ?)", bird.Name, bird.Color)
	if err != nil {
		return err
	}
	return nil
}

func (r *BirdRepository) GetByID(id int) (*Bird, error) {
	row := r.DB.QueryRow("SELECT id, name, color FROM birds WHERE id = ?", id)
	bird := &Bird{}
	err := row.Scan(&bird.ID, &bird.Name, &bird.Color)
	if err != nil {
		return nil, err
	}
	return bird, nil
}

func (r *BirdRepository) Update(bird *Bird) error {
	_, err := r.DB.Exec("UPDATE birds SET name = ?, color = ? WHERE id = ?", bird.Name, bird.Color, bird.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *BirdRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM birds WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *BirdRepository) GetAll() ([]*Bird, error) {
	rows, err := r.DB.Query("SELECT id, name, color FROM birds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var birds []*Bird
	for rows.Next() {
		bird := &Bird{}
		err := rows.Scan(&bird.ID, &bird.Name, &bird.Color)
		if err != nil {
			return nil, err
		}
		birds = append(birds, bird)
	}

	return birds, nil
}
