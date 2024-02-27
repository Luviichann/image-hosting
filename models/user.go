package models

type User struct {
	Id       int
	Username string `form:"username"`
	Password string `form:"password"`
	Auth     int
}

func (User) TableName() string {
	return "user"
}
