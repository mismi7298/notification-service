package model

import (
	"notification-service/src/scheduler/repository/model"
)

type CreateNotificationRequest struct {
	Id                   string                      `json:"id"`
	Status               string                      `json:"status"`
	Message              string                      `json:"message"`
	NotificationSchedule *model.NotificationSchedule `json:"notification_schedule"`
	Sender               string                      `json:"sender"`
	Receiver             string                      `json:"receiver"`
	Type                 string                      `json:"type"`
	Channels             []string                    `json:"channels"` // eg: email, slack, inApp
}

type UpdateNotificationRequest struct {
	Id                   string                      `json:"id"`
	Status               string                      `json:"status"`
	Message              string                      `json:"message"`
	Sender               string                      `json:"sender"`
	Receiver             string                      `json:"receiver"`
	Type                 string                      `json:"type"`
	Channels             []string                    `json:"channels"` // eg: email, slack, inApp
	NotificationSchedule *model.NotificationSchedule `json:"notification_schedule"`
}

type GetNotificationRequest struct {
	Id string `json:"id"`
}

type DeleteNotificationRequest struct {
	Id string `json:"id"`
}

type MarkNotificationAsSentRequest struct {
	NotificationId string `json:"notification_id"`
	Channel        string `json:"channel"`
}
