package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/LoanTracker/domain/entities"
	"github.com/saleamlakw/LoanTracker/domain/forms"
)

type UserController struct {
	UserUsecase entities.UserUserCase
}

func NewUserController(userusecase entities.UserUserCase) *UserController {
	return &UserController{
		UserUsecase: userusecase,
	}
}
func (uc *UserController) GetUsers(c *gin.Context) {

	users, err := uc.UserUsecase.GetUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetProfile(c *gin.Context) {
	id, exists := c.Get("x-user-id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid x-user-id"})
		return
	}
	idstr, ok := id.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid x-user-id"})
		return
	}
	
	user, err := uc.UserUsecase.GetProfile(c.Request.Context(), idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	id:=c.Param("id")

	if err := uc.UserUsecase.DeleteUser(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
func (uc *UserController) SignupUser(c *gin.Context) {
	var signuprequest forms.SignupForm

	if err := c.ShouldBindJSON(&signuprequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.UserUsecase.SignupUser(c.Request.Context(), &signuprequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "email verification link sent"})
}
func (uc *UserController) VerifyEmail(c *gin.Context) {
	Verificationtoken := c.Param("token")
	err:=uc.UserUsecase.VerifyEmail(c.Request.Context(), Verificationtoken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
		return
	}	
	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}

func (uc *UserController) Login(c *gin.Context) {
	var request forms.LoginForm

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginResponse, err := uc.UserUsecase.Login(c.Request.Context(), &request)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, loginResponse)
}

func (uc *UserController) RefreshToken(c *gin.Context) {
	var request forms.RefreshTokenRequestForm

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	refreshDataID, exists := c.Get("x-user-refresh-data-id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "x-user-refresh-data-id not found"})
		c.Abort()
		return
	}
	refreshDataIDStr, ok := refreshDataID.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid x-user-refresh-data-id"})
		c.Abort()
		return
	}

	refreshTokenResponse, err := uc.UserUsecase.RefreshToken(c.Request.Context(), &request,refreshDataIDStr)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
	}
	c.JSON(http.StatusOK, refreshTokenResponse)
}


