package database

import (
	"github.com/go-toschool/platon/database/postgres"
	"github.com/go-toschool/platon/entities"
	"github.com/jmoiron/sqlx"
)

// Store define the behavior for a Store.
type Store interface {
	Get(*entities.TalksQuery) (*entities.Talk, error)
	Select() ([]*entities.Talk, error)
	Create(*entities.Talk) error
	Update(*entities.Talk) error
	Delete(*entities.Talk) error
}

// NewPostgres return a new talks service.
func NewPostgres(dsn string) (Store, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &postgres.TalksService{
		Store: db,
	}, nil
}
