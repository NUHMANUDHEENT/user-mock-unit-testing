package models

type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email string `json:"username"`
	Password string `json:"password"`
}
