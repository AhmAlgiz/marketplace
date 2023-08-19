package structures

type Item struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	UserId      int    `json:"user_id"`
}
