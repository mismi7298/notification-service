package repository

import (
	"context"
	"database/sql"
	"notification-service/src/server/handler/model"
)

type SenderRepositoryInterface interface {
	MarkNotificationAsSent(ctx context.Context, notification *model.MarkNotificationAsSentRequest) error
}

type SenderRepository struct {
	db *sql.DB
}

func NewSenderRepository(db *sql.DB) SenderRepositoryInterface {
	return &SenderRepository{
		db: db,
	}
}

func (r *SenderRepository) MarkNotificationAsSent(ctx context.Context, notification *model.MarkNotificationAsSentRequest) error {
	_, err := r.db.ExecContext(ctx, `
		Insert into delivery_ledger (notification_id, channel)
		VALUES (?, ?)
	`, notification.NotificationId, notification.Channel)
	return err
}
