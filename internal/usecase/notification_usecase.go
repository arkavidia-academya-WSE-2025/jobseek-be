package usecase

import (
	"context"
	"fp-academya-be/internal/entity"
	"fp-academya-be/internal/repository"

	"firebase.google.com/go/v4/messaging"
)

type NotificationUsecase interface {
	SendNotification(ctx context.Context, notification *entity.Notification) error
}

type notificationUsecase struct {
	Repo repository.NotificationRepository
}

func NewNotificationUsecase(repo repository.NotificationRepository) NotificationUsecase {
	return &notificationUsecase{Repo: repo}
}

func (u *notificationUsecase) SendNotification(ctx context.Context, notification *entity.Notification) error {
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: notification.Title,
			Body:  notification.Body,
		},
		Token: notification.Token, // Device Token
	}
	return u.Repo.SendNotification(ctx, message)
}
