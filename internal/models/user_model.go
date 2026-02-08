package models

type User struct {
	BaseModel

	Email    string `json:"email"`
	Password string `json:"-"`
	Name     string `json:"name"`
}
