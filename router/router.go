package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	InitializeRoutes(router)

	router.Run("0.0.0.0:5555")
}
