package appNotifier

import "fmt"

// Sends in-app notifications to the user

type AppClient struct {
	apiKey string
}

type AppClientInterface interface {
	SendNotification(userId string, message string) error
}

func NewAppClient(apiKey string) AppClientInterface {
	return &AppClient{apiKey: apiKey}
}

func (c *AppClient) SendNotification(userId string, message string) error {
	fmt.Println("Sending in-app notification to user:", userId, "with message:", message)
	return nil
}
