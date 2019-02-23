package service

import (
	"github.com/go-toschool/platon/database"
	"github.com/go-toschool/platon/entities"
)

// NewTalksService expose a new citizens service.
func NewTalksService(store database.Store) *TalksService {
	return &TalksService{
		Store: store,
	}
}

// TalksService postgres implementation.
type TalksService struct {
	Store database.Store
}

// GetByID gets a record from db.
func (cs *TalksService) GetByID(id string) (*entities.Talk, error) {
	q := &entities.TalksQuery{
		ID: id,
	}
	return cs.Store.Get(q)
}

// GetByUserID gets a record from db.
func (cs *TalksService) GetByUserID(userID string) (*entities.Talk, error) {
	q := &entities.TalksQuery{
		UserID: userID,
	}

	return cs.Store.Get(q)
}

// GetByIDAndUserID gets a record from db.
func (cs *TalksService) GetByIDAndUserID(id, userID string) (*entities.Talk, error) {
	q := &entities.TalksQuery{
		ID:     id,
		UserID: userID,
	}

	return cs.Store.Get(q)
}

// Select returns a collectio of users from db.
func (cs *TalksService) Select() ([]*entities.Talk, error) {
	return cs.Store.Select()
}

// Create creates a new user.
func (cs *TalksService) Create(t *entities.Talk) error {
	return cs.Store.Create(t)
}

// Update updates the given user.
func (cs *TalksService) Update(t *entities.Talk) error {
	return cs.Store.Update(t)
}

// Delete logical delete.
func (cs *TalksService) Delete(t *entities.Talk) error {
	return cs.Store.Delete(t)
}
