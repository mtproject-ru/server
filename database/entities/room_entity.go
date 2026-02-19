package entities

import (
	"database/sql"
	"time"
)

type RoomEntity struct {
	Id         uint       `db:"id"`
	Visibility Visibility `db:"visibility"`
	// Массив пользователей, которые могут видеть эту комнату (исключая самих участников)
	VisibleFor []uint `db:"visible_for"`
	// VARCHAR(64)
	RoomTitle string `db:"room_title"`
	// VARCHAR(512)
	RoomDescription sql.NullString `db:"room_description"`
	// Даты, на которых основана команата
	RoomDates []time.Time `db:"room_dates"`
	// Айди пользователя, создавшего комнату
	CreatedBy uint      `db:"created_by"`
	CreatedAt time.Time `db:"created_at"`
}

// RoomToUser реляция вида many-to-many
type RoomToUser struct {
	RoomId uint `db:"room_id"`
	UserId uint `db:"user_id"`
}
