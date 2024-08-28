package entities

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BorrowRequest struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	BookID    primitive.ObjectID `bson:"book_id,omitempty" json:"book_id,omitempty" binding:"required"`
	UserID    primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty" binding:"required"`
	Status    string             `bson:"status" json:"status" binding:"required,oneof='pending approved rejected'"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type BookRepositary interface {
	UpdateBorrowRequestStatus(ctx context.Context,id primitive.ObjectID, status string) error
	GetAllBorrowRequests(ctx context.Context,status string, order string) ([]BorrowRequest, error)
	GetBorrowRequestByID(ctx context.Context,id primitive.ObjectID) (BorrowRequest, error)
	ApplyBorrowRequest(ctx context.Context,borrowRequest BorrowRequest) (*mongo.InsertOneResult, error)
	DeleteBorrowRequest(ctx context.Context,id primitive.ObjectID) error
}

type BookUsecase interface {
	DeleteBorrowRequest(ctx context.Context,id primitive.ObjectID) error
	UpdateBorrowRequestStatus(ctx context.Context,id primitive.ObjectID, status string) error
	GetAllBorrowRequests(ctx context.Context,status, order string) ([]BorrowRequest, error)
	GetBorrowRequestByID(ctx context.Context,id primitive.ObjectID) (BorrowRequest, error)
	ApplyBorrowRequest(ctx context.Context,bookID, userID primitive.ObjectID) error
}