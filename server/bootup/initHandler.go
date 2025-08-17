package bootup

import (
	"database/sql"
	"notification-service/src/clients"
	schedulerRepository "notification-service/src/scheduler/repository"
	schedulerService "notification-service/src/scheduler/service"
	"notification-service/src/sender"
	senderRepository "notification-service/src/sender/repository"
	notificationHandler "notification-service/src/server/handler"
	notificationRepository "notification-service/src/server/repository"
	notificationService "notification-service/src/server/service"
	templateHandler "notification-service/src/templates/handler"
	templateRepository "notification-service/src/templates/repository"
	templateService "notification-service/src/templates/service"
)

type Handlers struct {
	NotificationHandler *notificationHandler.NotificationHandler
	TemplateHandler     *templateHandler.TemplateHandler
}

func InitHandlers(db *sql.DB) *Handlers {
	//
	templateRepository := templateRepository.NewTemplateRepository(db)
	notificationRepository := notificationRepository.NewRepository(db)
	schedulerRepository := schedulerRepository.NewSchedulerRepository(db)
	senderRepository := senderRepository.NewSenderRepository(db)

	clientManager := clients.NewClientManager()

	templateService := templateService.NewTemplateService(templateRepository)
	senderService := sender.NewSenderService(templateService, clientManager, senderRepository)
	schedulerService := schedulerService.NewSchedulerService(schedulerRepository, senderService, notificationRepository)
	notificationService := notificationService.NewService(notificationRepository, senderService, schedulerService)

	return &Handlers{
		NotificationHandler: notificationHandler.NewNotificationHandler(notificationService),
		TemplateHandler:     templateHandler.NewTemplateHandler(templateService),
	}
}
