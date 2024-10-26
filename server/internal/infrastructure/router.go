package infrastructure

import (
	"github.com/Ateto1204/swep-user-serv/server/internal/delivery"
	"github.com/Ateto1204/swep-user-serv/server/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(chatUseCase usecase.UserUseCase) *gin.Engine {
	router := gin.Default()
	router.Use(corsMiddleware())

	userHandler := delivery.NewUserHandler(chatUseCase)

	router.GET("/", userHandler.Handle)
	router.POST("/api/user", userHandler.SaveUser)
	router.POST("api/user/id", userHandler.GetUser)

	return router
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
