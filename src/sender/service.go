package sender

import (
	"context"
	"fmt"
	"notification-service/src/clients"
	"notification-service/src/sender/repository"
	"notification-service/src/server/handler/model"
	templateModel "notification-service/src/templates/handler/model"
	templateService "notification-service/src/templates/service"
)

type SenderServiceInterface interface {
	SendNotification(ctx context.Context, notification *model.NotificationResponse) error
	QueueNotification(notification *model.NotificationResponse)
	Start()
	Stop()
}

type senderService struct {
	templateService templateService.TemplateServiceInterface
	clientManager   clients.ClientManagerInterface
	messageChan     chan *model.NotificationResponse
	repo            repository.SenderRepositoryInterface
}

func NewSenderService(templateService templateService.TemplateServiceInterface, clientManager clients.ClientManagerInterface, repo repository.SenderRepositoryInterface) SenderServiceInterface {
	return &senderService{
		templateService: templateService,
		clientManager:   clientManager,
		repo:            repo,
		messageChan:     make(chan *model.NotificationResponse, 100), // Buffered channel for async processing
	}
}

func (s *senderService) SendNotification(ctx context.Context, notification *model.NotificationResponse) (err error) {

	for _, channel := range notification.Channels {
		// get template for every type
		template, err := s.templateService.GetTemplateByID(ctx, &templateModel.GetTemplateRequest{
			Type:    notification.Type,
			Channel: channel,
		})
		if err != nil {
			return err
		}
		//Todo: template filling can be done with Name and other information
		// get client for the type of message
		client, err := s.clientManager.GetClient(channel)
		if err != nil {
			return err
		}
		// send message
		err = client.SendNotification(notification.Receiver, template.Content)
		if err != nil {
			return err
		}
		// mark notification as sent
		err = s.repo.MarkNotificationAsSent(ctx, &model.MarkNotificationAsSentRequest{
			NotificationId: notification.Id,
			Channel:        channel,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *senderService) Start() {
	go func() {
		for notification := range s.messageChan {
			err := s.SendNotification(context.Background(), notification)
			if err != nil {
				// Handle error (e.g., log it)
				fmt.Println("Error sending notification:", err, "Notification:", notification)
			}
		}
	}()
}

func (s *senderService) Stop() {
	close(s.messageChan)
}

func (s *senderService) QueueNotification(notification *model.NotificationResponse) {
	select {
	case s.messageChan <- notification:
	default:
		// Handle backpressure (e.g., log a warning, drop the notification, etc.)
	}
}
