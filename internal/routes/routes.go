package routes

import (
	"subscription-service/internal/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(subscriptionHandler *handlers.SubscriptionHandler) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		sub := v1.Group("/subscriptions")
		{
			sub.POST("", subscriptionHandler.Create)
			sub.GET("", subscriptionHandler.List)
			sub.GET("/:id", subscriptionHandler.GetByID)
			sub.PUT("/:id", subscriptionHandler.Update)
			sub.DELETE("/:id", subscriptionHandler.Delete)
		}

		v1.GET("/total", subscriptionHandler.GetTotalCost)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
