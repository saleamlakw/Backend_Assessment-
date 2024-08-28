package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saleamlakw/LoanTracker/domain/entities"
)

type logController struct {
    usecase entities.LogUsecase
}

func NewLogController(usecase entities.LogUsecase) *logController {
    return &logController{
        usecase: usecase,
    }
}

func (lc *logController) GetAllLogs(ctx *gin.Context) {
    logs, err := lc.usecase.GetAllLogs(ctx)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, logs)
}
