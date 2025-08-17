package clients

import (
	"fmt"
	"notification-service/src/clients/appNotifier"
	"notification-service/src/clients/email"
	"notification-service/src/clients/slack"
)

// sends notification client to requester according to requested client type

type ClientType string

const (
	EmailClientType ClientType = "email"
	SlackClientType ClientType = "slack"
	AppClientType   ClientType = "app"
)

type ClientManager struct {
	clients map[ClientType]Client
}

type Client interface {
	SendNotification(userId string, message string) error
}

type ClientManagerInterface interface {
	GetClient(name string) (Client, error)
}

func NewClientManager() ClientManagerInterface {

	return &ClientManager{clients: map[ClientType]Client{
		EmailClientType: email.NewEmailClient(""),
		SlackClientType: slack.NewSlackClient(""),
		AppClientType:   appNotifier.NewAppClient(""),
	}}
}

func (c *ClientManager) GetClient(name string) (Client, error) {
	clientType := ClientType(name)

	client, ok := c.clients[clientType]
	if !ok {
		return nil, fmt.Errorf("client %s not found", name)
	}
	return client, nil
}
