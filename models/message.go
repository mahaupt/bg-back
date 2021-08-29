package models

type Message struct {
	Name    string `form:"name" binding:"required"`
	Email   string `form:"email" binding:"required,email"`
	Message string `form:"message" binding:"required"`
}
