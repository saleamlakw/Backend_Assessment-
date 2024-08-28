package entities

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Event     string             `bson:"event" json:"event"`
    Details   string             `bson:"details" json:"details"`
    CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type LogRepository interface {
    GetAllLogs(ctx context.Context) ([]Log, error)
    LogEvent(ctx context.Context,event, details string) error 
}

type LogUsecase interface {
    GetAllLogs(ctx context.Context) ([]Log, error)
}