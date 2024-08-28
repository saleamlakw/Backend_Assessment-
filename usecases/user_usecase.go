package usecases

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	b64 "encoding/base64"

	"github.com/saleamlakw/LoanTracker/domain/entities"
	"github.com/saleamlakw/LoanTracker/domain/forms"
	"github.com/saleamlakw/LoanTracker/internal/emailutil"
	"github.com/saleamlakw/LoanTracker/internal/passwordutil"
	"github.com/saleamlakw/LoanTracker/internal/tokenutil"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userUserCase struct {
	UserRepository entities.UserRepository
}

func NewUserUserCase(UserRepository entities.UserRepository) *userUserCase {
	return &userUserCase{
		UserRepository: UserRepository,
	}
}

func (uu *userUserCase) SignupUser(ctx context.Context,signuprequest *forms.SignupForm) error {
	count,err:=uu.UserRepository.AccountExists(ctx,signuprequest.Email)
	if err != nil {
		return err
	}
	if count>0 {
		return errors.New("account already exists")
	}

	hashedPassword,err:=passwordutil.HashPassword(signuprequest.Password)
	if err != nil {
		return err
	}
	
	var user entities.User
	user.ID=primitive.NewObjectID()
	user.FirstName=signuprequest.FirstName
	user.LastName=signuprequest.LastName
	user.Email=signuprequest.Email
	user.Password=hashedPassword
	user.IsActive=false
	result,err:=uu.UserRepository.CountUsers(ctx)
	if err != nil {
		return err
	}
	if result==0{
		user.Role="admin"
	}else{
		user.Role="user"
	}
	uu.UserRepository.SignupUser(ctx, &user)
	VerificationTokenSecret:=os.Getenv("VERIFICATION_TOKEN_SECRET")

	VerificationTokenExpiryMin,_:=strconv.Atoi(os.Getenv("VERIFICATION_TOKEN_EXPIRY_MIN"))

	VerificationToken, err := tokenutil.CreateVerificationToken(&user, VerificationTokenSecret, VerificationTokenExpiryMin)
	if err != nil {
		return errors.New("failed to create verification token")
	}
	encodedToken := b64.URLEncoding.EncodeToString([]byte(VerificationToken))
	err = emailutil.SendVerificationEmail(user.Email, encodedToken, )

	if err != nil {
		return errors.New("failed to send verification email")
	}
	return nil

}

func (uu *userUserCase) VerifyEmail(ctx context.Context, verificationToken string) error {
	decodedToken, _ := b64.URLEncoding.DecodeString(verificationToken)
	VerificationTokenSecret:=os.Getenv("VERIFICATION_TOKEN_SECRET")
	valid, err := tokenutil.IsAuthorized(string(decodedToken), VerificationTokenSecret)
	if !valid || err != nil {
		return errors.New("invalid verification token")
	}



	claims, err := tokenutil.ExtractUserClaimsFromToken(string(decodedToken), VerificationTokenSecret)
	userID := claims["id"].(string)
	if err != nil {
		return err
	}
	user, err := uu.UserRepository.GetUserById(context.TODO(), userID)
	if err != nil {
		return err
	}
	if user.IsActive {
		return errors.New("user is already active")
	}

	err = uu.UserRepository.ActivateUser(context.TODO(), userID)
	if err != nil {
		return err
	}

	return nil
}

func (uu *userUserCase) Login (ctx context.Context, loginRequest *forms.LoginForm) (*forms.LoginResponseForm, error){
	user, err := uu.UserRepository.GetUserByEmail(ctx, loginRequest.Email)
	if err != nil {
		return nil, err
	}

	if !user.IsActive {
		return nil, errors.New("user is not active")
	}

	result, err := passwordutil.VerifyPassword(loginRequest.Password, user.Password)
	if !result || err != nil {
		return nil, errors.New("invalid credentials")
	}
		

	var refreshData entities.RefreshData
	refreshData.Id = primitive.NewObjectID()
	refreshData.UserId = user.ID.Hex()

	RefreshTokenExpiryHour,err:= strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRY_HOUR"))

	if err != nil {
		return nil, err
	}
	refreshData.Expire_date =time.Now().Add(time.Hour * time.Duration(RefreshTokenExpiryHour))
	AccessTokenSecret:=os.Getenv("ACCESS_TOKEN_SECRET")
	AccessTokenExpiryHour,err:= strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR"))

	if err != nil {	
		return nil, err
	}

	accessToken, err := tokenutil.CreateAccessToken(user, AccessTokenSecret, AccessTokenExpiryHour, refreshData.Id.Hex())
	if err != nil {
		return nil, err
	}
	RefreshTokenSecret:=os.Getenv("REFRESH_TOKEN_SECRET")

	refreshToken, err := tokenutil.CreateRefreshToken(user, RefreshTokenSecret, RefreshTokenExpiryHour, refreshData.Id.Hex())
	if err != nil {
		return nil, err
	}
	refreshData.RefreshToken = refreshToken
	err = uu.UserRepository.CreateRefreshData(ctx, refreshData)
	if err != nil {
		return nil, err
	}

	loginResponse := forms.LoginResponseForm{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &loginResponse, nil

}

func (uu *userUserCase) RefreshToken(ctx context.Context, request *forms.RefreshTokenRequestForm,refreshDataIDStr string) (*forms.RefreshTokenResponseForm, error) {
	RefreshTokenSecret:=os.Getenv("REFRESH_TOKEN_SECRET")
	if valid, err := tokenutil.IsAuthorized(string(request.RefreshToken), RefreshTokenSecret); !valid || err != nil {
		return nil, errors.New("invalid refresh token")
	}
	
	id, err := uu.UserRepository.ExtractIDFromToken(request.RefreshToken, RefreshTokenSecret)
	if err != nil {
		return nil, err
	}

	user, err := uu.UserRepository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}


	
	_, err = uu.UserRepository.GetRefreshData(ctx, refreshDataIDStr)
	if err != nil {
		return nil, err
	}
	AccessTokenSecret:=os.Getenv("ACCESS_TOKEN_SECRET")
	AccessTokenExpiryHour,err:=strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR"))
	if err!=nil{
		return nil, err
	}
	accessToken, err := tokenutil.CreateAccessToken(user, AccessTokenSecret, AccessTokenExpiryHour, refreshDataIDStr)
	if err!=nil{
		return nil, err
	}


	RefreshTokenExpiryHour,err:=strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRY_HOUR"))
	if err!=nil{
		return nil, err
	}
	refreshToken, err := tokenutil.CreateRefreshToken(user, RefreshTokenSecret, RefreshTokenExpiryHour, refreshDataIDStr)
	
	if err != nil {
		return nil, err
	}
	ID, err := primitive.ObjectIDFromHex(refreshDataIDStr)

	if err != nil {
		return nil, err
	}
	uu.UserRepository.DeleteRefreshData(ctx, refreshDataIDStr)
	var refreshData entities.RefreshData
	refreshData.Id = ID
	refreshData.UserId = user.ID.Hex()
	refreshData.Expire_date = refreshData.Expire_date
	refreshData.RefreshToken = refreshToken

	uu.UserRepository.CreateRefreshData(ctx, refreshData)

	refreshTokenResponse := forms.RefreshTokenResponseForm{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &refreshTokenResponse, nil
}
