package entities

import (
	"database/sql"
)

type UserEntity struct {
	Id uint `db:"id"`
	// VARCHAR(64) UNIQUE
	Username string `db:"username"`
	// VARCHAR(64)
	FirstName string `db:"first_name"`
	// VARCHAR(64)
	LastName sql.NullString `db:"last_name"`
	// VARCHAR(64)
	Password string `db:"password"`
}

// UserToEvent Реляция вида many-to-many
type UserToEvent struct {
	UserId  uint `db:"user_id"`
	EventId uint `db:"event_id"`
}
