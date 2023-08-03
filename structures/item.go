package marketplace

type Item struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type UsersItem struct {
	Id     int
	UserId int
}
