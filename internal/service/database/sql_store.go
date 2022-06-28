package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/takclark/schedulator/api"
)

type SqlStore struct {
	db *sqlx.DB
}

func NewSqlStore(file string) (*SqlStore, error) {
	db, err := sqlx.Open("sqlite3", file)
	if err != nil {
		return nil, fmt.Errorf("opening db: %w", err)
	}

	return &SqlStore{db: db}, nil
}

func (s *SqlStore) Rules() ([]api.Rule, error) {
	data := []api.Rule{}

	if err := s.db.Select(&data, "SELECT * FROM rules;"); err != nil {
		return []api.Rule{}, fmt.Errorf("loading rules: %w", err)
	}

	return data, nil
}

func (s *SqlStore) Rule(id int64) (api.Rule, error) {
	data := api.Rule{}

	if err := s.db.Get(&data, "SELECT * FROM rules WHERE id = ?", id); err != nil {
		return api.Rule{}, fmt.Errorf("loading rule: %w", err)
	}

	return data, nil
}

func (s *SqlStore) CreateRule(data api.CreateRule) (api.Rule, error) {
	q := `
	INSERT INTO rules (
		name,
		expression,
		rule_type,
		data
	)
	VALUES (
		:name,
		:expression,
		:rule_type,
		:data
	)
	`

	res, err := s.db.NamedExec(q, data)
	if err != nil {
		return api.Rule{}, fmt.Errorf("inserting rule: %w", err)
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return api.Rule{}, fmt.Errorf("finding last inserted ID: %w", err)
	}

	return s.Rule(insertedID)
}

func (s *SqlStore) UpdateRule(data api.UpdateRule) (api.Rule, error) {
	return api.Rule{}, fmt.Errorf("not implemented")
}
