package email

import "fmt"

// Sends notification via emails to the user

type EmailClient struct {
	apiKey string
}

type EmailClientInterface interface {
	SendNotification(userId string, message string) error
}

func NewEmailClient(apiKey string) EmailClientInterface {
	return &EmailClient{apiKey: apiKey}
}

func (c *EmailClient) SendNotification(userId string, message string) error {
	fmt.Println("Sending Email notification to user:", userId, "with message:", message)
	return nil
}
