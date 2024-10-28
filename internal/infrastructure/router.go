package infrastructure

import (
	"github.com/Ateto1204/swep-user-serv/internal/delivery"
	"github.com/Ateto1204/swep-user-serv/internal/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(chatUseCase usecase.UserUseCase) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// router.Use(corsMiddleware())
	router.Use(cors.Default())

	userHandler := delivery.NewUserHandler(chatUseCase)

	router.GET("/", userHandler.Handle)
	router.POST("/api/user", userHandler.SaveUser)
	router.POST("api/user/id", userHandler.GetUser)

	return router
}

// func corsMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}
// 		c.Next()
// 	}
// }
