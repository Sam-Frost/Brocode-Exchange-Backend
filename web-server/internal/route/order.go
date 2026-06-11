package route

import "github.com/gin-gonic/gin"

func OrderRouter(gin *gin.Engine) {
	orderRouter := gin.Group("/api/v1/order")

	orderRouter.GET("")
}
