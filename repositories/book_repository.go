package repositories

import (
	"context"
	"time"

	"github.com/saleamlakw/LoanTracker/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type bookRepository struct {
	bookCollection *mongo.Collection
}

func NewBookRepository(client *mongo.Client) *bookRepository {
	return &bookRepository{
		bookCollection:    client.Database("book").Collection("books"),
	}
}

func (rr *bookRepository) ApplyBorrowRequest(ctx context.Context,borrowRequest entities.BorrowRequest) (*mongo.InsertOneResult, error) {
    borrowRequest.CreatedAt = time.Now()
    borrowRequest.UpdatedAt = time.Now()
    return rr.bookCollection.InsertOne(ctx, borrowRequest)
}

func (rr *bookRepository) GetBorrowRequestByID(ctx context.Context,id primitive.ObjectID) (entities.BorrowRequest, error) {
    var borrowRequest entities.BorrowRequest
    err := rr.bookCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&borrowRequest)
    return borrowRequest, err
}

func (rr *bookRepository) GetAllBorrowRequests(ctx context.Context,status string, order string) ([]entities.BorrowRequest, error) {
    var borrowRequests []entities.BorrowRequest
    filter := bson.M{}
    if status != "all" {
        filter["status"] = status
    }

    opts := options.Find()
    if order == "asc" {
        opts.SetSort(bson.D{{Key: "created_at", Value: 1}})
    } else {
        opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
    }

    cursor, err := rr.bookCollection.Find(ctx, filter, opts)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var borrowRequest entities.BorrowRequest
        cursor.Decode(&borrowRequest)
        borrowRequests = append(borrowRequests, borrowRequest)
    }

    return borrowRequests, nil
}

func (rr *bookRepository) UpdateBorrowRequestStatus(ctx context.Context,id primitive.ObjectID, status string) error {

    filter := bson.M{"_id": id}
    update := bson.M{"$set": bson.M{"status": status, "updated_at": time.Now()}}
    _, err := rr.bookCollection.UpdateOne(ctx, filter, update)
    return err
}

func (rr *bookRepository) DeleteBorrowRequest(ctx context.Context,id primitive.ObjectID) error {
    _, err := rr.bookCollection.DeleteOne(ctx, bson.M{"_id": id})
    return err
}

