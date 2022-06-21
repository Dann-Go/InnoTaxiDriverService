package handler

import (
	_ "github.com/Dann-Go/InnoTaxiDriverService/docs"
	"github.com/Dann-Go/InnoTaxiDriverService/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	driverService        *service.DriverService
	authorizationService *service.AuthorizationService
}

func NewHandler(driverService *service.DriverService, authorizationService *service.AuthorizationService) *Handler {
	return &Handler{
		driverService:        driverService,
		authorizationService: authorizationService,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	return router
}
