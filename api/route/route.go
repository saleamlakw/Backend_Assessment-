package route

import (
	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/LoanTracker/repositories"
	"github.com/saleamlakw/LoanTracker/usecases"
	"github.com/saleamlakw/LoanTracker/api/controller"
	"go.mongodb.org/mongo-driver/mongo"
)

func Route(router *gin.Engine, client *mongo.Client) {
	ur:=repositories.NewUserRepository(client)
	uu:=usecases.NewUserUserCase(ur)
	uc:=controller.NewUserController(uu)

	router.POST("/users/register",uc.SignupUser)
	router.POST( "/users/verify-email/:token",uc.VerifyEmail)
	router.POST("/users/login",uc.Login)
}