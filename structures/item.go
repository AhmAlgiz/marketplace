package structures

type Item struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price"`
	UserId      int    `json:"user_id" db:"user_id"`
}
