package repository

import (
	"context"

	"firebase.google.com/go/v4/messaging"
)

type NotificationRepository interface {
	SendNotification(ctx context.Context, message *messaging.Message) error
}

type notificationRepository struct {
	FCMClient *messaging.Client
}

func NewNotificationRepository(client *messaging.Client) NotificationRepository {
	return &notificationRepository{FCMClient: client}
}

func (r *notificationRepository) SendNotification(ctx context.Context, message *messaging.Message) error {
	_, err := r.FCMClient.Send(ctx, message)
	return err
}
