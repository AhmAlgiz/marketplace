package repository

import (
	"fmt"

	"github.com/AhmAlgiz/marketplace/structures"
	"github.com/jmoiron/sqlx"
)

type ItemPostgres struct {
	db *sqlx.DB
}

func NewItemPostgres(db *sqlx.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

func (r *ItemPostgres) CreateItem(input structures.Item) (int, error) {

	//making SQL transaction
	tx, err := r.db.Begin()
	if err != nil {
		return 0, nil
	}

	var id int

	query := fmt.Sprintf(
		`INSERT INTO %s (title, description, price, user_id)
			VALUES ($1, $2, $3, $4) RETURNING id`, itemsTable)
	row := tx.QueryRow(
		query, input.Title, input.Description, input.Price, input.UserId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
