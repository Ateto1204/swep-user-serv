package infrastructure

import (
	"github.com/Ateto1204/swep-user-serv/internal/delivery"
	"github.com/Ateto1204/swep-user-serv/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(chatUseCase usecase.UserUseCase) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(corsMiddleware())

	userHandler := delivery.NewUserHandler(chatUseCase)

	router.GET("/", userHandler.Handle)
	router.POST("/api", userHandler.SaveUser)
	router.POST("/api/id", userHandler.GetUser)
	router.DELETE("/api/user-del", userHandler.DeleteUser)
	router.PATCH("/api/friend-add", userHandler.AddNewFriend)
	router.PATCH("/api/friend-remove", userHandler.RemoveFriend)
	router.PATCH("/api/chat-add", userHandler.AddNewChat)
	router.PATCH("/api/chat-remove", userHandler.RemoveChat)

	return router
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
