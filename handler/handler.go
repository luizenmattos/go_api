package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizenmattos/go_api/config"
	"github.com/luizenmattos/go_api/schemas"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("bind error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("validation erro: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}

func GetOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}

func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Opening",
	})
}

func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Opening",
	})
}

func ListOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}
