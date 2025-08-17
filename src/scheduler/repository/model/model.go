package model

import "time"

const (
	NotificationRuleOneTime = "one_time"
	NotificationRuleDaily   = "daily"
	NotificatioonRule
)

type NotificationSchedule struct {
	Repeated       bool      `json:"repeated"`
	CronExpression string    `json:"cron_expression"`
	NextRunAt      time.Time `json:"next_run_at"`
}

// notification rule table design
type NotificationRuleTable struct {
	ID             string    `db:"id"`
	Name           string    `db:"name"`
	Description    string    `db:"description"`
	CronExpression string    `db:"cron_expression"` // if null then schedule for one time
	NotificationId string    `db:"notification_id"` // foreign key of notification table
	Repeated       bool      `db:"repeated"`
	NextRunAt      time.Time `db:"next_run_at"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	DeletedAt      time.Time `db:"deleted_at"`
}
