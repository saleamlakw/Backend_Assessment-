package entities

import (
	"context"
	"time"

	"github.com/saleamlakw/LoanTracker/domain/forms"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	FirstName string             `json:"first_name,omitempty" bson:"first_name" binding:"required,min=3,max=30"`
	LastName  string             `json:"last_name,omitempty" bson:"last_name" binding:"required,min=3,max=30"`
	Email     string             `json:"email,omitempty" bson:"email" binding:"required,email"`
	Password  string             `json:"password,omitempty" bson:"password" binding:"required,min=8,max=30"`
	Role      string             `json:"role,omitempty" bson:"role"`
	IsActive  bool               `json:"is_active,omitempty" bson:"is_active"`
}

type RefreshData struct{
	Id primitive.ObjectID `json:"_id"`
	UserId string `json:"userId"`
	RefreshToken string `json:" refreshToken"`
	Expire_date  time.Time `json:"expire_date"`
}

type UserUserCase interface {
	SignupUser(ctx context.Context,signuprequest *forms.SignupForm) error
	VerifyEmail(ctx context.Context, Verificationtoken string)	error
	Login(ctx context.Context, loginRequest *forms.LoginForm) (*forms.LoginResponseForm, error)
	RefreshToken(ctx context.Context, request *forms.RefreshTokenRequestForm,refreshDataIDStr string) (*forms.RefreshTokenResponseForm, error)
	DeleteUser(c context.Context, userID string) error
	GetProfile(ctx context.Context, userID string) (*User, error)
	GetUsers(ctx context.Context) ([]*User, error)
}
type UserRepository interface {
	SignupUser(ctx context.Context,user *User) error
	AccountExists(ctx context.Context,email string)(int64,error)
	CountUsers(ctx context.Context)(int64,error)
	ActivateUser(ctx context.Context, userID string) error
	GetUserById(ctx context.Context, userID string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CreateRefreshData(c context.Context, refreshData RefreshData) error
	ExtractIDFromToken(requestToken string, secret string) (string, error) 
	GetRefreshData(c context.Context, id string) (*RefreshData, error)
	DeleteRefreshData(c context.Context, id string) error
	DeleteUser(c context.Context, userID string) error
	GetUsers(ctx context.Context) ([]*User, error)
}