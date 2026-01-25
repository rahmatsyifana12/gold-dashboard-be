package dtos

import "gold-dashboard-be/internal/models"

type GetUserByIDResponse struct {
	models.User
}

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Name     string `json:"name" validate:"omitempty,min=2"`
}

type UpdateUserParams struct {
	ID          uint   `param:"id" validate:"required"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

type UserIDParams struct {
	ID uint `param:"id" validate:"required"`
}
