package route

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/LoanTracker/api/controller"
	"github.com/saleamlakw/LoanTracker/api/middleware"
	"github.com/saleamlakw/LoanTracker/repositories"
	"github.com/saleamlakw/LoanTracker/usecases"
	"go.mongodb.org/mongo-driver/mongo"
)

func Route(router *gin.Engine, client *mongo.Client) {
	ur:=repositories.NewUserRepository(client)
	uu:=usecases.NewUserUserCase(ur)
	uc:=controller.NewUserController(uu)
	AccessTokenSecret:=os.Getenv("ACCESS_TOKEN_SECRET")


	private:=router.Group("/",middleware.JwtAuthMiddleware(AccessTokenSecret))
	//user routes
	//public 
	router.POST("/users/register",uc.SignupUser)
	router.POST( "/users/verify-email/:token",uc.VerifyEmail)
	router.POST("/users/login",uc.Login)

	//private
	private.GET("/users",middleware.IsAdminMiddleware,uc.GetUsers)
	private.GET("/users/profile",uc.GetProfile)
	private.DELETE("/users/:userID",middleware.IsAdminMiddleware,uc.DeleteUser)


	

    bookRepo := repositories.NewBookRepository(client)
	LogRepo:=repositories.NewLogRepository(client)

    bookUsecase := usecases.NewBookUserCase(bookRepo,LogRepo)
	logUsecase:=usecases.NewLogUserCase(LogRepo)
	
    bookController := controller.NewbookController(bookUsecase)
	logController := controller.NewLogController( logUsecase)
	//book routes
	//public 
    // Book borrow routes
    private.POST("/books/borrow", bookController.ApplyBorrowRequest)
    private.GET("/books/borrow/:id", bookController.GetBorrowRequestByID)

    // Admin routes
    private.GET("/books/borrows",middleware.IsAdminMiddleware,bookController.GetAllBorrowRequests)
    private.PATCH("/books/borrows/:id/status",middleware.IsAdminMiddleware, bookController.UpdateBorrowRequestStatus)
    private.DELETE("/books/borrows/:id", middleware.IsAdminMiddleware,bookController.DeleteBorrowRequest)

	//log routes
	private.GET("/logs",middleware.IsAdminMiddleware,logController.GetAllLogs)
}