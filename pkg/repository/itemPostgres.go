package repository

import (
	"fmt"
	"strings"

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

func (r *ItemPostgres) GetItemById(id int) ([]structures.Item, error) {
	query := fmt.Sprintf(
		`SELECT * FROM %s WHERE id=$1`, itemsTable)

	var sl []structures.Item

	err := r.db.Select(&sl, query, id)

	return sl, err
}

func (r *ItemPostgres) GetItemByTitle(title string) ([]structures.Item, error) {
	query := fmt.Sprintf(
		`SELECT * FROM %s WHERE title=$1`, itemsTable)

	var sl []structures.Item

	err := r.db.Select(&sl, query, title)

	return sl, err
}

func (r *ItemPostgres) GetItemByUsername(username string) ([]structures.Item, error) {
	query := fmt.Sprintf(
		`SELECT it.id, it.title, it.description, it.price, it.user_id FROM %s it INNER JOIN %s ut ON it.user_id=ut.id WHERE ut.username=$1`, itemsTable, usersTable)

	var sl []structures.Item

	err := r.db.Select(&sl, query, username)

	return sl, err
}

func (r *ItemPostgres) DeleteItem(id, userId int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s it WHERE it.id=$1 AND it.user_id=$2 `, itemsTable)

	_, err := r.db.Exec(query, id, userId)

	return err
}

func (r *ItemPostgres) GetAllItems() ([]structures.Item, error) {
	query := fmt.Sprintf(
		`SELECT * FROM %s`, itemsTable)

	var sl []structures.Item

	err := r.db.Select(&sl, query)

	return sl, err
}

func (r *ItemPostgres) UpdateItem(input structures.UpdateItem, userId int) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}
	args = append(args, userId)
	args = append(args, input.Id)

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s it SET %s WHERE it.user_id=$%d AND it.id=$%d", itemsTable, setQuery, argId, argId+1)

	_, err := r.db.Exec(query, args...)

	return err
}
