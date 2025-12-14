package models

import "time"

// Ensure Model Has Swagger Tags

type User struct {
	ID        int       `json:"id" example:"1"`
	Name      string    `json:"name" example:"John Doe"`
	Email     string    `json:"email" example:"john@mail.com"`
	Age       int       `json:"age" example:"25"`
	CreatedAt time.Time `json:"created_at" example:"2025-12-14T10:12:30Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2025-12-14T10:12:30Z"`
}
