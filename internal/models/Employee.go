package models

type Employee struct {
	Name   string `json:"name" binding:"required"`
	Branch string `json:"branch" binding:"required"`
}
