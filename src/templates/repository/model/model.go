package model

import "time"

// template table definition
type TemplateTable struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Content   string    `db:"content"`
	Channel   string    `db:"channel"`
	Type      string    `db:"type"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}
