package structures

type User struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"username" binding:"required" db:"username"`
	Pass string `json:"pass" binding:"required" db:"pass_hash"`
}

type UpdateUser struct {
	Name string `json:"username" binding:"required" db:"username"`
}
