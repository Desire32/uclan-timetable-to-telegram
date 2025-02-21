package timetable

import (
	"context"
	th "github.com/mymmrac/telego/telegohandler"
	"go.mongodb.org/mongo-driver/mongo"
)

type MicrosoftAuthService interface {
	MicrosoftLogin(ctx context.Context) error
}

type TimetableService interface {
	TimetableAuth(ctx context.Context) error
	TimetableRetrieve(ctx context.Context) string
}

type MongoDbService interface {
	MongoConnect() (*mongo.Client, error)
	MongoSend(jsonData string) error
}

type TelegramService interface {
	TgConnection() error
	TgHandlers(botHandler *th.BotHandler)
}

type BadgesService interface {
	BadgeRetrieve(ctx context.Context) error
	ModulesRetrieve(ctx context.Context) (string, error)
}
