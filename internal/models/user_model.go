package models

type User struct {
	BaseModel

	Email    string `gorm:"email;unique" json:"email"`
	Password string `gorm:"password" json:"-"`
	Name     string `gorm:"name" json:"name"`
}
