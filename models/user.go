package perfume

type User struct {
	Id         int    `json:"id" db:"id"`
	Firstname  string `json:"firstname" binding:"required""`
	Secondname string `json:"secondname" binding:"required"`
	Username   string `json:"username" binding:"required"`
	Sex        string `json:"sex"`
	Password   string `json:"password" binding:"required"`
}
