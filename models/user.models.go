package models

type User struct {
	Id       int    `form:"id"`
	Username string `form:"username"`
	Email    string `form:"email"`
	Password string `form:"password"`
}
