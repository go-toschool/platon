package postgres

import (
	"database/sql"
	"errors"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/go-toschool/platon/entities"
	"github.com/jmoiron/sqlx"
)

// TalksService postgres implementation
type TalksService struct {
	Store *sqlx.DB
}

// Get gets a record from db
func (cs *TalksService) Get(q *entities.TalksQuery) (*entities.Talk, error) {
	query := squirrel.Select("*").From("talks").Where("deleted_at is null")

	if q.ID == "" && q.UserID == "" {
		return nil, errors.New("must provide a query")
	}

	if q.ID != "" {
		query = query.Where("id = ?", q.ID)
	}

	if q.UserID != "" {
		query = query.Where("user_id = ?", q.UserID)
	}

	sql, args, err := query.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	row := cs.Store.QueryRowx(sql, args...)

	t := &entities.Talk{}
	if err := row.StructScan(t); err != nil {
		return nil, err
	}

	return t, nil
}

// Select returns a collectio of users from db.
func (cs *TalksService) Select() ([]*entities.Talk, error) {
	query := squirrel.Select("*").From("talks").Where("deleted_at is null")

	sql, args, err := query.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := cs.Store.Queryx(sql, args...)
	if err != nil {
		return nil, err
	}

	tt := make([]*entities.Talk, 0)

	for rows.Next() {
		t := &entities.Talk{}
		if err := rows.StructScan(t); err != nil {
			return nil, err
		}
		tt = append(tt, t)
	}

	return tt, nil
}

// Create creates a new user.
func (cs *TalksService) Create(t *entities.Talk) error {
	sql, args, err := squirrel.
		Insert("talks").
		Columns("title", "description", "repository", "attendance_date", "tags", "user_id").
		Values(t.Title, t.Description, t.Repository, t.Date, t.Tags, t.UserID).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	row := cs.Store.QueryRowx(sql, args...)
	if err := row.StructScan(t); err != nil {
		return err
	}

	return nil
}

// Update updates the given user.
func (cs *TalksService) Update(t *entities.Talk) error {
	sql, args, err := squirrel.Update("talks").
		Set("title", t.Title).
		Set("description", t.Description).
		Set("repository", t.Repository).
		Set("date", t.Date).
		Set("tags", t.Tags).
		Where("id = ?", t.ID).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	row := cs.Store.QueryRowx(sql, args...)
	return row.StructScan(t)
}

// Delete logical delete.
func (cs *TalksService) Delete(t *entities.Talk) error {
	row := cs.Store.QueryRowx(
		"update talks set deleted_at = $1 where id = $2 returning *",
		time.Now(), t.ID,
	)

	if err := row.StructScan(t); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	return nil
}
