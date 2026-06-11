package middleware

import (
	"fmt"
	"net/http"

	"github.com/Sam-Frost/web-server/internal/dto"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process the request first

		// Check if any errors were added to the context
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, dto.APIResponse{
				Success: false,
				Error:   err,
			})
		}
	}
}
