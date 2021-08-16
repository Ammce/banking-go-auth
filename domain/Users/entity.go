package users

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	email    string
	password string
	role     string
}
