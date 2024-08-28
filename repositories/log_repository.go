package repositories

import (
	"context"
	"time"

	"github.com/saleamlakw/LoanTracker/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogRepository struct {
    collection *mongo.Collection
}

func NewLogRepository(client *mongo.Client) *LogRepository {
    return &LogRepository{
        collection: client.Database("book").Collection("logs"),
    }
}

func (r *LogRepository) LogEvent(ctx context.Context,event, details string) error {
    log := entities.Log{
        Event:     event,
        Details:   details,
        CreatedAt: time.Now(),
    }
    _, err := r.collection.InsertOne(ctx, log)
    return err
}

func (r *LogRepository) GetAllLogs(ctx context.Context) ([]entities.Log, error) {
    var logs []entities.Log

    cursor, err := r.collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var log entities.Log
        cursor.Decode(&log)
        logs = append(logs, log)
    }

    return logs, nil
}
