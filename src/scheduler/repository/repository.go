package repository

import (
	"context"
	"database/sql"
	"notification-service/src/scheduler/repository/model"
)

type schedulerRepository struct {
	db *sql.DB
}

type SchedulerRepositoryInterface interface {
	GetRecurringNotificationRules(ctx context.Context) ([]model.NotificationRuleTable, error)
	CreateNotificationSchedule(ctx context.Context, notificationId string, schedule *model.NotificationSchedule) error
	UpdateNotificationSchedule(ctx context.Context, notificationId string, schedule *model.NotificationSchedule) error
	DeleteNotificationSchedule(ctx context.Context, id string) error
	GetNotificationSchedule(ctx context.Context, id string) (*model.NotificationSchedule, error)
}

func NewSchedulerRepository(db *sql.DB) SchedulerRepositoryInterface {
	return &schedulerRepository{
		db: db,
	}
}

func (r *schedulerRepository) GetRecurringNotificationRules(ctx context.Context) ([]model.NotificationRuleTable, error) {
	return nil, nil
}
func (r *schedulerRepository) CreateNotificationSchedule(ctx context.Context, notificationId string, schedule *model.NotificationSchedule) error {

	query := "INSERT INTO notification_schedules (notification_id, repeated, cron_expression, next_run_at) VALUES (?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, notificationId, schedule.Repeated, schedule.CronExpression, schedule.NextRunAt)
	return err
}

func (r *schedulerRepository) UpdateNotificationSchedule(ctx context.Context, notificationId string, schedule *model.NotificationSchedule) error {
	query := "UPDATE notification_schedules SET repeated = ?, cron_expression = ?, next_run_at = ? WHERE notification_id = ?"
	_, err := r.db.ExecContext(ctx, query, schedule.Repeated, schedule.CronExpression, schedule.NextRunAt, notificationId)
	return err
}

func (r *schedulerRepository) DeleteNotificationSchedule(ctx context.Context, id string) error {
	query := "DELETE FROM notification_schedules WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *schedulerRepository) GetNotificationSchedule(ctx context.Context, id string) (*model.NotificationSchedule, error) {
	query := "SELECT repeated, cron_expression, next_run_at FROM notification_schedules WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, id)

	var schedule model.NotificationSchedule
	if err := row.Scan(&schedule.Repeated, &schedule.CronExpression, &schedule.NextRunAt); err != nil {
		return nil, err
	}
	return &schedule, nil
}
