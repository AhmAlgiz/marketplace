package structures

import "errors"

type Item struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price"`
	UserId      int    `json:"user_id" db:"user_id"`
}

type UpdateItem struct {
	Id          int     `json:"id" binding:"required" db:"id"`
	Title       *string `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
	Price       *int    `json:"price"`
}

func (u *UpdateItem) Validate() error {
	if u.Title == nil && u.Description == nil && u.Price == nil {
		return errors.New("empty update parameters")
	}
	return nil
}
