package entities

import (
	"time"

	"github.com/go-toschool/platon/talks"
)

// Talk struct
type Talk struct {
	ID          string    `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Repository  string    `json:"repository" db:"repository"`
	Date        time.Time `json:"attendance_date" db:"attendance_date"`
	Tags        string    `json:"tags" db:"tags"`
	UserID      string    `json:"user_id" db:"user_id"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
}

// TalksQuery represents values to query database.
type TalksQuery struct {
	ID     string
	UserID string
}

// Talks represents basic behavior for plato service.
type Talks interface {
	TalksGetter

	Select() ([]*Talk, error)
	Create(*Talk) error
	Update(*Talk) error
	Delete(*Talk) error
}

// TalksGetter represents getter actions.
type TalksGetter interface {
	GetByID(id string) (*Talk, error)
	GetByUserID(userID string) (*Talk, error)
	GetByIDAndUserID(id, userID string) (*Talk, error)
}

// ToProto helper function to transform to proto message.
func (t *Talk) ToProto() *talks.Talk {
	return &talks.Talk{
		Id:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Repository:  t.Repository,
		Date:        t.Date.Unix(),
		Tags:        t.Tags,
		UserId:      t.UserID,
		CreatedAt:   t.CreatedAt.Unix(),
		UpdatedAt:   t.UpdatedAt.Unix(),
	}
}

// FromProto helper function to transform to plato struct.
func (t *Talk) FromProto(tt *talks.Talk) *Talk {
	t.ID = tt.Id
	t.Title = tt.Title
	t.Description = tt.Description
	t.Repository = tt.Repository
	t.Date = time.Unix(tt.Date, 0)
	t.Tags = tt.Tags
	t.UserID = tt.UserId
	t.CreatedAt = time.Unix(tt.CreatedAt, 0)
	t.UpdatedAt = time.Unix(tt.UpdatedAt, 0)

	return t
}
