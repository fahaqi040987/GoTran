package user

import (
	"time"
)

// Struct di go membuat declaration berdasarkan database yang sudah di buat
type User struct {
	Id             int
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
