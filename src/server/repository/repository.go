package repository

import (
	"context"
	"database/sql"
	"strconv"

	handlerModel "notification-service/src/server/handler/model"
)

// repository stores all incoming notification in a DB

type RepositoryInterface interface {
	CreateNotification(ctx context.Context, notification *handlerModel.CreateNotificationRequest) (*handlerModel.NotificationResponse, error)
	UpdateNotification(ctx context.Context, notification *handlerModel.UpdateNotificationRequest) (*handlerModel.NotificationResponse, error)
	GetNotificationByID(ctx context.Context, id string) (*handlerModel.NotificationResponse, error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) RepositoryInterface {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateNotification(ctx context.Context, notification *handlerModel.CreateNotificationRequest) (*handlerModel.NotificationResponse, error) {
	result, err := r.db.ExecContext(ctx, `
		INSERT INTO notifications (id, status, message,  sender, receiver, type, channels)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, notification.Id, notification.Status, notification.Message, notification.Sender, notification.Receiver, notification.Type, notification.Channels)
	if err != nil {
		return nil, err
	}

	// Get the ID of the newly created notification
	notificationID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Return the created notification
	return &handlerModel.NotificationResponse{
		Id:       strconv.FormatInt(notificationID, 10),
		Status:   notification.Status,
		Message:  notification.Message,
		Sender:   notification.Sender,
		Receiver: notification.Receiver,
		Type:     notification.Type,
		Channels: notification.Channels,
	}, nil
}
func (r *Repository) UpdateNotification(ctx context.Context, notification *handlerModel.UpdateNotificationRequest) (*handlerModel.NotificationResponse, error) {
	_, err := r.db.ExecContext(ctx, `
		UPDATE notifications
		SET status = ?, message = ?, sender = ?, receiver = ?, type = ?, channels = ?
		WHERE id = ?
	`, notification.Status, notification.Message, notification.Sender, notification.Receiver, notification.Type, notification.Channels, notification.Id)
	if err != nil {
		return nil, err
	}

	// Return the updated notification
	return &handlerModel.NotificationResponse{
		Id:       notification.Id,
		Status:   notification.Status,
		Message:  notification.Message,
		Sender:   notification.Sender,
		Receiver: notification.Receiver,
		Type:     notification.Type,
		Channels: notification.Channels,
	}, nil
}

func (r *Repository) GetNotificationByID(ctx context.Context, id string) (*handlerModel.NotificationResponse, error) {
	var notification handlerModel.NotificationResponse
	err := r.db.QueryRowContext(ctx, `
		SELECT id, status, message, created_at, updated_at, sent_at, sender, receiver, type, channels
		FROM notifications
		WHERE id = ?
	`, id).Scan(&notification.Id, &notification.Status, &notification.Message, &notification.CreatedAt, &notification.UpdatedAt, &notification.SentAt, &notification.Sender, &notification.Receiver, &notification.Type, &notification.Channels)
	if err != nil {
		return nil, err
	}
	return &notification, nil
}
