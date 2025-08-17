package model

import "time"

type NotificationResponse struct {
	Id                   string    `db:"id"`
	Status               string    `db:"status"`
	Message              string    `db:"message"`
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
	SentAt               time.Time `db:"sent_at"`
	Sender               string    `db:"sender"`
	Receiver             string    `db:"receiver"`
	Type                 string    `db:"type"`
	NotificationSchedule time.Time `db:"notification_schedule"` // Optional, for scheduled notifications
	Channels             []string  `db:"channels"`              // eg: email, slack, inApp
}
