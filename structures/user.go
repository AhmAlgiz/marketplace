package structures

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
	Pass string `json:"pass" binding:"required"`
}
