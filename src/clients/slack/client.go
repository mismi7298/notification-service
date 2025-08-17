package slack

import "fmt"

// Sends notification via slack to the user

type SlackClient struct {
	apiKey string
}

type SlackClientInterface interface {
	SendNotification(userId string, message string) error
}

func NewSlackClient(apiKey string) SlackClientInterface {
	return &SlackClient{apiKey: apiKey}
}

func (c *SlackClient) SendNotification(userId string, message string) error {
	fmt.Println("Sending Slack notification to user:", userId, "with message:", message)
	return nil
}
