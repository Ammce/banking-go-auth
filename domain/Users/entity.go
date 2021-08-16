package users

import (
	"github.com/gofrs/uuid"
)

type User struct {
	ID       uuid.UUID
	email    string
	password string
	role     string
}
