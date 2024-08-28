package repositories

import (
	"context"

	"github.com/saleamlakw/LoanTracker/domain/entities"
	"github.com/saleamlakw/LoanTracker/internal/tokenutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	UserCollection *mongo.Collection
	RefreshCollection *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *userRepository {
	return &userRepository{
		UserCollection : client.Database("book").Collection("users"),
		RefreshCollection : client.Database("book").Collection("refresh"),
	}
}
func (ur *userRepository) SignupUser(ctx context.Context,user *entities.User) error {
	_, err := ur.UserCollection.InsertOne(ctx, user)
	return err
}

func (ur *userRepository) AccountExists(ctx context.Context,email string)(int64,error){
	count,err:=ur.UserCollection.CountDocuments(ctx,bson.M{"email":email})
	return count,err
}

func (ur *userRepository)CountUsers(ctx context.Context)(int64,error){
	result,err:=ur.UserCollection.CountDocuments(ctx,bson.M{})
	if err!=nil{
		return -1,err
	}
	return result,nil
 }
 
 func (ur *userRepository)GetUserById(ctx context.Context, userID string) (*entities.User, error){
	var user entities.User

	id, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return nil, err
	}

	err = ur.UserCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, err
 }

 func (ur *userRepository)ActivateUser(ctx context.Context, userID string) error{
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"is_active": true}}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var ResultUser entities.User
	err = ur.UserCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&ResultUser)

	if err != nil {
		return err
	}

	return  nil

 }

 func (ur *userRepository)GetUserByEmail(ctx context.Context, email string) (*entities.User, error){
 	var user entities.User
 	err := ur.UserCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
 	return &user, err
 }

 func (ur *userRepository) CreateRefreshData(c context.Context, refreshData entities.RefreshData) error {
	
	_, err := ur.RefreshCollection.InsertOne(c, refreshData)
	if err != nil {
		return err
	}
	return nil
}
func (ur *userRepository) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	userClaims, err := tokenutil.ExtractUserClaimsFromToken(requestToken, secret)
	if err != nil {
		return "", err
	}
	return userClaims["id"].(string), nil
}

func (ur *userRepository) GetRefreshData(c context.Context, id string) (*entities.RefreshData, error) {
	var refreshData entities.RefreshData
	err := ur.RefreshCollection.FindOne(c, bson.M{"_id": id}).Decode(&refreshData)
	if err != nil {
		return nil, err
	}
	return &refreshData, err
}

func (ur *userRepository) DeleteRefreshData(c context.Context, id string) error {
	_, err := ur.RefreshCollection.DeleteOne(c, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) DeleteUser(c context.Context, userID string) error {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	res, err := ur.UserCollection.DeleteOne(c, filter)

	if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	if err != nil {
		return err
	}

	return nil
}

func  (ur *userRepository)  GetUsers(ctx context.Context) ([]*entities.User, error) {
	var users []*entities.User
	cursor, err := ur.UserCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var user entities.User
		if err = cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

