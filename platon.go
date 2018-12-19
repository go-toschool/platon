package platon

import "time"

// Talk struct
type Talk struct {
	ID           string    `json:"id" db:"id"`
	Title        string    `json:"title" db:"title"`
	Description  string    `json:"description" db:"description"`
	Repository   string    `json:"repository" db:"repository"`
	Date         time.Time `json:"date" db:"date"`
	Tags         string    `json:"tags" db:"tags"`
	User_user_id string    `json:"user_id" db:"user_id"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
}

// Talks interface
type Talks interface {
	Get(string) (*Talk, error)
	Select() ([]*Talk, error)
	Create(*Talk) error
	Update(*Talk) error
	Delete(*Talk) error
}
