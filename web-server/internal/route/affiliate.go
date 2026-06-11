package route

import "github.com/gin-gonic/gin"

func AffiliateRouter(gin *gin.Engine) {
	affiliateRouter := gin.Group("/api/v1/affiliate")

	affiliateRouter.GET("")
}
