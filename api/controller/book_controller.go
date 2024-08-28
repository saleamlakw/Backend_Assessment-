package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/LoanTracker/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type bookController struct {
	bookUsecase entities.BookUsecase
}

func NewbookController(bookusecase entities.BookUsecase) *bookController {
	return &bookController{
		bookUsecase: bookusecase,
	}
}

func (c *bookController) ApplyBorrowRequest(ctx *gin.Context) {
    var req struct {
        BookID string `json:"book_id"`
        UserID string `json:"user_id"`
    }
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    bookID, err := primitive.ObjectIDFromHex(req.BookID)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }
    userID, err := primitive.ObjectIDFromHex(req.UserID)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    err = c.bookUsecase.ApplyBorrowRequest(ctx,bookID, userID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"status": "request submitted"})
}

func (c *bookController) GetBorrowRequestByID(ctx *gin.Context) {
    id := ctx.Param("id")
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
        return
    }

    borrowRequest, err := c.bookUsecase.GetBorrowRequestByID(ctx,objectID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, borrowRequest)
}

func (c *bookController) GetAllBorrowRequests(ctx *gin.Context) {
    status := ctx.DefaultQuery("status", "all")
    order := ctx.DefaultQuery("order", "asc")

    borrowRequests, err := c.bookUsecase.GetAllBorrowRequests(ctx,status, order)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, borrowRequests)
}

func (c *bookController) UpdateBorrowRequestStatus(ctx *gin.Context) {
    id := ctx.Param("id")
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
        return
    }

    var req struct {
        Status string `json:"status"`
    }
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err = c.bookUsecase.UpdateBorrowRequestStatus(ctx,objectID, req.Status)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (c *bookController) DeleteBorrowRequest(ctx *gin.Context) {
    id := ctx.Param("id")
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
        return
    }

    err = c.bookUsecase.DeleteBorrowRequest(ctx,objectID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

