package entities

import (
	"database/sql"
	"time"
)

type EventEntity struct {
	Id         uint       `db:"id"`
	Visibility Visibility `db:"visibility"`
	// VARCHAR(128)
	EventTitle string `db:"event_title"`
	// VARCHAR(512)
	EventDescription sql.NullString `db:"event_description"`
	// Если событие начинается в какой то момент и имеет продолжительность
	// (в т.ч. и бесконечную), то указывается дата и время начала события
	EventStart sql.NullTime `db:"event_start"`

	// Если событие ограничено по времени до какой либо даты и какого либо
	// времени, то указывается дата и время окончания события
	//
	// При этом, если начало события указано, то конец события должен быть больше (позже) начала
	EventEnd sql.NullTime `db:"event_end"`

	// Дата и время создания события
	EventCreatedAt time.Time `db:"event_created_at"`
}
