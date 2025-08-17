package service

import (
	"context"
	schedulerService "notification-service/src/scheduler/service"
	"notification-service/src/sender"
	"notification-service/src/server/handler/model"
	"notification-service/src/server/repository"
)

type ServiceInterface interface {
	CreateNotification(ctx context.Context, notification *model.CreateNotificationRequest) (*model.NotificationResponse, error)
	UpdateNotification(ctx context.Context, notification *model.UpdateNotificationRequest) (*model.NotificationResponse, error)
	GetNotificationByID(ctx context.Context, notification *model.GetNotificationRequest) (*model.NotificationResponse, error)
	// DeleteNotification(ctx context.Context, notification *model.DeleteNotificationRequest) error
}

type Service struct {
	repo        repository.RepositoryInterface
	sender      sender.SenderServiceInterface
	scheduler   schedulerService.SchedulerServiceInterface
	receiveDone chan *model.MarkNotificationAsSentRequest
}

func NewService(repo repository.RepositoryInterface, sender sender.SenderServiceInterface, scheduler schedulerService.SchedulerServiceInterface) ServiceInterface {
	return &Service{
		repo:        repo,
		sender:      sender,
		scheduler:   scheduler,
		receiveDone: make(chan *model.MarkNotificationAsSentRequest, 100),
	}
}

func (s *Service) CreateNotification(ctx context.Context, notification *model.CreateNotificationRequest) (*model.NotificationResponse, error) {

	// Call the repository to create the notification
	resp, err := s.repo.CreateNotification(ctx, notification)
	if err != nil {
		return nil, err
	}

	if notification.NotificationSchedule == nil {
		// If the notification is scheduled for the past, send it immediately
		err = s.sender.SendNotification(ctx, resp)
		if err != nil {
			return nil, err
		}
		// s.sender.QueueNotification(resp)
	} else {
		err = s.scheduler.CreateNotificationSchedule(ctx, resp.Id, notification.NotificationSchedule)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

func (s *Service) UpdateNotification(ctx context.Context, notification *model.UpdateNotificationRequest) (*model.NotificationResponse, error) {
	// Call the repository to update the notification
	resp, err := s.repo.UpdateNotification(ctx, notification)
	if err != nil {
		return nil, err
	}
	if notification.NotificationSchedule != nil {
		err = s.scheduler.UpdateNotificationSchedule(ctx, resp.Id, notification.NotificationSchedule)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (s *Service) GetNotificationByID(ctx context.Context, notification *model.GetNotificationRequest) (*model.NotificationResponse, error) {
	// Call the repository to get the notification by ID
	resp, err := s.repo.GetNotificationByID(ctx, notification.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// func (s *Service) DeleteNotification(ctx context.Context, notification *model.DeleteNotificationRequest) error {
// 	// Call the repository to delete the notification
// 	err := s.repo.DeleteNotification(ctx, notification.Id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
