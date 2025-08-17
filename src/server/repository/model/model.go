package model

import "time"

const (
	NotificationStatePending   = "pending"
	NotificationStateSent      = "sent"
	NotificationStateFailed    = "failed"
	NotificationStateScheduled = "scheduled"
)

// notification table schema

type NotificationTable struct {
	ID        string    `db:"id"`
	Sender    string    `db:"sender"`
	Receiver  string    `db:"receiver"`
	Message   string    `db:"message"`
	Status    string    `db:"status"`
	Type      string    `db:"type"`
	Channels  []string  `db:"channels"` // eg: email, slack, inApp
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}
