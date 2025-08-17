package service

import (
	"context"
	"notification-service/src/scheduler/repository"
	"notification-service/src/scheduler/repository/model"
	"notification-service/src/sender"
	serviceRepository "notification-service/src/server/repository"
)

type SchedulerServiceInterface interface {
	CreateNotificationSchedule(ctx context.Context, notificationId string, schedule *model.NotificationSchedule) error
	UpdateNotificationSchedule(ctx context.Context, notificationId string, schedule *model.NotificationSchedule) error
	DeleteNotificationSchedule(ctx context.Context, id string) error
	GetNotificationSchedule(ctx context.Context, id string) (*model.NotificationSchedule, error)
	GetRecurringNotificationRules(ctx context.Context) ([]model.NotificationRuleTable, error)
}

type schedulerService struct {
	repo              repository.SchedulerRepositoryInterface
	sender            sender.SenderServiceInterface
	serviceRepository serviceRepository.RepositoryInterface
}

func NewSchedulerService(repo repository.SchedulerRepositoryInterface, sender sender.SenderServiceInterface, serviceRepository serviceRepository.RepositoryInterface) SchedulerServiceInterface {
	return &schedulerService{
		repo:              repo,
		sender:            sender,
		serviceRepository: serviceRepository,
	}
}
func (s *schedulerService) ProcessNotificationSchedules(ctx context.Context) error {
	rules, err := s.repo.GetRecurringNotificationRules(ctx)
	if err != nil {
		return err
	}

	for _, rule := range rules {
		// Process each rule
		notification, err := s.serviceRepository.GetNotificationByID(ctx, rule.NotificationId)
		if err != nil {
			return err
		}
		s.sender.QueueNotification(notification)
		// Use the notification as needed
		// sender.QueueNotification(notification)
	}

	return nil
}

func (s *schedulerService) CreateNotificationSchedule(ctx context.Context, notificationId string, schedule *model.NotificationSchedule) error {
	return s.repo.CreateNotificationSchedule(ctx, notificationId, schedule)
}

func (s *schedulerService) UpdateNotificationSchedule(ctx context.Context, notificationId string, schedule *model.NotificationSchedule) error {
	return s.repo.UpdateNotificationSchedule(ctx, notificationId, schedule)
}

func (s *schedulerService) DeleteNotificationSchedule(ctx context.Context, id string) error {
	return s.repo.DeleteNotificationSchedule(ctx, id)
}

func (s *schedulerService) GetNotificationSchedule(ctx context.Context, id string) (*model.NotificationSchedule, error) {
	return s.repo.GetNotificationSchedule(ctx, id)
}
func (s *schedulerService) GetRecurringNotificationRules(ctx context.Context) ([]model.NotificationRuleTable, error) {
	return s.repo.GetRecurringNotificationRules(ctx)
}
